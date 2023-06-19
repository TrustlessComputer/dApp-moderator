package request

type UpdateBNSDefaultRequest struct {
	TokenID  string `json:"token_id"`
	Resolver string `json:"-"`
}
