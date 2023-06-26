package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/erc20"
	"dapp-moderator/utils/generative_nft_contract"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CaptureSoulImageChan struct {
	Nft              entity.Nfts
	Err              error
	Image            *string
	AnimationFileUrl *string
	Traits           *[]entity.NftAttr
	ReplacedTraits   *map[string]string
}

type ReplaceHtmlWithTraits struct {
	URL            string
	CapturedImage  string
	Traits         *[]entity.NftAttr
	ReplacedTraits *map[string]string
}

type CaptureSoulImagesChan struct {
	Nft entity.Nfts
	Err error

	//thumbnail + original html
	Image             *string
	Html              *string
	AnimationFileUrls *[]*ReplaceHtmlWithTraits
}

type CaptureSoulOwnerChan struct {
	Err          error
	Nft          entity.Nfts
	Owner        *string
	BlockNumber  *uint64
	Erc20Address *string
	Erc20Amount  *string
}

func (u *Usecase) SoulNftImageCrontab() error {

	limit := 3
	page := 1

	for {
		offset := (page - 1) * limit

		addr := os.Getenv("SOUL_CONTRACT")
		nfts, err := u.Repo.NftWithoutCapturedImage(addr, offset, limit)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit), zap.Error(err))
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit), zap.Int("nfts", len(nfts)))
		if len(nfts) == 0 {
			break
		}

		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		inputWorker1Chan := make(chan entity.Nfts, len(nfts))
		outputFromWorker1Chan := make(chan CaptureSoulImagesChan, len(nfts))
		outputFromWorker2Chan := make(chan CaptureSoulImagesChan, len(nfts))

		for i := 0; i < len(nfts); i++ {
			go u.GetSoulNftAnimationURLWorkerNew(&wg1, inputWorker1Chan, outputFromWorker1Chan)
		}

		for i := 0; i < len(nfts); i++ {
			go u.CaptureSoulNftImagesWorker(&wg2, outputFromWorker1Chan, outputFromWorker2Chan)
		}

		for _, nft := range nfts {
			wg1.Add(1)
			wg2.Add(1)
			inputWorker1Chan <- nft
		}

		for i := 0; i < len(nfts); i++ {
			out := <-outputFromWorker2Chan

			if out.Err != nil {
				logger.AtLog.Logger.Error("Debug", zap.Error(err))
				continue
			}

			output := *out.AnimationFileUrls
			if len(output) < 5 {
				err = errors.New("Cannot capture images")
				logger.AtLog.Logger.Error("Debug", zap.Error(err))
				continue
			}

			for _, soulImage := range output {
				wg3.Add(1)
				go u.CreateSoulNftImages(&wg3, CaptureSoulImageChan{
					Err:              out.Err,
					Nft:              out.Nft,
					Image:            &soulImage.CapturedImage,
					AnimationFileUrl: &soulImage.URL,
					Traits:           soulImage.Traits,
					ReplacedTraits:   soulImage.ReplacedTraits,
				})
			}

			//TODO - logic will be applied here
			wg3.Add(1)

			image := output[0].CapturedImage
			traits := output[0].Traits
			animationURL := out.Html
			updatedData := CaptureSoulImageChan{
				Err:              out.Err,
				Nft:              out.Nft,
				Image:            &image,
				AnimationFileUrl: animationURL,
				Traits:           traits,
			}

			logger.AtLog.Logger.Info("Debug", zap.Any("updatedData", updatedData), zap.Any("output", output))
			go u.UpdateSoulNftImageWorker(&wg3, updatedData)
		}

		wg1.Wait()
		wg2.Wait()
		wg3.Wait()

		page++
	}

	return nil
}

