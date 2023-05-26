package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) InsertUploadedFile(obj *entity.UploadedFile) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTxHashForUploadedFile(fileID string, txHash string) error {
	pID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": pID,
	}

	update := bson.M{
		"tx_hash": txHash,
	}

	result, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILES).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r *Repository) UpdateChunksForUploadedFile(fileID string, chunks int, chunkSize int) error {
	pID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": pID,
	}

	update := bson.M{
		"chunks":     chunks,
		"chunk_size": chunkSize,
	}

	result, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILES).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r *Repository) UpdateChunksTxHashForUploadedFile(fileID string, txHash string, walletAddress string, chunks int, chunkSize int) error {
	pID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": pID,
	}

	update := bson.M{
		"tx_hash":        txHash,
		"chunks":         chunks,
		"chunk_size":     chunkSize,
		"wallet_address": walletAddress,
	}

	result, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILES).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r *Repository) GetUploadedFile(fileID string) (*entity.UploadedFile, error) {
	pID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": pID,
	}

	result := r.DB.Collection(utils.COLLECTION_UPLOADED_FILES).FindOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	resp := &entity.UploadedFile{}
	err = result.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
