package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (c *Usecase) Tokens(ctx context.Context, filter request.PaginationReq, key string) (interface{}, error) {
	var data interface{}
	var err error

	params := filter.ToNFTServiceUrlQuery()

	query := entity.TokenFilter{}
	query.FromPagination(filter)

	if key != "" {
		query.Key = key
	}

	if key == "" {
		data, err = c.TokenExplorer.Tokens(params)
	} else {
		params["query"] = []string{key}
		data, err = c.TokenExplorer.Search(params)
	}

	if err != nil {
		logger.AtLog.Logger.Error("Tokens", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) Token(ctx context.Context, address string) (interface{}, error) {

	data, err := c.TokenExplorer.Token(address)
	if err != nil {
		logger.AtLog.Logger.Error("Token", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Token", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) CrawToken(ctx context.Context, fromPage int) (int, error) {
	perPage := 100
	toPage := fromPage

	tokenCount := 1
	for tokenCount > 0 {
		params := request.PaginationReq{
			Page:  &toPage,
			Limit: &perPage,
		}.ToNFTServiceUrlQuery()
		Tokens, err := c.TokenExplorer.Tokens(params)
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
			dbToken, err := c.Repo.FindToken(ctx, entity.TokenFilter{
				Address: token.Address,
			})
			if err != nil && err != mongo.ErrNoDocuments {
				logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
				return toPage, nil
			}

			if dbToken != nil {
				continue
			}

			// save token to DB
			_, err = c.Repo.InsertOne(&token)
			if err != nil {
				logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
				return toPage, nil
			}
		}
		toPage++
	}
	return toPage, nil
}
