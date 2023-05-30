package entity

import "dapp-moderator/utils"

type FilterUploadedFile struct {
	BaseFilters
	TokenID         *string
	ContractAddress *string
	WalletAddress   *string
	Name            *string
	TxHash          *string
	Status          []int
}

type QueriedUploadedFile struct {
	UploadedFile   `bson:",inline"`
	ProcessedChunk int `bson:"processed_chunk" json:"processed_chunks"`
	Status         int `bson:"status" json:"status"`
}

type UploadedFile struct {
	BaseEntity      `bson:",inline"`
	Name            string `bson:"name" json:"name"`
	Path            string `bson:"path" json:"path"`
	FullPath        string `bson:"full_path" json:"full_path"`
	FileType        string `bson:"file_type" json:"file_type"`
	Size            int    `bson:"size" json:"size"` //kb
	Chunks          int    `bson:"chunks" json:"total_chunks"`
	ChunkSize       int    `bson:"chunk_size" json:"chunk_size"` //kb
	TxHash          string `bson:"tx_hash" json:"tx_hash"`
	TokenID         string `bson:"token_id" json:"token_id"`
	ContractAddress string `bson:"contract_address" json:"contract_address"`
	WalletAddress   string `bson:"wallet_address" json:"wallet_address"`
}

func (t *UploadedFile) CollectionName() string {
	return utils.COLLECTION_UPLOADED_FILES
}
