package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) InsertUploadedFileChunk(obj *entity.UploadedFileChunk) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListChunks(filter *entity.FilterChunks) ([]entity.UploadedFileChunk, error) {
	match := bson.D{}
	chunks := []entity.UploadedFileChunk{}
	if filter.Status != nil {
		match = append(match, bson.E{"status", *filter.Status})
	}

	if filter.FileID != nil {
		match = append(match, bson.E{"file_id", *filter.FileID})
	}

	f := bson.A{
		bson.D{
			{"$match", match},
		},
		bson.D{{"$sort", bson.D{{"chunk_index", 1}}}},
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	if filter.Offset == 0 {
		filter.Offset = 0
	}

	f = append(f, bson.D{{"$skip", filter.Offset}})
	f = append(f, bson.D{{"$limit", filter.Limit}})

	cursor, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILE_CHUNKS).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &chunks)
	if err != nil {
		return nil, err
	}

	return chunks, nil
}

func (r *Repository) ListUploadedFiles(filter *entity.FilterUploadedFile) ([]entity.UploadedFile, error) {
	match := bson.D{}
	files := []entity.UploadedFile{}
	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"contract_address", *filter.ContractAddress})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"token_id", *filter.TokenID})
	}

	if filter.WalletAddress != nil && *filter.WalletAddress != "" {
		match = append(match, bson.E{"wallet_address", *filter.WalletAddress})
	}

	if filter.TxHash != nil && *filter.TxHash != "" {
		match = append(match, bson.E{"tx_hash", *filter.TxHash})
	}

	f := bson.A{
		bson.D{
			{"$match", match},
		},
		bson.D{{"$sort", bson.D{{"chunk_index", 1}}}},
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	if filter.Offset == 0 {
		filter.Offset = 0
	}

	f = append(f, bson.D{{"$skip", filter.Offset}})
	f = append(f, bson.D{{"$limit", filter.Limit}})

	cursor, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILES).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) UpdateChunkTxHash(fileID string, chunkID string, txHash string) error {
	cID, err := primitive.ObjectIDFromHex(chunkID)
	if err != nil {
		return err
	}

	fID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":     cID,
		"file_id": fID,
	}

	update := bson.M{
		"tx_hash": txHash,
	}

	result, err := r.DB.Collection(utils.COLLECTION_UPLOADED_FILE_CHUNKS).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r *Repository) FindChunk(fileID string, chunkID string) (*entity.UploadedFileChunk, error) {
	cID, err := primitive.ObjectIDFromHex(chunkID)
	if err != nil {
		return nil, err
	}

	fID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"_id", cID},
		{"file_id", fID},
	}

	result, err := r.FindOne(utils.COLLECTION_UPLOADED_FILE_CHUNKS, filter)
	if err != nil {
		return nil, err
	}

	resp := &entity.UploadedFileChunk{}
	err = result.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
