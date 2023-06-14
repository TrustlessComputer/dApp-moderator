package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/erc20"
	"dapp-moderator/utils/contracts/generative_nft_contract"
	"dapp-moderator/utils/contracts/soul"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type CheckGMBalanceOutputChan struct {
	Nft         entity.Nfts
	Err         error
	Balance     *big.Int
	IsAvailable bool
}

func (u *Usecase) SoulCrontab() error {
	maxProcess := 10
	minBalance := float64(1)
	erc20Addr := strings.ToLower(os.Getenv("SOUL_GM_ADDRESS"))

	erc20instance, err := erc20.NewErc20(common.HexToAddress(erc20Addr), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	soulInstance, err := soul.NewSoul(common.HexToAddress(os.Getenv("SOUL_CONTRACT")), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	collection, err := u.Repo.GetSoulCollection()
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	nfts, err := u.Repo.SoulNfts(collection.Contract)
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	inputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan CheckGMBalanceOutputChan, len(nfts))
	wg := sync.WaitGroup{}
	logger.AtLog.Logger.Info("SoulCrontab", zap.String("contract_address", collection.Contract), zap.Int("nfts", len(nfts)))

	for i := 0; i < len(nfts); i++ {
		go u.CheckGMBalanceWorker(&wg, erc20instance, soulInstance, inputChan, outputChan)
	}

	for i, nft := range nfts {
		wg.Add(1)
		inputChan <- nft
		if i%maxProcess == 0 && i > 0 || i == len(nfts)-1 {
			wg.Wait()
		}
	}

	for i := 0; i < len(nfts); i++ {
		out := <-outputChan
		if out.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(out.Err))
			continue
		}

		tokenIDInt, err := strconv.Atoi(out.Nft.TokenID)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(err))
			continue
		}

		isAuction := false
		isAvailable := out.IsAvailable

		value := helpers.GetValue(fmt.Sprintf("%d", out.Balance.Int64()), 18)
		if value < minBalance && isAvailable {
			isAuction = true
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.String("owner", out.Nft.Owner), zap.String("balance", fmt.Sprintf("%d", out.Balance.Int64())))

		insertData := &entity.NftAuctionsAvailable{
			TokenID:         out.Nft.TokenID,
			TokenIDInt:      int64(tokenIDInt),
			ContractAddress: strings.ToLower(out.Nft.ContractAddress),
			IsAuction:       isAuction,
		}

		err = u.Repo.InsertAuction(insertData)
	}
	return nil
}

func (u *Usecase) CheckGMBalanceWorker(wg *sync.WaitGroup, erc20Instance *erc20.Erc20, soulInstance *soul.Soul, input chan entity.Nfts, output chan CheckGMBalanceOutputChan) {
	defer wg.Done()
	nft := <-input
	var err error
	isAvailableP := new(bool)
	isAvailable := false
	isAvailableP = &isAvailable
	balanceOf := &big.Int{}

	defer func() {
		outData := CheckGMBalanceOutputChan{
			Nft:         nft,
			Balance:     balanceOf,
			Err:         err,
			IsAvailable: *isAvailableP,
		}

		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("CheckGMBalanceWorker - %s", nft.TokenID), zap.String("tokenID", nft.TokenID), zap.Error(err), zap.Any("outData", outData))
		} else {
			logger.AtLog.Logger.Info(fmt.Sprintf("CheckGMBalanceWorker - %s", nft.TokenID), zap.String("tokenID", nft.TokenID), zap.Any("outData", outData))
		}

		output <- outData
	}()

	owner := nft.Owner
	balanceOf, err = erc20Instance.BalanceOf(nil, common.HexToAddress(owner))
	if err != nil {
		return
	}

	tokenID, isSet := new(big.Int).SetString(nft.TokenID, 10)
	if isSet == false {
		err = errors.New("Cannot parse tokenID")
		return
	}

	isAvailable, err = soulInstance.Available(nil, tokenID)
	if err != nil {
		return
	}
}

