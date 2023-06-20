package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type ArtifactChan struct {
	Nfts []entity.Nfts
	Err  error
	Done bool
}

type NftSizeChan struct {
	Nft entity.Nfts
	Err error
}

func (u *Usecase) CrontabUpdateImageSize(ctx context.Context) error {

	nftsChan := make(chan ArtifactChan)
	stopSig := make(chan bool)
	nfts := []entity.Nfts{}

	go func(stopSig chan bool) {
		contract := "0x16EfDc6D3F977E39DAc0Eb0E123FefFeD4320Bc0"
		page := 1
		limit := 50

		for {

			go u.GetNftArtifacts(ctx, contract, page, limit, nftsChan)
			page++

			stop := <-stopSig
			if stop {
				return
			}

		}

	}(stopSig)

	for {
		dataFromChan := <-nftsChan
		stopSig <- dataFromChan.Done

		nftsFromChan := dataFromChan.Nfts
		nfts = append(nfts, nftsFromChan...)

		if dataFromChan.Done {
			break
		}
	}

	logger.AtLog.Logger.Info("CrontabUpdateImageSize", zap.Int("nfts", len(nfts)))
	workerInputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan NftSizeChan, len(nfts))

	for i := 0; i < len(nfts); i++ {
		go u.UpdateImageSize(workerInputChan, outputChan)
	}

	for i := 0; i < len(nfts); i++ {
		workerInputChan <- nfts[i]
		if i > 0 && i%5 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	data := []entity.Nfts{}
	for i := 0; i < len(nfts); i++ {
		dataFromChan := <-outputChan
		if dataFromChan.Err != nil {
			logger.AtLog.Logger.Error("CrontabUpdateImageSize", zap.Error(dataFromChan.Err))
			continue
		}

		nft := dataFromChan.Nft
		//update DB
		updated, err := u.Repo.UpdateNftSize(nft.ContractAddress, nft.TokenID, nft.Size)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("CrontabUpdateImageSize - %s - %s", nft.ContractAddress, nft.TokenID), zap.Error(err), zap.Int64("size", nft.Size))
			return err
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("CrontabUpdateImageSize - %s - %s", nft.ContractAddress, nft.TokenID), zap.Any("updated", updated), zap.Int64("size", nft.Size))

		//only used for testing
		data = append(data, nft)
	}

	//err := helpers.CreateFile("nfts.json", data)
	return nil
}

func (u *Usecase) UpdateImageSize(input chan entity.Nfts, output chan NftSizeChan) {
	var err error
	inputData := <-input
	image := inputData.Image

	tokenURI := fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s", inputData.ContractAddress, inputData.TokenID)
	inputData.TokenURI = tokenURI

	image = fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s/content", inputData.ContractAddress, inputData.TokenID)
	inputData.Image = image

	res, headers, status, err := helpers.HttpRequest(image, "GET", make(map[string]string), nil)
	if err != nil {
		logger.AtLog.Logger.Info(fmt.Sprintf("UpdateImageSize - %s", inputData.TokenID), zap.Any("headers", headers), zap.Any("status", status))
	} else {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateImageSize - %s", inputData.TokenID), zap.Any("headers", headers), zap.Any("status", status), zap.Error(err))

	}

	inputData.Size = int64(len(res))
	defer func() {
		output <- NftSizeChan{
			Nft: inputData,
			Err: err,
		}
	}()

}

func (u *Usecase) GetNftArtifacts(ctx context.Context, contract string, page, limit int, nftsChan chan ArtifactChan) {
	offset := (page - 1) * limit
	var err error
	nfts := []entity.Nfts{}

	defer func() {
		done := false
		if len(nfts) == 0 || err != nil {
			done = true
		}

		nftsChan <- ArtifactChan{
			Nfts: nfts,
			Err:  err,
			Done: done,
		}
	}()

	nfts, err = u.Repo.GetNftsWithoutSize(contract, offset, limit)
	if err != nil {
		logger.AtLog.Logger.Error("GetNftArtifacts", zap.Error(err),
			zap.String("contract", contract),
			zap.Int("page", page),
			zap.Int("limit", limit),
			zap.Int("offset", offset),
		)

		return
	}

	logger.AtLog.Logger.Info("GetNftArtifacts",
		zap.String("contract", contract),
		zap.Int("page", page),
		zap.Int("limit", limit),
		zap.Int("offset", offset),
		zap.Int("nfts", len(nfts)),
	)
}