func (u *Usecase) SoulNftImageHistoriesCrontab(specialNfts []string) error {
	logger.AtLog.Logger.Info("SoulNftImageHistoriesCrontab", zap.Any("specialNfts", specialNfts))
	gmAddress := os.Getenv("SOUL_GM_ADDRESS")
	url := fmt.Sprintf("https://www.fprotocol.io/api/swap/token/report?address=%s", gmAddress)
	rate, _, _, err := helpers.JsonRequest(url, "GET", map[string]string{}, nil)
	if err != nil {
		return err
	}

	resp := &structure.ReportErc20{}
	err = json.Unmarshal(rate, resp)
	if err != nil {
		return err
	}

	erc20Contract, err := erc20.NewErc20(common.HexToAddress(gmAddress), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulNftImageHistoriesCrontab", zap.Error(err))
		return err
	}

	nftContract, err := generative_nft_contract.NewGenerativeNftContract(common.HexToAddress(os.Getenv("SOUL_CONTRACT")), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulNftImageHistoriesCrontab", zap.Error(err))
		return err
	}
	limit := 3
	page := 1

	for {
		offset := (page - 1) * limit

		addr := os.Getenv("SOUL_CONTRACT")
		nfts, err := u.Repo.NftCapturedImageHistories(addr, offset, limit, specialNfts)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulNftImageHistoriesCrontab - page: %d, limit: %d", page, limit), zap.Error(err))
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("SoulNftImageHistoriesCrontab - page: %d, limit: %d", page, limit), zap.Int("nfts", len(nfts)))
		if len(nfts) == 0 {
			break
		}

		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		var wg4 sync.WaitGroup
		inputWorker1Chan := make(chan entity.Nfts, len(nfts))
		inputWorker3Chan := make(chan entity.Nfts, len(nfts))

		outputFromWorker1Chan := make(chan CaptureSoulImageChan, len(nfts))
		outputFromWorker2Chan := make(chan CaptureSoulImageChan, len(nfts))
		outputFromWorker3Chan := make(chan CaptureSoulOwnerChan, len(nfts))

		for i := 0; i < len(nfts); i++ {
			go u.GetSoulNftAnimationURLWorker(&wg1, inputWorker1Chan, outputFromWorker1Chan)
		}

		for i := 0; i < len(nfts); i++ {
			go u.CaptureSoulNftImageWorker(&wg2, outputFromWorker1Chan, outputFromWorker2Chan)
		}

		for i := 0; i < len(nfts); i++ {
			go u.GetSoulNftOwnerWorker(&wg4, inputWorker3Chan, erc20Contract, nftContract, outputFromWorker3Chan)
		}

		for _, nft := range nfts {
			wg1.Add(1)
			wg2.Add(1)
			wg4.Add(1)

			inputWorker1Chan <- nft
			inputWorker3Chan <- nft
		}

		for i := 0; i < len(nfts); i++ {
			out := <-outputFromWorker2Chan
			out1 := <-outputFromWorker3Chan

			wg3.Add(1)

			go u.UpdateSoulNftImageImageHistoriesWorker(&wg3, resp, out, out1)

		}

		wg1.Wait()
		wg2.Wait()
		wg3.Wait()

		page++
	}

	return nil
}

