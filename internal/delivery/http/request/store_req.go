package request

type CreateDappInfoReq struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Link    string `json:"link"`
	Image   string `json:"image"`
	Creator string `json:"creator"`
}
