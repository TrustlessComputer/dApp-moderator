package entity

type Bns struct {
	Name    string `json:"name" bson:"name"`
	TokenId string `json:"tokenId" bson:"token_id"`
	Owner   string `json:"owner" bson:"owner"`
}
