package repository

import (
	"dapp-moderator/internal/entity"
)

func (r *Repository) InsertSoulImageHistory(obj *entity.SoulImageHistories) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RunAgregateFunction() error {
	//cursor, err := r.DB.RunCommand()
	//if err != nil {
	//	return nil
	//}
	//
	//_ = cursor
	return nil
}