func (u *Usecase) GetSoulNftAnimationURLWorker(wg *sync.WaitGroup, inputChan chan entity.Nfts, outputChan chan CaptureSoulImageChan) {
	ctx := context.Background()
	defer wg.Done()
	nft := <-inputChan
	var err error
	animationFileUrlP := new(string)

	defer func() {

		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("GetSoulNftAnimationURLWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("animationFileUrlP", animationFileUrlP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("GetSoulNftAnimationURLWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

		outputChan <- CaptureSoulImageChan{
			Err:              err,
			Nft:              nft,
			AnimationFileUrl: animationFileUrlP,
		}
	}()

	animationFileUrl, err := u.GetAnimationFileUrl(ctx, &nft)
	animationFileUrlP = &animationFileUrl
}

func (u *Usecase) GetSoulNftAnimationURLWorkerNew(wg *sync.WaitGroup, inputChan chan entity.Nfts, outputChan chan CaptureSoulImagesChan) {

	defer wg.Done()
	nft := <-inputChan
	var err error
	animationFileUrlP := new([]*ReplaceHtmlWithTraits)
	animationHtmlOriginal := new(string)
	defer func() {

		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("GetSoulNftAnimationURLWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("animationFileUrlP", animationFileUrlP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("GetSoulNftAnimationURLWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

		outputChan <- CaptureSoulImagesChan{
			Err:               err,
			Nft:               nft,
			AnimationFileUrls: animationFileUrlP,
			Html:              animationHtmlOriginal,
		}
	}()

	contractS, err := generative_nft_contract.NewGenerativeNftContract(common.HexToAddress(nft.ContractAddress),
		u.TCPublicNode.GetClient())

	if err != nil {
		return
	}

	tokenIdInt, _ := strconv.Atoi(nft.TokenID)
	tokenBigInt := big.NewInt(int64(tokenIdInt))

	tokenUriData, err := contractS.TokenURI(&bind.CallOpts{Context: context.Background()}, tokenBigInt)
	if err != nil {
		return
	}

	tokenUri := entity.TokenUri{}
	if err := json.Unmarshal([]byte(tokenUriData), &tokenUri); err != nil {
		return
	}
	if tokenUri.AnimationUrl == "" {
		err = errors.New("animation url is empty")
		return
	}

	animationHtmlOriginal = &tokenUri.AnimationUrl
	imageUrls := []*ReplaceHtmlWithTraits{}
	if strings.Contains(tokenUri.AnimationUrl, "base64") {

		html, err := u.ReplaceSoulHtml(tokenUri.AnimationUrl)
		if err != nil {
			return
		}

		for i := 0; i <= 4; i++ {
			//TODO - replace via random number here
			html1 := *html
			randomArray := make(map[string]string)
			if i != 0 {
				capKey := fmt.Sprintf("capture%d", i)
				replaced := fmt.Sprintf("%s=!1", capKey)
				replaceTo := fmt.Sprintf("%s=true", capKey)
				randomArray[replaced] = replaceTo
				html1 = strings.ReplaceAll(html1, replaced, replaceTo)
			}

			encoded := helpers.Base64Encode(html1)
			fileName := fmt.Sprintf("%v_%v_%v.html", nft.ContractAddress, nft.TokenID, time.Now().UTC().Unix())
			resp, err := u.Storage.UploadBaseToBucket(encoded, fmt.Sprintf("capture_animation_file/%v", fileName))
			if err != nil {
				return
			}

			item := &ReplaceHtmlWithTraits{}
			htmlFileLink := fmt.Sprintf("https://storage.googleapis.com%v", resp.Path)

			item.URL = htmlFileLink
			item.ReplacedTraits = &randomArray
			imageUrls = append(imageUrls, item)
		}

		animationFileUrlP = &imageUrls
		return
	}

	return
}

func (u *Usecase) GetSoulNftOwnerWorker(wg *sync.WaitGroup, inputChan chan entity.Nfts, erc20Contract *erc20.Erc20, nftContract *generative_nft_contract.GenerativeNftContract, outputChan chan CaptureSoulOwnerChan) {

	defer wg.Done()
	nft := <-inputChan
	var err error
	ownerP := new(string)
	erc20AmountP := new(string)
	blockNumberP := new(uint64)

	defer func() {
		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("GetSoulNftOwnerWorker - %s, %s", nft.ContractAddress, nft.TokenID))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("GetSoulNftOwnerWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

		outputChan <- CaptureSoulOwnerChan{
			Err:         err,
			Nft:         nft,
			Owner:       ownerP,
			Erc20Amount: erc20AmountP,
			BlockNumber: blockNumberP,
		}
	}()

	blockNumber, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		return
	}

	bn := blockNumber.Uint64()
	blockNumberP = &bn

	tokenID, _ := new(big.Int).SetString(nft.TokenID, 10)
	owner, err := nftContract.OwnerOf(nil, tokenID)
	if err != nil {
		return
	}

	balance, err := erc20Contract.BalanceOf(nil, owner)
	if err != nil {
		return
	}

	b := fmt.Sprintf("%d", balance.Int64())
	o := owner.String()

	ownerP = &o
	erc20AmountP = &b
}

func (u *Usecase) CaptureSoulNftImageWorker(wg *sync.WaitGroup, inputChan chan CaptureSoulImageChan, outputChan chan CaptureSoulImageChan) {
	defer wg.Done()
	inChan := <-inputChan
	newImagePathP := new(string)
	traitP := new([]entity.NftAttr)
	var err error
	nft := inChan.Nft

	defer func() {

		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("CaptureSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("newImagePathP", newImagePathP), zap.Any("traitP", traitP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("CaptureSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

		inChan.Image = newImagePathP
		inChan.Traits = traitP
		inChan.Err = err
		outputChan <- inChan
	}()

	if inChan.Err != nil {
		err = inChan.Err
		return
	}

	newImagePath, traits, err := u.ParseHtmlImage(*inChan.AnimationFileUrl)
	if err != nil {
		return
	}

	newImagePathP = &newImagePath

	traitObjs := []entity.NftAttr{}
	for key, value := range traits {
		t := entity.NftAttr{}
		t.TraitType = key
		t.Value = value
		traitObjs = append(traitObjs, t)
	}

	traitP = &traitObjs
}

func (u *Usecase) UpdateSoulNftImageWorker(wg *sync.WaitGroup, inputChan CaptureSoulImageChan) {
	defer wg.Done()
	out := inputChan
	nft := out.Nft

	newImagePathP := new(string)
	traitP := new([]entity.NftAttr)
	var err error

	defer func() {
		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("newImagePathP", newImagePathP), zap.Any("traitP", traitP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

	}()

	err = out.Err
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.Error(err),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
		)
		return
	}

	image := out.Image
	traits := out.Traits
	animationFileUrl := out.AnimationFileUrl

	updated, err := u.Repo.UpdateOne(utils.COLLECTION_NFTS, bson.D{{"_id", nft.ID}}, bson.M{"$set": bson.M{
		"image_capture":      *image,
		"animation_file_url": animationFileUrl,
		"attributes":         *traits,
		"image_capture_at":   time.Now().UTC(),
	}})

	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.Error(err),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
			zap.Any("image", image),
			zap.Any("traits", traits),
			zap.Any("animationFileUrl", animationFileUrl),
		)
		return
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID),
		zap.String("tokenID", nft.TokenID),
		zap.String("contractAddress", nft.ContractAddress),
		zap.Any("image", image),
		zap.Any("traits", traits),
		zap.Any("animationFileUrl", animationFileUrl),
		zap.Any("updated", updated),
	)
}