func (u *Usecase) FilterSoulNfts(ctx context.Context, filter entity.FilterNfts) ([]*nft_explorer.SoulNft, error) {
	resp := []*nft_explorer.SoulNft{}
	f := bson.D{}

	maxFileSize := os.Getenv("FILE_CHUNK_SIZE")
	if filter.IsBigFile != nil {
		maxFileSizeInt, _ := strconv.Atoi(maxFileSize)
		if *filter.IsBigFile == true {
			f = append(f, bson.E{"size", bson.M{"$gte": maxFileSizeInt}})
		} else {
			f = append(f, bson.E{"size", bson.M{"$lt": maxFileSizeInt}})
		}
	}

	if filter.IsBuyable != nil {
		f = append(f, bson.E{"buyable", *filter.IsBuyable})
	}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		f = append(f, bson.E{"collection_address", *filter.ContractAddress})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		f = append(f, bson.E{"token_id", *filter.TokenID})
	}

	if filter.Rarity != nil {
		filter.Rarity.Min = filter.Rarity.Min / 100
		filter.Rarity.Max = filter.Rarity.Max / 100
		//f = append(f, bson.E{"$and", bson.A{
		//	bson.E{"attributes.percent", bson.M{"$lte": filter.Rarity.Max / 100}},
		//	bson.E{"attributes.percent", bson.M{"$gte": filter.Rarity.Min / 100}},
		//}})

		attrs, err := u.Repo.FilterCollectionAttributeByPercent(entity.FilterMarketplaceCollectionAttribute{
			ContractAddress: filter.ContractAddress,
			MaxPercent:      &filter.Rarity.Max,
			MinPercent:      &filter.Rarity.Min,
		})

		if err != nil {
			return nil, err
		}

		key := []string{}
		value := []string{}
		for _, attr := range attrs {
			key = append(key, attr.TraitType)
			value = append(value, attr.Value)
		}

		filter.AttrKey = key
		filter.AttrValue = value
	}

	if filter.Price != nil {
		btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))
		ethRate := u.GetExternalRate(os.Getenv("WETH_ADDRESS"))
		rate := btcRate / ethRate

		minPrice := filter.Price.Min
		maxPrice := filter.Price.Max

		minPriceEth := minPrice * rate
		maxPriceEth := maxPrice * rate

		fPrice := bson.A{
			bson.D{
				{"$and",
					bson.A{
						bson.D{{"erc20", strings.ToLower(os.Getenv("WBTC_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gt", minPrice}}}},
						bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
					},
				},
			},
			bson.D{
				{"$and",
					bson.A{
						bson.D{{"erc20", strings.ToLower(os.Getenv("WETH_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gt", minPriceEth}}}},
						bson.D{{"price", bson.D{{"$lte", maxPriceEth}}}},
					},
				},
			},
		}

		f = append(f, bson.E{"$or", fPrice})

	}

	if len(filter.AttrKey) > 0 {
		f = append(f, bson.E{"attributes.trait_type", bson.M{"$in": filter.AttrKey}})
	}

	if len(filter.AttrValue) > 0 {
		f = append(f, bson.E{"attributes.value", bson.M{"$in": filter.AttrValue}})
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	sortBy := "token_id_int"
	if filter.SortBy != "" {
		sortBy = filter.SortBy
	}

	sort := -1
	if filter.Sort != 0 {
		sort = int(filter.Sort)
	}

	s := bson.D{
		{"buyable", -1},
		{"price_erc20.price", 1},
		{sortBy, sort},
	}
	//old: VIEW_MARKETPLACE_NFTS, VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES has attributes + percent

	projections := bson.D{
		{"activities", 0},
	}

	err := u.Repo.FindWithProjections(utils.VIEW_NFT_AUCTION, f, int64(filter.Limit), int64(filter.Offset), &resp, s, projections)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) CreateSignature(requestData request.CreateSignatureRequest) (*structure.CreateSignatureResp, error) {
	soulChainID := os.Getenv("SOUL_CHAIN_ID")
	chainID, _ := new(big.Int).SetString(soulChainID, 10)
	contractAddr := strings.ToLower(os.Getenv("SOUL_CONTRACT"))
	signerMint := strings.ToLower(os.Getenv("SOUL_SIGNATURE_PUBLIC_KEY"))
	userWalletAddress := strings.ToLower(requestData.WalletAddress)
	gmTokenAddress := strings.ToLower(os.Getenv("SOUL_GM_ADDRESS"))
	var err error
	gm := float64(0)

	soulInstance, err := soul.NewSoul(common.HexToAddress(contractAddr), u.TCPublicNode.GetClient())
	if err != nil {
		return nil, err
	}

	minted, err := soulInstance.Minted(nil, common.HexToAddress(userWalletAddress))
	if err != nil {
		return nil, err
	}

	if minted.Int64() > 0 {
		return nil, errors.New("User has a minted token")
	}

	testAccounts := strings.Split(strings.ReplaceAll(os.Getenv("SOUL_TEST_ACCOUNT"), " ", ""), ",")

	if !helpers.InArray(userWalletAddress, testAccounts) {
		key := fmt.Sprintf("gm.deposit.%s", userWalletAddress)
		existed, _ := u.Cache.Exists(key)
		if !*existed {
			gm, err = u.GMDeposit(userWalletAddress)
			if err != nil {
				return nil, err
			}

			err = u.Cache.SetData(key, gm)
			if err != nil {
				return nil, err
			}
		}

		cached, _ := u.Cache.GetData(key)
		gm, _ = strconv.ParseFloat(*cached, 10)

		if gm < 1 {
			return nil, errors.New("Not enough GM")
		}
	}

	gmAmount := helpers.ConvertAmount(gm)
	//deposit GM - generative
	totalGM, _ := gmAmount.Int64()
	g := big.NewInt(totalGM)
	messageHash, signature, err := u.PnftReferralPaymentSignMessage(contractAddr, *chainID, signerMint, userWalletAddress, gmTokenAddress, *g)
	if err != nil {
		return nil, err
	}

	resp := &structure.CreateSignatureResp{
		GM:          g.String(),
		Signature:   signature,
		MessageHash: messageHash,
	}
	return resp, nil
}

