package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// UserCredits godoc
// @Summary Get Proposals
// @Description Get Proposals
// @Tags DAO - Proposal
// @Accept  json
// @Produce  json
// @Param proposer query string false "proposer"
// @Param proposal_id query string false "proposal_id"
// @Param contract_address query string false "contract_address"
// @Param states query string false "separated by comma"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by field: default volume"
// @Param sort query int false "sort default: -1 desc"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /dao/proposals [GET]
func (h *httpDelivery) proposals(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			if p.SortBy == nil {
				sortBy := "created_at"
				p.SortBy = &sortBy
			}
			if p.Sort == nil {
				s := int(entity.SORT_DESC)
				p.Sort = &s
			}

			f := entity.FilterProposal{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}

			proposer := r.URL.Query().Get("proposer")
			proposalID := r.URL.Query().Get("proposal_id")
			states := r.URL.Query().Get("states")
			contractAddress := r.URL.Query().Get("contract_address")

			if proposer != "" {
				f.Proposer = &proposer
			}
			if contractAddress != "" {
				f.ContractAddress = &contractAddress
			}
			if proposalID != "" {
				f.ProposalID = &proposalID
			}
			if states != "" {
				s := []int{}
				states = strings.ReplaceAll(states, " ", "")
				stateArr := strings.Split(states, ",")
				for _, stateStr := range stateArr {
					stateINT, err := strconv.Atoi(stateStr)
					if err == nil {
						s = append(s, stateINT)
					}
				}
				f.State = s
			}

			return h.Usecase.GetProposals(f)

		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Create Proposals
// @Description Create Proposals
// @Tags DAO - Proposal
// @Accept  json
// @Produce  json
// @Param requestBody body entity.ProposalDetail false "requestBody"
// @Success 200 {object} response.JsonResponse{}
// @Router /dao/proposals [POST]
func (h *httpDelivery) createDraftProposals(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			reqBody := &entity.ProposalDetail{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			return h.Usecase.CreateDraftProposal(reqBody)
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Proposal
// @Description Get Proposal
// @Tags DAO - Proposal
// @Accept  json
// @Produce  json
// @Param proposal_id path string true "proposal_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /dao/proposals/{proposal_id} [GET]
func (h *httpDelivery) getProposal(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			proposalID := vars["proposal_id"]
			return h.Usecase.GetProposal(proposalID)
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Proposal's votes
// @Description Get Proposal's votes
// @Tags DAO - Proposal
// @Accept  json
// @Produce  json
// @Param contract_address query string false "contract_address"
// @Param voter query string false "voter"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by field: default volume"
// @Param sort query int false "sort default: -1 desc"
// @Param page query int false "page"
// @Param proposal_id path string true "proposal_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /dao/proposals/{proposal_id}/votes [GET]
func (h *httpDelivery) getProposalVotes(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			proposalID := vars["proposal_id"]

			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			if p.SortBy == nil {
				sortBy := "created_at"
				p.SortBy = &sortBy
			}
			if p.Sort == nil {
				s := int(entity.SORT_DESC)
				p.Sort = &s
			}

			f := entity.FilterProposalVotes{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}
			f.ProposalID = &proposalID

			contractAddress := r.URL.Query().Get("contract_address")
			voter := r.URL.Query().Get("voter")

			if contractAddress != "" {
				f.ContractAddress = &contractAddress
			}

			if voter != "" {
				f.Voter = &voter
			}

			return h.Usecase.GetProposalVotes(f)
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Update  Proposal's proposalID
// @Description Get Proposal's proposalID
// @Tags DAO - Proposal
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param proposal_id path string true "proposal_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /dao/proposals/{id}/{proposal_id} [PUT]
func (h *httpDelivery) mapOffAndOnChainProposal(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iD := vars["id"]
			proposalID := vars["proposal_id"]
			return h.Usecase.MapOffToOnChainProposal(iD, proposalID)
		},
	).ServeHTTP(w, r)
}