func (u *Usecase) UpdateSoulNftImageImageHistoriesWorker(wg *sync.WaitGroup, bitcoindex *structure.ReportErc20, inputChan CaptureSoulImageChan, input1 CaptureSoulOwnerChan) {
	defer wg.Done()
	out := inputChan
	nft := out.Nft

	newImagePathP := new(string)
	traitP := new([]entity.NftAttr)
	var err error

	defer func() {
		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("UpdateSoulNftImageImageHistoriesWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("newImagePathP", newImagePathP), zap.Any("traitP", traitP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageImageHistoriesWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

	}()

	err = bitcoindex.Error
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageImageHistoriesWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.Error(err),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
		)
		return
	}

	err = input1.Err
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageImageHistoriesWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.Error(err),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
		)
		return
	}

	err = out.Err
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageImageHistoriesWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.Error(err),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
		)
		return
	}

	ownerP := input1.Owner
	erc20P := input1.Erc20Amount

	image := out.Image
	now := time.Now().UTC()

	owner := ""
	if ownerP != nil {
		owner = *ownerP
	}

	erc20Amount := "0"
	if erc20P != nil {
		erc20Amount = *erc20P
	}

	bn := uint64(0)
	if input1.BlockNumber != nil {
		bn = *input1.BlockNumber
	}

	obj := &entity.SoulImageHistories{
		ContractAddress:  strings.ToLower(nft.ContractAddress),
		TokenID:          nft.TokenID,
		TokenIDInt:       nft.TokenIDInt,
		ImageCapture:     *image,
		ImageCaptureAt:   &now,
		ImageCaptureDate: fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()),
		Erc20Address:     strings.ToLower(os.Getenv("SOUL_GM_ADDRESS")),
		Erc20Amount:      erc20Amount,
		BlockNumber:      bn,
		Owner:            strings.ToLower(owner),
	}

	if len(bitcoindex.Data) >= 1 {
		price := bitcoindex.Data[0]
		obj.BitcoinDexWETHPrice = price.Price
		obj.BitcoinDexWBTCPrice = price.BtcPrice
		obj.BitcoinDexUSDTPrice = price.UsdPrice
	} else {
		obj.BitcoinDexWETHPrice = "0"
		obj.BitcoinDexWBTCPrice = "0"
		obj.BitcoinDexUSDTPrice = "0"
	}

	err = u.Repo.InsertSoulImageHistory(obj)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID),
			zap.String("tokenID", nft.TokenID),
			zap.String("contractAddress", nft.ContractAddress),
			zap.Any("image", image),
			zap.Error(err),
		)
		return
	}

	logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID),
		zap.String("tokenID", nft.TokenID),
		zap.String("contractAddress", nft.ContractAddress),
		zap.Any("image", image),
		zap.Any("histories", obj),
	)
}

