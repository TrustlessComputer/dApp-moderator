package repository

import (
	"dapp-moderator/internal/entity"
)

func (r Repository) CreateTokenUriHistory(data *entity.TokenUriHistories) error {
	err := r.InsertOne(data.TableName(), data)
	if err != nil {
		return err
	}
	return nil
}
