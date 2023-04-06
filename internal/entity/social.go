package entity

type Social struct {
	Website   string `json:"website" bson:"website"`
	DisCord   string `json:"discord" bson:"discord"`
	Twitter   string `json:"twitter" bson:"twitter"`
	Telegram  string `json:"telegram" bson:"telegram"`
	Medium    string `json:"medium" bson:"medium"`
	Instagram string `json:"instagram" bson:"instagram"`
}
