package response

import "dapp-moderator/internal/entity"

type SwapRouteResponse struct {
	PathPairs  []*entity.SwapPair `json:"path_pairs"`
	PathTokens []*entity.Token    `json:"path_tokens"`
}
