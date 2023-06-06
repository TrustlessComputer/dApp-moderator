package usecase

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
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
	symbol := ""
	decimal := 0
	if tokenAddress == strings.ToLower(os.Getenv("WETH_ADDRESS")) {
		symbol = "eth"
		decimal = 18
	}

	if tokenAddress == strings.ToLower(os.Getenv("WBTC_ADDRESS")) {
		symbol = "eth"
		decimal = 8
	}

	if symbol == "" {
		err = errors.New("Cannot detect erc20 token")
		logger.AtLog.Logger.Error("StartWorker - GetExternalPrice", zap.Error(err), zap.String("tokenAddress", tokenAddress))

	} else {

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

		volume.Erc20Rate = rate
		volume.Erc20Decimal = decimal
		if decimal != 0 {
			erc20Val := helpers.GetValue(fmt.Sprintf("%d", volume.TotalVolume), float64(volume.Erc20Decimal))
			volume.USDTValue = erc20Val * rate
		} else {
			volume.USDTValue = 0
		}

	}

	return err
}

func (u *Usecase) StartWorker(inputItemChan chan entity.MarketplaceCollectionAggregation, outputChan chan outputMkpCollectionChan) {
	inputItem := <-inputItemChan
	minNumber := float64(99999999)
	var err error

	//calculate usdt by erc20 token
	total := float64(0)
	min := minNumber
	volumes := inputItem.MarketPlaceVolumes
	floorPriceVolumes := inputItem.FloorPriceMarketPlaceVolumes

	for _, volume := range floorPriceVolumes {
		err = u.calculateRate(volume)
		if min > volume.USDTValue {
			min = volume.USDTValue
		}
	}

	for _, volume := range volumes {
		err = u.calculateRate(volume)
		total += volume.USDTValue
	}

	inputItem.Volume = total
	if min == minNumber {
		min = 0
	}

	inputItem.FloorPrice = min
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
			logger.AtLog.Logger.Error(fmt.Sprintf("AggregateCollectionStats %s ", dataFromChan.Item.ID.Hex()), zap.Error(dataFromChan.Err))
			return dataFromChan.Err
		}

		//save to view
		err := u.Repo.InsertMarketPlaceAggregation(dataFromChan.Item)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("AggregateCollectionStats %s ", dataFromChan.Item.ID.Hex()), zap.Error(err))
			return err
		}
	}

	return nil
}