func (u *Usecase) PnftReferralPaymentSignMessage(contractAddr string, chainID big.Int, signerMint string, userWalletAddress string, gmTokenAddress string, totalGM big.Int) (string, string, error) {
	privateKey := strings.ToLower(os.Getenv("SOUL_SIGNATURE_PRIVATE_KEY"))
	datas := []byte{}

	datas = append(datas, common.HexToHash(contractAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(chainID.Bytes()).Bytes()...)
	datas = append(datas, common.HexToHash(signerMint).Bytes()...)
	datas = append(datas, common.HexToHash(userWalletAddress).Bytes()...)
	datas = append(datas, common.HexToHash(gmTokenAddress).Bytes()...)
	datas = append(datas, common.BytesToHash(totalGM.Bytes()).Bytes()...)

	dataByteHash := crypto.Keccak256Hash(
		datas,
	)

	signature, err := u.SignWithEthereum(privateKey, dataByteHash.Bytes())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("PnftReferralPaymentSignMessage - %s", userWalletAddress), zap.String("userWalletAddress", userWalletAddress), zap.String("contractAddr", contractAddr), zap.String("chainID", chainID.String()), zap.Error(err))
		return "", "", err
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("PnftReferralPaymentSignMessage - %s", userWalletAddress), zap.String("userWalletAddress", userWalletAddress), zap.String("contractAddr", contractAddr), zap.String("chainID", chainID.String()), zap.String("signature", signature))

	return dataByteHash.Hex(), signature, nil
}

