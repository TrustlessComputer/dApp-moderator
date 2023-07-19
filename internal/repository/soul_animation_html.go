package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strings"
)

func (r *Repository) InsertSoulAnimationHtml(obj *entity.SoulAnimationHtml) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindAnimationHtml(contractAddress, tokenID string) (*entity.SoulAnimationHtml, error) {
	resp := &entity.SoulAnimationHtml{}

	f := bson.D{
		{"collection_address", strings.ToLower(contractAddress)},
		{"token_id", strings.ToLower(contractAddress)},
	}

	r1, err := r.FindOne(utils.COLLECTION_SOUL_ANIMATION_HTML, f)
	if err != nil {
		return nil, err
	}

	err = r1.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (r *Repository) FindAnimationHtmls(tokenIDs []string) ([]entity.SoulAnimationHtml, error) {
	resp := &[]entity.SoulAnimationHtml{}

	f := bson.D{
		{"token_id", bson.M{"$in": tokenIDs}},
		{"collection_address", strings.ToLower(os.Getenv("SOUL_CONTRACT"))},
	}

	r1, err := r.DB.Collection(utils.COLLECTION_SOUL_ANIMATION_HTML).Find(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	err = r1.All(context.TODO(), resp)
	if err != nil {
		return nil, err
	}

	return *resp, err

}
