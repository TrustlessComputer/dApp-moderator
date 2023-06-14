package txTCServer

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type eventLog struct {
	Log       types.Log
	Event     entity.TokenActivityType
	Activity  *entity.MarketplaceTokenActivity
	EventData interface{}
	Err       error
	PFunc     proccessFunc
}

type proccessFunc func(eventData interface{}, chainLog types.Log) error

func (c *txTCServer) ProcessLog(resultChan chan *eventLog) {
	dataFromChan := <-resultChan
	err := dataFromChan.Err
	if err == nil {
		_func := dataFromChan.PFunc
		err := _func(dataFromChan.EventData, dataFromChan.Log)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("Proccess funnc - %s", dataFromChan.Log.TxHash), zap.Error(err))
		}

		activity := dataFromChan.Activity
		err = c.Usecase.InsertActivity(activity)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("InsertActivity - %s", dataFromChan.Log.TxHash), zap.Error(err))
		}
	}
	//} else {
	//	//logger.AtLog.Logger.Error(fmt.Sprintf("ProcessLog - %s", dataFromChan.Log.TxHash), zap.Error(err))
	//}
}

func (c *txTCServer) TokenEvents(wg *sync.WaitGroup, ctx context.Context, fromBlock int64, toBlock int64) {
	defer wg.Done()

	//logs are only heard by our collection addresses
	logs, err := c.Blockchain.GetEventLogs(*big.NewInt(fromBlock), *big.NewInt(toBlock), c.ListenedContracts())
	if err != nil {
		logger.AtLog.Logger.Error("err.GetEventLogs", zap.String("err", err.Error()))
		return
	}

	logger.AtLog.Logger.Info("resolveTransaction", zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock), zap.Int("logs", len(logs)))

	//start channel
	inputDataChan := make(chan types.Log, len(logs))
	resultChan := make(chan *eventLog, len(logs))

	//init workers (pool)
	for _, _ = range logs {
		go c.Worker(inputDataChan, resultChan)
	}

	// push data to worker | 100 logs / 500 milisecond
	for i, log := range logs {
		inputDataChan <- log
		if i > 0 && i%100 == 0 {
			time.Sleep(time.Millisecond * 500)
		}
	}

	//processing
	for _, _ = range logs {
		c.ProcessLog(resultChan)
	}
}

func (c *txTCServer) Worker(inputDataChan chan types.Log, result chan *eventLog) {
	log := <-inputDataChan

	topics := log.Topics
	event := topics[0]
	mkpEventName := strings.ToLower(event.String())

	var pFunction proccessFunc
	var eventType entity.TokenActivityType
	var err error
	activity := &entity.MarketplaceTokenActivity{}
	var eventData interface{}

	switch mkpEventName {
	case c.MarketPlace.Events["MARKETPLACE_TRANSFER_EVENT"]:
		pFunction = c.Usecase.TransferToken
		eventType = entity.TokenTransfer
		break
	case c.MarketPlace.Events["MARKETPLACE_LIST_TOKEN"]:
		pFunction = c.Usecase.MarketplaceCreateListing
		eventType = entity.TokenListing
		break
	case c.MarketPlace.Events["MARKETPLACE_PURCHASE_TOKEN"]:
		pFunction = c.Usecase.MarketplacePurchaseListing
		eventType = entity.TokenPurchase
		break
	case c.MarketPlace.Events["MARKETPLACE_CANCEL_LISTING"]:
		pFunction = c.Usecase.MarketplaceCancelListing
		eventType = entity.TokenCancelListing
		break
	case c.MarketPlace.Events["MARKETPLACE_MAKE_OFFER"]:
		pFunction = c.Usecase.MarketplaceMakeOffer
		eventType = entity.TokenMakeOffer
		break
	case c.MarketPlace.Events["MARKETPLACE_ACCEPT_MAKE_OFFER"]:
		pFunction = c.Usecase.MarketplaceAcceptOffer
		eventType = entity.TokenAcceptOffer
		break
	case c.MarketPlace.Events["MARKETPLACE_CANCEL_MAKE_OFFER"]:
		pFunction = c.Usecase.MarketplaceCancelOffer
		eventType = entity.TokenCancelOffer
		break
	case c.MarketPlace.Events["MARKETPLACE_BNS_RESOLVER_UPDATED"]:
		pFunction = c.Usecase.MarketplaceBNSResolverUpdated
		eventType = entity.BNSResolverUpdated
		break
	case c.MarketPlace.Events["MARKETPLACE_BNS_REGISTERED"]:
		pFunction = c.Usecase.MarketplaceBNSCreated
		eventType = entity.BNSResolverCreated
		break
	case c.MarketPlace.Events["MARKETPLACE_BNS_SET_FPF"]:
		pFunction = c.Usecase.MarketplacePFPUpdated
		eventType = entity.BNSPfpUpdated
		break
	case c.MarketPlace.Events["AUCTION_CREATED_EVENT"]:
		pFunction = c.Usecase.AuctionCreated
		eventType = entity.AuctionCreated
		break
	}

	activity, eventData, err = c.Usecase.ParseMkplaceData(log, eventType)
	result <- &eventLog{
		Log:       log,
		EventData: eventData,
		Event:     eventType,
		Activity:  activity,
		Err:       err,
		PFunc:     pFunction,
	}

}
