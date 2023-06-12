package usecase

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"go.uber.org/zap"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type collectionsChan struct {
	Items []entity.MarketplaceCollectionAggregation
	Err   error
	Done  bool
}

type outputMkpCollectionChan struct {
	Item *entity.MarketplaceCollectionAggregation
	Err  error
}

func (u *Usecase) LoadCollections(limit int64, page int64) ([]entity.MarketplaceCollectionAggregation, error) {

	offset := int64((page - 1) * limit)
	items, err := u.Repo.AggregatetMarketPlaceData(entity.FilterMarketplaceAggregationData{
		BaseFilters: entity.BaseFilters{
			Limit:  limit,
			Offset: offset,
		},
	})

	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("LoadCollections - page: %d, limit: %d, offset: %d", page, limit, offset), zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (u *Usecase) ProccessLoadingCollections(limit int64, page int64, dataChan chan collectionsChan) {
	var err error

	items := []entity.MarketplaceCollectionAggregation{}
	//output := &[]entity.MarketplaceCollectionAggregation{}

	defer func() {
		isDOne := false
		if err != nil || len(items) == 0 {
			isDOne = true
		}

		dataChan <- collectionsChan{
			Items: items,
			Err:   err,
			Done:  isDOne,
		}
	}()

	items, err = u.LoadCollections(limit, page)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("AggregateCollectionStats - page: %d, limit: %d ", page, limit), zap.Error(err))
		return
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("AggregateCollectionStats - page: %d, limit: %d ", page, limit), zap.Int("items", len(items)))
	if len(items) == 0 {
		return
	}
}

func (u *Usecase) calculateRate(volume *entity.MarketPlaceVolume) error {
	rate := float64(0)
	tokenAddress := strings.ToLower(volume.Erc20Token)
	var err error
	decimal := 18

	btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))
	ethRate := u.GetExternalRate(os.Getenv("WETH_ADDRESS"))

	volume.Erc20Rate = rate
	volume.Erc20Decimal = decimal
	volume.WBTCRate = btcRate
	volume.WEthRate = ethRate

	if volume.WBTCRate != 0 && volume.WEthRate != 0 {
		erc20Val := helpers.GetValue(fmt.Sprintf("%d", volume.TotalVolume), float64(volume.Erc20Decimal))

		if strings.ToLower(tokenAddress) == strings.ToLower(os.Getenv("WBTC_ADDRESS")) {
			rate = volume.WBTCRate
		}

		if strings.ToLower(tokenAddress) == strings.ToLower(os.Getenv("WETH_ADDRESS")) {
			rate = volume.WEthRate
		}

		volume.USDTValue = erc20Val * rate
		volume.BTCValue = volume.USDTValue / volume.WBTCRate
		volume.EthValue = volume.USDTValue / volume.WEthRate
	} else {
		volume.USDTValue = 0
		volume.BTCValue = 0
		volume.EthValue = 0
	}

	return err
}

func (u *Usecase) GetExternalRate(tokenAddress string) float64 {
	rate := float64(0)
	symbol := "BTC"
	if strings.ToLower(tokenAddress) == strings.ToLower(os.Getenv("WETH_ADDRESS")) {
		symbol = "ETH"
	}

	var err error
	key := helpers.TokenRateKey(tokenAddress)
	existed, _ := u.Cache.Exists(key)

	if existed != nil && *existed == false {
		rate, err = helpers.GetExternalPrice(symbol)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("StartWorker - GetExternalPrice -  %s", tokenAddress), zap.Error(err))
		}

		u.Cache.SetDataWithExpireTime(key, rate, 1200) // 20 min
	} else {
		cached, _ := u.Cache.GetData(key)
		if cached != nil {
			rate, err = strconv.ParseFloat(*cached, 10)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("StartWorker - ParseFloat -  %s", tokenAddress), zap.Error(err))
				rate = 0
			}
		}
	}

	return rate
}

