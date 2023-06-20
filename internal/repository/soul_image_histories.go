package repository

import "dapp-moderator/internal/entity"

func (r *Repository) InsertSoulImageHistory(obj *entity.SoulImageHistories) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}
