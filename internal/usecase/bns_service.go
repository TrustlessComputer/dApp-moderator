package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/bns"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (u *Usecase) BnsNames(ctx context.Context, filter request.FilterBNSNames) ([]*entity.FilteredBNS, error) {
	f := entity.FilterBns{}
	err := helpers.JsonTransform(filter, &f)
	if err != nil {
		return nil, err
	}

	resp, err := u.Repo.FilterBNS(f)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsName(ctx context.Context, tokenID string) (*entity.FilteredBNS, error) {
	resp, err := u.Repo.FindBNS(tokenID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsNameAvailable(ctx context.Context, name string) (*bool, error) {
	resp, err := u.BnsService.NameAvailable(name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsNamesOnwedByWalletAddress(ctx context.Context, walletAdress string, filter request.FilterBNSNames) ([]*bns_service.NameResp, error) {
	params := url.Values{}
	if filter.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *filter.Limit))
	}

	if filter.Offset != nil {
		params.Set("offset", fmt.Sprintf("%d", *filter.Offset))
	}

	resp, err := u.BnsService.NameOnwedByWalletAddress(walletAdress, params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsDefault(ctx context.Context, resolver string) (*entity.Bns, error) {
	result, err := u.Repo.FindOne(utils.COLLECTION_BNS_DEFAULT, bson.D{{"resolver", resolver}})
	if err == nil { // Nếu đã có data bns_default rồi thì lấy bns ra và trả về
		bnsDefault := &entity.BNSDefault{}
		if err := result.Decode(bnsDefault); err == nil {
			var bnsEntity = &entity.Bns{} // SHOULD Optimize to 1 query
			if result, err := u.Repo.FindOne(utils.COLLECTION_BNS, bson.D{{"_id", bnsDefault.BNSDefaultID}}); err == nil {
				if err := result.Decode(bnsEntity); err == nil {
					return bnsEntity, nil
				}
			}
		}
	}

	// Nếu chưa có data bns_default thì lấy bns từ resolver và lưu vào db
	bnsEntities := []*entity.Bns{}
	err = u.Repo.Find(utils.COLLECTION_BNS, bson.D{{"resolver", resolver}}, 100, 0, &bnsEntities, bson.D{
		{"_id", -1},
	})
	if err != nil || len(bnsEntities) == 0 {
		logger.AtLog.Logger.Error("BNSDefault but dont have any bns items", zap.String("resolver", resolver))
		return nil, errors.New("BNSDefault but dont have any bns items")
	}

	// Ưu tiên lấy bns nào mà có pfp_data, ko thì lấy cái đầu tiên
	var bnsEntityPickToDefault *entity.Bns
	for index, item := range bnsEntities {
		if item.PfpData != nil {
			bnsEntityPickToDefault = bnsEntities[index]
			break
		}
	}
	if bnsEntityPickToDefault == nil {
		bnsEntityPickToDefault = bnsEntities[0]
	}

	// Lưu bns default vào db
	bnsDefault := &entity.BNSDefault{
		Resolver:     resolver,
		BNSDefaultID: bnsEntityPickToDefault.ID,
	}

	_, err = u.Repo.InsertOne(bnsDefault)
	if err != nil {
		return nil, err
	}

	return bnsEntityPickToDefault, nil
}

// Start: only called by test/main.go
func (u *Usecase) CrontabBns(ctx context.Context) error {
	page := 1
	limit := 100
	bnsAddress := strings.ToLower("0x8b46f89bba2b1c1f9ee196f43939476e79579798")
	done := make(chan bool)

	for {
		go u.GetBnsCollection(ctx, done, bnsAddress, page, limit)

		s := <-done
		if s {
			break
		}

		time.Sleep(time.Millisecond * 500)
		page++
	}

	return nil
}

func (u *Usecase) GetBnsCollection(ctx context.Context, done chan bool, bnsAddress string, page int, limit int) error {
	offset := (page - 1) * limit
	nfts, err := u.Repo.GetNfts(bnsAddress, offset, limit)
	isDone := false
	isDoneP := &isDone

	defer func() {
		if err != nil {
			isDone = true
		}

		done <- *isDoneP
	}()

	bnsS, err := bns.NewBns(common.HexToAddress(bnsAddress), u.TCPublicNode.GetClient())
	if err != nil {
		return err
	}

	key := fmt.Sprintf("CrontabBns - page %d, offset: %d, items: %d", page, offset, len(nfts))

	logger.AtLog.Info(key)
	if len(nfts) == 0 {
		isDone = true
		return nil
	}

	inputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan structure.BnsRespChan, len(nfts))

	for i := 0; i < len(nfts); i++ {
		go u.BnsItemWorker(ctx, bnsS, inputChan, outputChan)
	}

	go func() {
		for _, nft := range nfts {
			inputChan <- nft
		}
	}()

	for _, _ = range nfts {
		processed := <-outputChan
		if processed.Err == nil {
			bns := processed.Bns

			u.Repo.InsertOne(bns)
		}
	}

	return nil
}

type pfpChan struct {
	Data []byte
	Err  error
}

type ownerChan struct {
	Data common.Address
	Err  error
}

func (u *Usecase) BnsItemWorker(ctx context.Context, bns *bns.Bns, inputChan chan entity.Nfts, outputChan chan structure.BnsRespChan) {
	var err error
	nft := <-inputChan
	bnsData := &entity.Bns{}

	defer func() {
		outputChan <- structure.BnsRespChan{
			Bns: bnsData,
			Nft: nft,
			Err: err,
		}
	}()

	tokenID, _ := new(big.Int).SetString(nft.TokenID, 10)
	pfpFchan := make(chan pfpChan)
	nameFchan := make(chan pfpChan)
	ownerFchan := make(chan ownerChan)
	resolverFchan := make(chan ownerChan)

	go func(pfpFchan chan pfpChan) {
		pfp, err := bns.GetPfp(nil, tokenID)
		pfpFchan <- pfpChan{
			Data: pfp,
			Err:  err,
		}
	}(pfpFchan)

	go func(ownerFchan chan ownerChan) {
		owner, err := bns.OwnerOf(nil, tokenID)
		ownerFchan <- ownerChan{
			Data: owner,
			Err:  err,
		}
	}(ownerFchan)

	go func(resolverFchan chan ownerChan) {
		owner, err := bns.Resolver(nil, tokenID)
		resolverFchan <- ownerChan{
			Data: owner,
			Err:  err,
		}
	}(resolverFchan)

	go func(nameFchan chan pfpChan) {
		names, err := bns.Names(nil, tokenID)
		nameFchan <- pfpChan{
			Data: names,
			Err:  err,
		}
	}(nameFchan)

	pfp := <-pfpFchan
	owner := <-ownerFchan
	resolver := <-resolverFchan
	name := <-nameFchan

	if pfp.Err != nil {
		err = pfp.Err
		return
	}

	if name.Err != nil {
		err = name.Err
		return
	}

	if owner.Err != nil {
		err = owner.Err
		return
	}

	if resolver.Err != nil {
		err = resolver.Err
		return
	}

	bnsData.TokenID = nft.TokenID
	bnsData.Owner = strings.ToLower(owner.Data.Hex())
	bnsData.Resolver = strings.ToLower(resolver.Data.Hex())
	bnsData.Pfp = string(pfp.Data)
	bnsData.Name = string(name.Data)
	bnsData.CollectionAddress = strings.ToLower(nft.ContractAddress)
}

//End: only called by test/main.go
