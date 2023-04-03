package repository

import (
	"dapp-moderator/internal/entity"
)

func (r Repository) CreateTokenUriMetadata(data *entity.TokenUriMetadata) error {
	err := r.InsertOne(data.TableName(), data)
	if err != nil {
		return err
	}
	return nil
}
