package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) SwapAddOrUpdateIdo(ctx context.Context, idoReq *request.IdoRequest) (interface{}, error) {
	var err error
	user, err := u.Repo.FindUserByWalletAddress(idoReq.UserWalletAddress)
	if err != nil {
		err := errors.New("User is not exist")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}
	if user == nil {
		err := errors.New("User is not exist")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}

	tokenFilter := entity.TokenFilter{}
	tokenFilter.Address = idoReq.TokenAddress
	tokenFilter.CreatedBy = idoReq.UserWalletAddress

	token, err := u.Repo.FindToken(ctx, tokenFilter)
	if err != nil {
		err := errors.New("Token is not exist")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}
	if token == nil {
		err := errors.New("Token is not exist")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}

	isVeried, err := u.verify(idoReq.Signature, user.WalletAddress, token.Address)
	if err != nil {
		logger.AtLog.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}
	if !isVeried {
		err := errors.New("Signature is not valid")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}

	var ido *entity.SwapIdo
	if idoReq.ID == "" {
		query := entity.SwapIdoFilter{}
		query.Address = idoReq.TokenAddress
		query.WalletAddress = strings.ToLower(idoReq.UserWalletAddress)
		ido, err = u.Repo.FindSwapIdo(ctx, query)
		if err != nil && err != mongo.ErrNoDocuments {
			logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
			return nil, err
		}

		if ido != nil {
			err = errors.New("This IDO is existed")
			logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
			return nil, err
		}
		ido = &entity.SwapIdo{}
		ido.UserWalletAddress = strings.ToLower(idoReq.UserWalletAddress)
	} else {
		query := entity.SwapIdoFilter{}
		query.ID = idoReq.ID
		ido, err = u.Repo.FindSwapIdo(ctx, query)
		if err != nil {
			logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
			return nil, err
		}
	}
	ido.Discord = idoReq.Discord
	ido.Link = idoReq.Link
	ido.Price = idoReq.Price
	ido.StartAt = idoReq.StartAt
	ido.Token = *token
	ido.Twitter = idoReq.Twitter
	ido.Website = idoReq.Website
	ido.WhitePaper = idoReq.WhitePaper

	if idoReq.ID == "" {
		_, err = u.Repo.InsertOne(ido)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return nil, err
		}
	} else {
		err = u.Repo.UpdateSwapIdo(ctx, ido)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return nil, err
		}
	}

	logger.AtLog.Logger.Info("SwapAddIdo", zap.Any("data", ido))
	return ido, nil
}

func (u *Usecase) SwapFindSwapIdoHistories(ctx context.Context, filter request.PaginationReq) (interface{}, error) {
	var err error
	query := entity.SwapIdoFilter{}
	query.FromPagination(filter)

	idos, err := u.Repo.FindSwapIdos(ctx, query)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("SwapFindSwapIdoHistories", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("SwapFindSwapIdoHistories", zap.Any("data", idos))
	return idos, nil
}

func (u *Usecase) SwapFindSwapIdoDetail(ctx context.Context, id string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapIdoFilter{}
	query.ID = id

	data, err = u.Repo.FindSwapIdo(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("SwapFindSwapIdoDetail", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("SwapFindSwapIdoDetail", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) SwapDeleteSwapIdo(ctx context.Context, id string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapIdoFilter{}
	query.ID = id

	err = u.Repo.DetelteSwapIdo(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("SwapDeleteSwapIdo", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (u *Usecase) SwapFindTokens(ctx context.Context, filter request.PaginationReq, owner string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.TokenFilter{}
	query.FromPagination(filter)
	query.CreatedBy = owner

	data, err = u.Repo.FindIdoTokens(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("Tokens", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
	return data, nil
}
