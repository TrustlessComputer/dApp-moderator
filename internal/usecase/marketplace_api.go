package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	soul_contract "dapp-moderator/utils/contracts/soul"
	"dapp-moderator/utils/helpers"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
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

func (u *Usecase) FilterTokenSoulHistories(ctx context.Context, filter entity.FilterTokenActivities) ([]*entity.SoulTokenHistoriesFiltered, error) {
	return u.Repo.FilterSoulHistories(filter)
}

func (u *Usecase) FilterMkplaceNfts(ctx context.Context, filter entity.FilterNfts) (*entity.MkpNftsPagination, error) {
	resp := []*entity.MkpNftsResp{}
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
		f = append(f, bson.E{"collection_address", strings.ToLower(*filter.ContractAddress)})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		f = append(f, bson.E{"token_id", *filter.TokenID})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		f = append(f, bson.E{"owner", strings.ToLower(*filter.Owner)})
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
		//btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))
		//ethRate := u.GetExternalRate(os.Getenv("WETH_ADDRESS"))
		//rate := btcRate / ethRate

		minPrice := filter.Price.Min
		maxPrice := filter.Price.Max

		//minPriceEth := minPrice * rate
		//maxPriceEth := maxPrice * rate

		fPrice := bson.A{
			bson.D{
				{"$and",
					bson.A{
						//bson.D{{"erc20", strings.ToLower(os.Getenv("WBTC_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gte", minPrice}}}},
						bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
					},
				},
			},
			//bson.D{
			//	{"$and",
			//		bson.A{
			//			//bson.D{{"erc20", strings.ToLower(os.Getenv("WETH_ADDRESS"))}},
			//			bson.D{{"price", bson.D{{"$gte", minPriceEth}}}},
			//			bson.D{{"price", bson.D{{"$lte", maxPriceEth}}}},
			//		},
			//	},
			//},
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
		{"price", 1},
		{sortBy, sort},
	}
	//old: VIEW_MARKETPLACE_NFTS, VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES has attributes + percent

	projections := bson.D{
		{"activities", 0},
	}

	queryFromView := utils.VIEW_NEW_MARKETPLACE_NFTS
	if filter.ContractAddress != nil && strings.ToLower(*filter.ContractAddress) == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
		queryFromView = utils.VIEW_SOUL_MARKETPLACE_NFTS_AUCTION_RARITY
	}

	err := u.Repo.FindWithProjections(queryFromView, f, filter.Limit, filter.Offset, &resp, s, projections)
	if err != nil {
		return nil, err
	}

	total, err := u.Repo.AllItems(queryFromView, f)
	if err != nil {
		return nil, err
	}

	for index, item := range resp {
		if len(item.BnsDefault) > 0 && item.BnsDefault[0].Resolver != "" {
			for j, bnsItem := range resp[index].BnsData {
				if bnsItem.ID.Hex() == item.BnsDefault[0].BNSDefaultID.Hex() {
					resp[index].BnsData[0], resp[index].BnsData[j] = resp[index].BnsData[j], resp[index].BnsData[0]
					break
				}
			}
		}
	}

	respData := &entity.MkpNftsPagination{
		Items:     resp,
		TotalItem: total,
	}
	return respData, nil
}

func (u *Usecase) FilterMkplaceNftNew(ctx context.Context, filter entity.FilterNfts) (*entity.MkpNftsPagination, error) {
	resp, err := u.Repo.FilterMKPNfts(filter)
	if err != nil {
		return nil, err
	}

	// Nếu contract là SOUL thì lấy name tu chain, mặc dù ban đầu có update vô rồi nhưng van co thể user change name
	if filter.ContractAddress != nil && strings.ToLower(*filter.ContractAddress) == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
		if soulContract, err := soul_contract.NewSoul(common.HexToAddress(*filter.ContractAddress), u.TCPublicNode.GetClient()); err == nil {
			for i, item := range resp.Items {
				resp.Items[i].Name = ""
				if name, err := u.SoulNFTName(item.TokenID, soulContract); err == nil {
					resp.Items[i].Name = name
				}
			}
		}
	}

	return resp, nil
}

func (u *Usecase) GetMkplaceNft(ctx context.Context, contractAddress string, tokenID string) (*nft_explorer.MkpNftsResp, error) {
	resp := &nft_explorer.MkpNftsResp{}

	f := bson.D{
		bson.E{"collection_address", strings.ToLower(contractAddress)},
		bson.E{"token_id", tokenID},
	}

	collectionName := utils.VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES
	if strings.ToLower(contractAddress) == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
		collectionName = utils.VIEW_SOUL_NFT_WITH_ATTRIBUTES
	}

	cursor, err := u.Repo.FindOne(collectionName, f)
	if err != nil {
		return nil, err
	}

	err = cursor.Decode(resp)
	if err != nil {
		return nil, err
	}

	bnsData, err := u.Repo.FilterBNS(entity.FilterBns{
		BaseFilters: entity.BaseFilters{
			SortBy: "_id",
			Sort:   entity.SORT_ASC,
		},
		Resolver: utils.ToPtr(resp.Owner),
	})
	if err == nil && len(bnsData) > 0 {
		resp.BnsData = bnsData
	}

	if strings.ToLower(contractAddress) == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
		if name, err := u.SoulNFTName(tokenID); err == nil {
			resp.Name = name
		} else {
			resp.Name = ""
		}
	}

	if strings.ToLower(contractAddress) == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
		attributes := resp.Attributes
		attrSorted := make(map[string]nft_explorer.MkpNftAttr)
		attrSorted[strings.ToLower("Color Palette")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Sea Level")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Neighborhood")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Soul Form")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Decoration")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Special Object")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Rendering Rate")] = nft_explorer.MkpNftAttr{}
		attrSorted[strings.ToLower("Activity In")] = nft_explorer.MkpNftAttr{}

		for _, attribute := range attributes {
			attrSorted[strings.ToLower(attribute.TraitType)] = attribute
		}

		respAttr := []nft_explorer.MkpNftAttr{}
		respAttr = append(respAttr, attrSorted[strings.ToLower("Color Palette")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Sea Level")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Neighborhood")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Soul Form")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Decoration")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Special Object")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Rendering Rate")])
		respAttr = append(respAttr, attrSorted[strings.ToLower("Activity In")])

		resp.Attributes = respAttr
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
			newUSDT := value.USDT + nft.USDT
			nft.USDT = newUSDT

			//nft.USDTRate = value.USDTRate + nft.USDTRate
			groupdata[nft.VolumeCreatedAtDate] = *nft
		}

	}

	btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))

	//response data
	resp := []*entity.CollectionChart{}
	for _, item := range groupdata {
		btc := helpers.ConvertAmount(item.USDT / btcRate)
		btcI, _ := btc.Int64()
		item.BTC = fmt.Sprintf("%d", btcI)
		resp = append(resp, &item)
	}

	return resp, nil
}

func (u *Usecase) FilterNftOwners(ctx context.Context, filter entity.FilterCollectionNftOwners) (*entity.CollectionNftOwnerFiltered, error) {
	owners, err := u.Repo.CollectionNftOwner(filter)
	if err != nil {
		return nil, err
	}
	return owners, nil
}
