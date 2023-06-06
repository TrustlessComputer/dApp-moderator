package entity

import (
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChunkStatus int

const (
	ChunkNew       ChunkStatus = 0 // chunk is created
	ChunkUploading ChunkStatus = 1 //uploading to blockchain
	ChunkUploaded  ChunkStatus = 2 //uploaded to blockchain
)

type FilterChunks struct {
	FileID *primitive.ObjectID
	Status *ChunkStatus
	BaseFilters
}

type UploadedFileChunk struct {
	BaseEntity `bson:",inline"`
	FileID     primitive.ObjectID `json:"file_id" bson:"file_id"` //ref to uploaded file's ID
	ChunkIndex int                `json:"chunk_index" bson:"chunk_index"`
	ChunkData  []byte             `json:"chunk_data" bson:"chunk_data"`
	Status     ChunkStatus        `json:"status" bson:"status"`
	TxHash     string             `json:"tx_hash" bson:"tx_hash"`
}

func (t *UploadedFileChunk) CollectionName() string {
	return utils.COLLECTION_UPLOADED_FILE_CHUNKS
}