func (u *Usecase) StartWorker(inputItemChan chan entity.MarketplaceCollectionAggregation, outputChan chan outputMkpCollectionChan) {
	inputItem := <-inputItemChan
	minNumber := float64(0)
	var err error

	//calculate usdt by erc20 token
	total := float64(0)
	totalBTC := float64(0)
	totalEth := float64(0)

	min := minNumber
	minBTC := minNumber
	minEth := minNumber
	volumes := inputItem.MarketPlaceVolumes
	floorPriceVolumes := inputItem.FloorPriceMarketPlaceVolumes

	for _, volume := range floorPriceVolumes {
		err = u.calculateRate(volume)
	}

	sort.SliceStable(floorPriceVolumes, func(i, j int) bool {
		return floorPriceVolumes[i].USDTValue < floorPriceVolumes[j].USDTValue
	})

	sort.SliceStable(floorPriceVolumes, func(i, j int) bool {
		return floorPriceVolumes[i].BTCValue < floorPriceVolumes[j].BTCValue
	})

	sort.SliceStable(floorPriceVolumes, func(i, j int) bool {
		return floorPriceVolumes[i].EthValue < floorPriceVolumes[j].EthValue
	})

	if len(floorPriceVolumes) >= 1 {
		min = floorPriceVolumes[0].USDTValue
	}

	if len(floorPriceVolumes) >= 1 {
		minBTC = floorPriceVolumes[0].BTCValue
	}

	if len(floorPriceVolumes) >= 1 {
		minEth = floorPriceVolumes[0].EthValue
	}

	for _, volume := range volumes {
		err = u.calculateRate(volume)
		total += volume.USDTValue
		totalBTC += volume.BTCValue
		totalEth += volume.EthValue
	}

	inputItem.Volume = total
	inputItem.BtcVolume = totalBTC
	inputItem.EthVolume = totalEth

	inputItem.FloorPrice = min
	inputItem.BtcFloorPrice = minBTC
	inputItem.EthFloorPrice = minEth

	outputChan <- outputMkpCollectionChan{
		Item: &inputItem,
		Err:  err,
	}
}

func (u *Usecase) AggregateCollectionStats() error {
	page := int64(1)
	limit := int64(100)

	data := []entity.MarketplaceCollectionAggregation{}
	dataChan := make(chan collectionsChan)
	stop := make(chan bool, 1)

	go func(stopSig chan bool) {
		for {
			go u.ProccessLoadingCollections(limit, page, dataChan)
			page++

			if <-stop {
				break
			}
		}
	}(stop)

	//collect data for collections
	for {
		dataFromChan := <-dataChan
		data = append(data, dataFromChan.Items...)
		if dataFromChan.Done {
			stop <- true
			break
		}
		stop <- false
	}

	logger.AtLog.Logger.Info("AggregateCollectionStats", zap.Int("collections", len(data)))

	//update USDT price - start worker
	inputItemChan := make(chan entity.MarketplaceCollectionAggregation, len(data))
	outputChan := make(chan outputMkpCollectionChan, len(data))

	for i := 0; i < len(data); i++ {
		go u.StartWorker(inputItemChan, outputChan)
	}

	//insert data to worker
	for i := 0; i < len(data); i++ {
		inputItemChan <- data[i]
		if i > 0 && i%100 == 0 {
			time.Sleep(time.Millisecond * 500)
		}
	}

	//get the processed data
	for i := 0; i < len(data); i++ {
		dataFromChan := <-outputChan
		if dataFromChan.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("AggregateCollectionStats %s ", dataFromChan.Item.Contract), zap.Error(dataFromChan.Err))
			return dataFromChan.Err
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("AggregateCollectionStats %s ", dataFromChan.Item.Contract), zap.Any("aggregated", dataFromChan.Item))

		//save to view
		err := u.Repo.InsertMarketPlaceAggregation(dataFromChan.Item)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("AggregateCollectionStats %s ", dataFromChan.Item.Contract), zap.Error(err))
			return err
		}
	}

	return nil
}