func (u *Usecase) SignWithEthereum(privateKey string, dataBytes []byte) (string, error) {
	signBytes := append([]byte("\x19Ethereum Signed Message:\n32"), dataBytes...)
	hash := crypto.Keccak256Hash(signBytes)
	prk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hash.Bytes(), prk)
	if err != nil {
		return "", err
	}
	signature[crypto.RecoveryIDOffset] += 27
	sigHex := hexutil.Encode(signature)

	return sigHex, nil
}

func (u *Usecase) GMDeposit(walletAddress string) (float64, error) {
	generativeURL := "https://generative.xyz/generative/api/charts/gm-collections/deposit"
	resp := &structure.GMDepositResponse{}

	data, _, _, err := helpers.JsonRequest(generativeURL, "GET", make(map[string]string), nil)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return 0, err
	}

	items := resp.Data.Items
	for _, item := range items {
		if strings.ToLower(item.From) == strings.ToLower(walletAddress) {
			return item.GmReceive, nil
		}
	}

	return 0, nil
}

func (u *Usecase) CaptureSoulImage(ctx context.Context, request *request.CaptureSoulTokenReq) (*entity.Nfts, error) {
	nftEntity, err := u.Repo.GetNft(request.ContractAddress, request.TokenID)
	if err != nil {
		return nil, err
	}

	animationFileUrl := nftEntity.AnimationFileUrl
	var imagePath = nftEntity.Image
	if animationFileUrl == "" {
		animationFileUrl, err = u.GetAnimationFileUrl(ctx, nftEntity)
		if err != nil {
			return nil, err
		}
	}

	newImagePath := u.ParseSvgImage(animationFileUrl)
	if newImagePath == animationFileUrl {
		return nil, errors.New("parse svg image error")
	}
	if newImagePath != "" {
		imagePath = newImagePath
	}

	_, err = u.Repo.UpdateOne(utils.COLLECTION_NFTS, bson.D{{"_id", nftEntity.ID}}, bson.M{"$set": bson.M{
		"image_capture":      imagePath,
		"animation_file_url": animationFileUrl,
	}})

	if err != nil {
		return nil, err
	}

	nftEntity.ImageCapture = imagePath
	nftEntity.AnimationFileUrl = animationFileUrl

	return nftEntity, nil
}

func (u *Usecase) GetAnimationFileUrl(ctx context.Context, nftEntity *entity.Nfts) (string, error) {
	contractS, err := generative_nft_contract.NewGenerativeNftContract(common.HexToAddress(nftEntity.ContractAddress),
		u.TCPublicNode.GetClient())

	if err != nil {
		return "", err
	}

	tokenIdInt, _ := strconv.Atoi(nftEntity.TokenID)
	tokenBigInt := big.NewInt(int64(tokenIdInt))

	tokenUriData, err := contractS.TokenURI(&bind.CallOpts{Context: context.Background()}, tokenBigInt)
	if err != nil {
		return "", err
	}

	tokenUri := entity.TokenUri{}
	if err := json.Unmarshal([]byte(tokenUriData), &tokenUri); err != nil {
		return "", err
	}
	if tokenUri.AnimationUrl == "" {
		return "", errors.New("animation url is empty")
	}
	if strings.Contains(tokenUri.AnimationUrl, "base64") {
		tokenUri.AnimationUrl = strings.Replace(tokenUri.AnimationUrl, "data:text/html;base64,", "", -1)

		fileName := fmt.Sprintf("%v_%v_%v.html", nftEntity.ContractAddress, nftEntity.TokenID, time.Now().UTC().Unix())
		resp, err := u.Storage.UploadBaseToBucket(tokenUri.AnimationUrl, fmt.Sprintf("capture_animation_file/%v", fileName))
		if err != nil {
			return "", err
		}

		htmlFileLink := fmt.Sprintf("https://storage.googleapis.com%v", resp.Path)
		return htmlFileLink, nil
	}

	return tokenUri.AnimationUrl, nil
}

func (u *Usecase) SoulNftDetail(ctx context.Context, contractAddress string, tokenID string) (*entity.NftAuctionsAvailable, error) {

	data, err := u.Repo.FindAuction(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNftDetail", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
	}

	return data, nil
}