func (u *Usecase) CaptureSoulNftImagesWorker(wg *sync.WaitGroup, inputChan chan CaptureSoulImagesChan, outputChan chan CaptureSoulImagesChan) {
	defer wg.Done()
	inChan := <-inputChan
	newImagePathP := new(string)
	traitP := new([]entity.NftAttr)
	var err error
	nft := inChan.Nft

	defer func() {

		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("CaptureSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Any("newImagePathP", newImagePathP), zap.Any("traitP", traitP))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("CaptureSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

		inChan.Image = newImagePathP
		inChan.Err = err
		outputChan <- inChan
	}()

	if inChan.Err != nil {
		err = inChan.Err
		return
	}

	for _, item := range *inChan.AnimationFileUrls {
		newImagePath, traits, err := u.ParseHtmlImage(item.URL)
		if err != nil {
			return
		}

		traitObjs := []entity.NftAttr{}
		for key, value := range traits {
			t := entity.NftAttr{}
			t.TraitType = key
			t.Value = value
			traitObjs = append(traitObjs, t)
		}

		item.Traits = &traitObjs
		item.CapturedImage = newImagePath
	}
}

func (u *Usecase) CreateSoulNftImages(wg *sync.WaitGroup, inputChan CaptureSoulImageChan) {
	defer wg.Done()
	var err error
	nft := entity.Nfts{}

	defer func() {
		if err == nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID))
		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateSoulNftImageWorker - %s, %s", nft.ContractAddress, nft.TokenID), zap.Error(err))
		}

	}()

	if inputChan.Err != nil {
		err = inputChan.Err
		return
	}

	nft = inputChan.Nft
	soulImage := &entity.SoulImages{
		ContractAddress:    strings.ToLower(nft.ContractAddress),
		TokenID:            strings.ToLower(nft.TokenID),
		TokenIDInt:         nft.TokenIDInt,
		Image:              inputChan.Image,
		AnimationURL:       inputChan.AnimationFileUrl,
		ReplacedAttributes: inputChan.ReplacedTraits,
	}

	_, err = u.Repo.InsertOne(soulImage)

}
