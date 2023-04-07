package usecase

import (
	"context"
	"crypto/rand"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u Usecase) GenerateMessage(ctx context.Context, data *structure.GenerateMessage) (*string, error) {
	addrr := data.Address
	addrr = strings.ToLower(addrr)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
		return nil, err
	}
	message := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	message = fmt.Sprintf(utils.NONCE_MESSAGE_FORMAT, message)
	now := time.Now().UTC()
	_, err = u.Repo.FindUserByWalletAddress(addrr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//insert
			user := &entity.Users{}
			user.WalletType = data.WalletType
			user.WalletAddress = addrr
			user.Message = message
			user.CreatedAt = &now

			_, err = u.Repo.InsertOne(user)
			if err != nil {
				logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
				return nil, err
			}

			return &message, nil

		} else {
			logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
			return nil, err
		}
	}
	
	_, err = u.Repo.UpdateUserMessage(addrr, message)
	if err != nil {
		logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
		return nil, err
	}

	return &message, nil
}