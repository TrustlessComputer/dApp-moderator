package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/erc20"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

type CheckGMBalanceOutputChan struct {
	Nft     entity.Nfts
	Err     error
	Balance *big.Int
}

func (u *Usecase) SoulCrontab() error {
	maxProcess := 10
	minBalance := float64(1)
	erc20Addr := "0x2fe8d5A64afFc1d703aECa8a566f5e9FaeE0C003"
	instance, err := erc20.NewErc20(common.HexToAddress(erc20Addr), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	collection, err := u.Repo.GetSoulCollection()
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	nfts, err := u.Repo.SoulNfts(collection.Contract)
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	inputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan CheckGMBalanceOutputChan, len(nfts))
	wg := sync.WaitGroup{}
	logger.AtLog.Logger.Info("SoulCrontab", zap.String("contract_address", collection.Contract), zap.Int("nfts", len(nfts)))

	for i := 0; i < len(nfts); i++ {
		go u.CheckGMBalanceWorker(&wg, instance, inputChan, outputChan)
	}

	for i, nft := range nfts {
		wg.Add(1)
		inputChan <- nft
		if i%maxProcess == 0 && i > 0 {
			wg.Wait()
		}
	}

	for i := 0; i < len(nfts); i++ {
		out := <-outputChan
		if out.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(out.Err))
			continue
		}

		tokenIDInt, err := strconv.Atoi(out.Nft.TokenID)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(err))
			continue
		}

		isAuction := false

		value := helpers.GetValue(fmt.Sprintf("%d", out.Balance.Int64()), 18)
		if value < minBalance {
			isAuction = true
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.String("owner", out.Nft.Owner), zap.String("balance", fmt.Sprintf("%d", out.Balance.Int64())))

		insertData := &entity.NftAuctions{
			TokenID:         out.Nft.TokenID,
			TokenIDInt:      int64(tokenIDInt),
			ContractAddress: strings.ToLower(out.Nft.ContractAddress),
			IsAuction:       isAuction,
		}

		err = u.Repo.InsertAuction(insertData)
	}
	return nil
}

func (u *Usecase) CheckGMBalanceWorker(wg *sync.WaitGroup, erc20Instance *erc20.Erc20, input chan entity.Nfts, output chan CheckGMBalanceOutputChan) {
	defer wg.Done()
	nft := <-input

	owner := nft.Owner
	balanceOf, err := erc20Instance.BalanceOf(nil, common.HexToAddress(owner))

	output <- CheckGMBalanceOutputChan{
		Nft:     nft,
		Balance: balanceOf,
		Err:     err,
	}
}

func (u *Usecase) FilterSoulNfts(ctx context.Context, filter entity.FilterNfts) ([]*nft_explorer.SoulNft, error) {
	resp := []*nft_explorer.SoulNft{}
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

	err := u.Repo.FindWithProjections(utils.VIEW_NFT_AUCTION, f, int64(filter.Limit), int64(filter.Offset), &resp, s, projections)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
