package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func (u *Usecase) FilterMKListing(ctx context.Context, filter entity.FilterMarketplaceListings) ([]entity.MarketplaceListings, error) {
	return u.Repo.FilterMarketplaceListings(filter)
}

func (u *Usecase) FilterMKOffers(ctx context.Context, filter entity.FilterMarketplaceOffer) ([]entity.MarketplaceOffers, error) {
	return u.Repo.FilterMarketplaceOffer(filter)
}

func (u *Usecase) FilterTokenActivities(ctx context.Context, filter entity.FilterTokenActivities) ([]*entity.MarketplaceTokenActivity, error) {
	return u.Repo.FilterTokenActivites(filter)
}

func (u *Usecase) FilterMkplaceNfts(ctx context.Context, filter entity.FilterNfts) ([]*nft_explorer.MkpNftsResp, error) {
	resp := []*nft_explorer.MkpNftsResp{}
	f := bson.D{}

	maxFileSize := os.Getenv("FILE_CHUNK_SIZE")
	if filter.IsBigFile != nil {
		maxFileSizeInt, _ := strconv.Atoi(maxFileSize)
		if *filter.IsBigFile == true {
			f = append(f, bson.E{"size", bson.M{"$gte": maxFileSizeInt}})
		} else {
			f = append(f, bson.E{"size", bson.M{"$lt": maxFileSizeInt}})
		}
	}

	if filter.IsBuyable != nil {
		f = append(f, bson.E{"buyable", *filter.IsBuyable})
	}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		f = append(f, bson.E{"collection_address", *filter.ContractAddress})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		f = append(f, bson.E{"token_id", *filter.TokenID})
	}

	if filter.Rarity != nil {
		filter.Rarity.Min = filter.Rarity.Min / 100
		filter.Rarity.Max = filter.Rarity.Max / 100
		//f = append(f, bson.E{"$and", bson.A{
		//	bson.E{"attributes.percent", bson.M{"$lte": filter.Rarity.Max / 100}},
		//	bson.E{"attributes.percent", bson.M{"$gte": filter.Rarity.Min / 100}},
		//}})

		attrs, err := u.Repo.FilterCollectionAttributeByPercent(entity.FilterMarketplaceCollectionAttribute{
			ContractAddress: filter.ContractAddress,
			MaxPercent:      &filter.Rarity.Max,
			MinPercent:      &filter.Rarity.Min,
		})

		if err != nil {
			return nil, err
		}

		key := []string{}
		value := []string{}
		for _, attr := range attrs {
			key = append(key, attr.TraitType)
			value = append(value, attr.Value)
		}

		filter.AttrKey = key
		filter.AttrValue = value
	}

	if filter.Price != nil {
		btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))
		ethRate := u.GetExternalRate(os.Getenv("WETH_ADDRESS"))
		rate := btcRate / ethRate

		minPrice := filter.Price.Min
		maxPrice := filter.Price.Max

		minPriceEth := minPrice * rate
		maxPriceEth := maxPrice * rate

		fPrice := bson.A{
			bson.D{
				{"$and",
					bson.A{
						bson.D{{"erc20", strings.ToLower(os.Getenv("WBTC_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gt", minPrice}}}},
						bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
					},
				},
			},
			bson.D{
				{"$and",
					bson.A{
						bson.D{{"erc20", strings.ToLower(os.Getenv("WETH_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gt", minPriceEth}}}},
						bson.D{{"price", bson.D{{"$lte", maxPriceEth}}}},
					},
				},
			},
		}

		f = append(f, bson.E{"$or", fPrice})

	}

	if len(filter.AttrKey) > 0 {
		f = append(f, bson.E{"attributes.trait_type", bson.M{"$in": filter.AttrKey}})
	}

	if len(filter.AttrValue) > 0 {
		f = append(f, bson.E{"attributes.value", bson.M{"$in": filter.AttrValue}})
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

	s := bson.D{
		{"buyable", -1},
		{"price_erc20.price", 1},
		{sortBy, sort},
	}
	//old: VIEW_MARKETPLACE_NFTS, VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES has attributes + percent

	projections := bson.D{
		{"activities", 0},
	}

	err := u.Repo.FindWithProjections(utils.VIEW_NEW_MARKETPLACE_NFTS, f, int64(filter.Limit), int64(filter.Offset), &resp, s, projections)
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

func (u *Usecase) FilterCollectionChart(ctx context.Context, filter entity.FilterCollectionChart) ([]*entity.CollectionChart, error) {
	nfts, err := u.Repo.GetCollectionChart(filter)
	if err != nil {
		return nil, err
	}

	//calculate rate
	for _, nft := range nfts {
		bigI, _ := new(big.Int).SetString(nft.Price, 10)

		e := &entity.MarketPlaceVolume{
			TotalVolume: float64(bigI.Int64()),
			Erc20Token:  nft.Erc20Token,
		}

		u.calculateRate(e)
		nft.USDT = e.USDTValue
		nft.USDTRate = e.Erc20Rate

	}

	//group and calculate usdt of a day
	groupdata := make(map[string]entity.CollectionChart)
	for _, nft := range nfts {
		value, ok := groupdata[nft.VolumeCreatedAtDate]
		if !ok {
			groupdata[nft.VolumeCreatedAtDate] = *nft
		} else {
			nft.USDT = value.USDT + nft.USDT
			nft.USDTRate = value.USDTRate + nft.USDTRate
			groupdata[nft.VolumeCreatedAtDate] = *nft
		}

	}

	//response data
	resp := []*entity.CollectionChart{}
	for _, item := range groupdata {
		resp = append(resp, &item)
	}

	return resp, nil
}
