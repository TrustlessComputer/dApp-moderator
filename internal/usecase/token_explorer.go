package usecase

import (
	"context"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) FindTokens(ctx context.Context, filter request.PaginationReq, key string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.TokenFilter{}
	query.FromPagination(filter)

	if key != "" {
		query.Key = key
	}

	data, err = u.Repo.FindTokens(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("Tokens", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) FindToken(ctx context.Context, address string) (interface{}, error) {

	query := entity.TokenFilter{
		Address: address,
	}
	data, err := u.Repo.FindToken(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("Token", zap.Error(err))
		return nil, err
	}

	return data, nil
}

func (u *Usecase) UpdateToken(ctx context.Context, address string, req request.UpdateTokenReq) error {
	query := entity.TokenFilter{
		Address: address,
	}
	token, err := u.Repo.FindToken(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("Token", zap.Error(err))
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("token not found")
		}
		return err
	}

	token.Name = req.Name
	token.Symbol = req.Symbol
	token.Slug = strings.ToLower(req.Symbol)
	token.Description = req.Description

	token.Social.DisCord = req.Social.DisCord
	token.Social.Telegram = req.Social.Telegram
	token.Social.Twitter = req.Social.Twitter
	token.Social.Website = req.Social.Website
	token.Social.Medium = req.Social.Medium
	token.Social.Instagram = req.Social.Instagram

	err = u.Repo.UpdateToken(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) CrawToken(ctx context.Context, fromPage int) (int, error) {
	perPage := 100
	toPage := fromPage

	tokenCount := 1
	for tokenCount > 0 {

		offset := perPage * (toPage - 1)
		params := request.PaginationReq{
			Page:   &toPage,
			Limit:  &perPage,
			Offset: &offset,
		}.ToNFTServiceUrlQuery()
		Tokens, err := u.TokenExplorer.Tokens(params)
		if err != nil {
			logger.AtLog.Logger.Error("Tokens() failed", zap.Error(err))
			return toPage, err
		}
		tokenCount = len(Tokens)
		if tokenCount == 0 {
			return toPage, nil
		}

		for _, t := range Tokens {

			// parse token
			token := entity.Token{}
			if err = token.FromTokenExplorer(t); err != nil {
				logger.AtLog.Logger.Error("FromTokenExplorer() failed", zap.Error(err))
				return toPage, nil
			}

			// check if token exist
			dbToken, err := u.Repo.FindToken(ctx, entity.TokenFilter{
				Address: token.Address,
			})
			if err != nil && err != mongo.ErrNoDocuments {
				logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
				return toPage, nil
			}

			if dbToken != nil {
				continue
			}

			countInt := int64(0)
			count, _, err := u.Repo.CountDocuments(utils.COLLECTION_TOKENS, bson.D{})
			if err == nil && count != nil {
				countInt = *count
			}

			countInt++
			token.Index = countInt
			// save token to DB
			_, err = u.Repo.InsertOne(&token)
			if err != nil {
				logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
				return toPage, nil
			}
			u.NewTokenNotify(&token)
		}

		if len(Tokens) < perPage {
			break
		}

		toPage++

	}
	return toPage, nil
}

func (u *Usecase) FindWalletAddressTokens(ctx context.Context, filter request.PaginationReq, walletAddress string) (interface{}, error) {
	query := entity.TokenFilter{}
	query.FromPagination(filter)

	contractAddresses := []string{}
	contractAddressBalance := make(map[string]token_explorer.WalletAddressToken)

	data, err := u.TokenExplorer.WalletAddressTokens(walletAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("FindWalletAddressTokens", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	for _, item := range data {
		contractAddresses = append(contractAddresses, item.Contract)
		contractAddressBalance[strings.ToLower(item.Contract)] = item
	}

	tokens, err := u.Repo.FindTokensByContracts(ctx, contractAddresses)
	if err != nil {
		logger.AtLog.Logger.Error("FindWalletAddressTokens", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	resp := []*entity.OwnedToken{}

	for _, token := range tokens {
		tmp := contractAddressBalance[strings.ToLower(token.Address)]
		tmpResp := token.OwnedToken()
		tmpResp.Balance = tmp.Balance
		tmpResp.Decimal = tmp.Decimal

		resp = append(resp, tmpResp)
	}

	logger.AtLog.Logger.Info("FindWalletAddressTokens", zap.String("walletAddress", walletAddress), zap.Any("data", resp))
	return resp, nil
}
