package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (u *Usecase) FilterMKListing(ctx context.Context, filter entity.FilterMarketplaceListings) ([]entity.MarketplaceListings, error) {
	return u.Repo.FilterMarketplaceListings(filter)
}

func (u *Usecase) FilterMKOffers(ctx context.Context, filter entity.FilterMarketplaceOffer) ([]entity.MarketplaceOffers, error) {
	return u.Repo.FilterMarketplaceOffer(filter)
}

func (u *Usecase) FilterTokenActivities(ctx context.Context, filter entity.FilterTokenActivities) ([]entity.MarketplaceTokenActivity, error) {
	return u.Repo.FilterTokenActivites(filter)
}

func (u *Usecase) FilterMkplaceNfts(ctx context.Context, filter entity.FilterNfts) ([]*nft_explorer.MkpNftsResp, error) {
	resp := []*nft_explorer.MkpNftsResp{}
	f := bson.D{}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		f = append(f, bson.E{"collection_address", *filter.ContractAddress})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		f = append(f, bson.E{"token_id", *filter.TokenID})
	}
	if len(filter.AttrKey) > 0 {
		f = append(f, bson.E{"attributes.trait_type", bson.M{"$in": filter.AttrKey}})
	}

	if len(filter.AttrValue) > 0 {
		f = append(f, bson.E{"attributes.value", bson.M{"$in": filter.AttrValue}})
	}

	if filter.Rarity != nil {
		f = append(f, bson.E{"$and", bson.A{
			bson.E{"attributes.percent", bson.M{"$lte": filter.Rarity.Max / 100}},
			bson.E{"attributes.percent", bson.M{"$gte": filter.Rarity.Min / 100}},
		}})
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	sortBy := "token_id_int"
	if filter.SortBy != "" {
		sortBy = filter.SortBy
	}

	sort := -1
	if filter.Sort != 0 {
		sort = int(filter.Sort)
	}

	s := bson.D{{sortBy, sort}}
	err := u.Repo.Find(utils.VIEW_MARKETPLACE_NFTS, f, int64(filter.Limit), int64(filter.Offset), &resp, s)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) GetMkplaceNft(ctx context.Context, contractAddress string, tokenID string) (*nft_explorer.MkpNftsResp, error) {
	resp := &nft_explorer.MkpNftsResp{}
	f := bson.D{
		bson.E{"collection_address", contractAddress},
		bson.E{"token_id", tokenID},
	}

	cursor, err := u.Repo.FindOne(utils.VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES, f)
	if err != nil {
		return nil, err
	}

	err = cursor.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
