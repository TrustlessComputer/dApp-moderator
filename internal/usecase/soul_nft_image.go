package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"os"
	"sync"
)

type CaptureSoulImageChan struct {
	Nft              entity.Nfts
	Err              error
	Image            *string
	AnimationFileUrl *string
	Traits           *[]entity.NftAttr
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
		inputWorker1Chan := make(chan entity.Nfts, len(nfts))
		outputFromWorker1Chan := make(chan CaptureSoulImageChan, len(nfts))
		outputFromWorker2Chan := make(chan CaptureSoulImageChan, len(nfts))

		for i := 0; i < len(nfts); i++ {
			go u.GetSoulNftAnimationURLWorker(&wg1, inputWorker1Chan, outputFromWorker1Chan)
		}

		for i := 0; i < len(nfts); i++ {
			go u.CaptureSoulNftImageWorker(&wg2, outputFromWorker1Chan, outputFromWorker2Chan)
		}

		for _, nft := range nfts {
			wg1.Add(1)
			wg2.Add(1)
			inputWorker1Chan <- nft
		}

		for i := 0; i < len(nfts); i++ {
			out := <-outputFromWorker2Chan
			nft := out.Nft
			err = out.Err
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit),
					zap.Error(err),
					zap.String("tokenID", nft.TokenID),
					zap.String("contractAddress", nft.ContractAddress),
				)
				continue
			}

			image := out.Image
			traits := out.Traits
			animationFileUrl := out.AnimationFileUrl

			if os.Getenv("ENV") != "mainnet" {
				updated, err := u.Repo.UpdateOne(utils.COLLECTION_NFTS, bson.D{{"_id", nft.ID}}, bson.M{"$set": bson.M{
					"image_capture":      *image,
					"animation_file_url": animationFileUrl,
					"attributes":         *traits,
				}})

				if err != nil {
					logger.AtLog.Logger.Error(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit),
						zap.Error(err),
						zap.String("tokenID", nft.TokenID),
						zap.String("contractAddress", nft.ContractAddress),
						zap.Any("image", image),
						zap.Any("traits", traits),
						zap.Any("animationFileUrl", animationFileUrl),
					)
					continue
				}

				logger.AtLog.Logger.Info(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit),
					zap.String("tokenID", nft.TokenID),
					zap.String("contractAddress", nft.ContractAddress),
					zap.Any("image", image),
					zap.Any("traits", traits),
					zap.Any("animationFileUrl", animationFileUrl),
					zap.Any("updated", updated),
				)
			} else {
				logger.AtLog.Logger.Info(fmt.Sprintf("SoulNftImageCrontab - page: %d, limit: %d", page, limit),
					zap.String("tokenID", nft.TokenID),
					zap.String("contractAddress", nft.ContractAddress),
					zap.Any("image", image),
					zap.Any("traits", traits),
					zap.Any("animationFileUrl", animationFileUrl),
				)
			}
		}

		wg1.Wait()
		wg2.Wait()

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

	newImagePath, traits := u.ParseHtmlImage(*inChan.AnimationFileUrl)
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
