package entity

type DappInfo struct {
	BaseEntity  `bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Image       string `json:"image" bson:"image"`
	Link        string `json:"link" bson:"link"`
	Creator     string `json:"creator" bson:"creator"`
	Description string `json:"desc" bson:"description"`
	Status      int    `json:"status" bson:"status"`
}

func (u DappInfo) CollectionName() string {
	return "dapp_infos"
}
