package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils/contracts/bns"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/url"
	"strings"
	"time"
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
