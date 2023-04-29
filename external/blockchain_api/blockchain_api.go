package blockchain_api

import (
	"bytes"
	"context"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/redis"
	"dapp-moderator/utils/uniswapfactory"
	"dapp-moderator/utils/uniswappair"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockResp struct {
	time   uint64
	hash   common.Hash
	number *big.Int
}

type BlockChainApi struct {
	BaseURL                      string
	ChainID                      int64
	UniswapV2FactoryContractAddr string
	UniswapV2RouterContractAddr  string
	client                       *ethclient.Client
	InterruptMili                int
	BlockMap                     map[uint64]*BlockResp
	ScanLimitBlockNumber         int64
}

func NewBlockChainApi(conf *config.Config, cache redis.IRedisCache) *BlockChainApi {
	return &BlockChainApi{
		BaseURL:                      conf.Swap.BaseURL,
		UniswapV2FactoryContractAddr: conf.Swap.UniswapV2FactoryContractAddr,
		UniswapV2RouterContractAddr:  conf.Swap.UniswapV2RouterContractAddr,
		InterruptMili:                100,
		ScanLimitBlockNumber:         50,
	}
}

type TcSwapPairCreatedEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Token0          string   `json:"token0"`
	Token1          string   `json:"token1"`
	Pair            string   `json:"pair"`
	Arg3            *big.Int `json:"arg3"`
	Index           uint     `json:"log_index"`
}

type TcSwapSwapEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Sender          string   `json:"sender"`
	To              string   `json:"to"`
	Amount0In       *big.Int `json:"amount0_in"`
	Amount1In       *big.Int `json:"amount1_in"`
	Amount0Out      *big.Int `json:"amount0_out"`
	Amount1Out      *big.Int `json:"amount1_out"`
	Index           uint     `json:"log_index"`
}

type TcSwapMintBurnEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Sender          string   `json:"sender"`
	To              string   `json:"to"`
	Amount0         *big.Int `json:"amount0"`
	Amount1         *big.Int `json:"amount1"`
	Index           uint     `json:"log_index"`
}

type TcSwapSyncEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Reserve0        *big.Int `json:"reserve0"`
	Reserve1        *big.Int `json:"reserve1"`
	Index           uint     `json:"log_index"`
}

type TcSwapEventResp struct {
	PairCreated     []*TcSwapPairCreatedEventResp `json:"pair_created"`
	Swap            []*TcSwapSwapEventResp        `json:"swap"`
	PairMint        []*TcSwapMintBurnEventResp    `json:"mint"`
	PairBurn        []*TcSwapMintBurnEventResp    `json:"burn"`
	PairSync        []*TcSwapSyncEventResp        `json:"sync"`
	LastBlockNumber int64                         `json:"last_block_number"`
}

func (c *BlockChainApi) getClient() (*ethclient.Client, error) {
	if c.client == nil {
		client, err := ethclient.Dial(c.BaseURL)
		if err != nil {
			return nil, err
		}
		c.client = client
	}
	return c.client, nil
}

func (c *BlockChainApi) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *BlockChainApi) postJSON(apiURL string, headers map[string]string, jsonObject interface{}, result interface{}) error {
	bodyBytes, _ := json.Marshal(jsonObject)
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.Unmarshal(bodyBytes, result)
	}
	return nil
}

func (c *BlockChainApi) TcSwapEventResp(resp *TcSwapEventResp, log *types.Log) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	fmt.Println(log.TxHash.Hex())

	uniswap, err := uniswapfactory.NewUniswapfactory(log.Address, client)
	if err != nil {
		return err
	}
	// ParsePoolCreated
	{
		logParsed, err := uniswap.ParsePairCreated(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.PairCreated = append(
				resp.PairCreated,
				&TcSwapPairCreatedEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					Token0:          logParsed.Token0.Hex(),
					Token1:          logParsed.Token1.Hex(),
					Arg3:            logParsed.Arg3,
					Pair:            logParsed.Pair.Hex(),
					Index:           log.Index,
				},
			)
		}
	}

	uniswappair, err := uniswappair.NewUniswappair(log.Address, client)
	if err != nil {
		return err
	}
	// ParseSwap
	{
		logParsed, err := uniswappair.ParseSwap(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.Swap = append(
				resp.Swap,
				&TcSwapSwapEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					Sender:          logParsed.Sender.Hex(),
					To:              logParsed.To.Hex(),
					Amount0In:       logParsed.Amount0In,
					Amount1In:       logParsed.Amount1In,
					Amount0Out:      logParsed.Amount0Out,
					Amount1Out:      logParsed.Amount1Out,
					Index:           log.Index,
				},
			)
		}
	}

	// ParseMint
	{
		logParsed, err := uniswappair.ParseMint(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.PairMint = append(
				resp.PairMint,
				&TcSwapMintBurnEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					Sender:          logParsed.Sender.Hex(),
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					Index:           log.Index,
				},
			)
		}
	}

	// ParseBurn
	{
		logParsed, err := uniswappair.ParseBurn(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.PairBurn = append(
				resp.PairBurn,
				&TcSwapMintBurnEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					Sender:          logParsed.Sender.Hex(),
					To:              logParsed.To.Hex(),
					Amount0:         logParsed.Amount0,
					Amount1:         logParsed.Amount1,
					Index:           log.Index,
				},
			)
		}
	}

	// ParseSync
	{
		logParsed, err := uniswappair.ParseSync(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.PairSync = append(
				resp.PairSync,
				&TcSwapSyncEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					Reserve0:        logParsed.Reserve0,
					Reserve1:        logParsed.Reserve1,
					Index:           log.Index,
				},
			)
		}
	}

	return nil
}

func (c *BlockChainApi) NewTcSwapEventResp() *TcSwapEventResp {
	return &TcSwapEventResp{}
}

func (c *BlockChainApi) TcSwapEvents(numBlocks, startBlock, endBlock int64) (*TcSwapEventResp, error) {
	resp := c.NewTcSwapEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	contractAddresses := []common.Address{}
	contractAddresses = append(contractAddresses, common.HexToAddress(c.UniswapV2FactoryContractAddr))
	contractAddresses = append(contractAddresses, common.HexToAddress(c.UniswapV2RouterContractAddr))

	ctx := context.Background()
	lastBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	lastNumber := lastBlock.Number.Int64()
	if endBlock != 0 && endBlock <= lastNumber {
		lastNumber = endBlock
	}

	resp.LastBlockNumber = lastNumber

	if startBlock != 0 {
		numBlocks = lastNumber - (startBlock - 10)
	}

	num := (numBlocks / c.ScanLimitBlockNumber) + 1
	for i := int64(0); i < num; i++ {
		c.Interrupt()
		logs, err := client.FilterLogs(
			ctx,
			ethereum.FilterQuery{
				FromBlock: big.NewInt(lastNumber - c.ScanLimitBlockNumber),
				ToBlock:   big.NewInt(lastNumber),
				Addresses: contractAddresses,
				Topics:    [][]common.Hash{},
			},
		)
		if err != nil {
			return nil, err
		}
		for _, log := range logs {
			err = c.TcSwapEventResp(resp, &log)
			if err != nil {
				return nil, err
			}
		}
		lastNumber -= 25
	}
	return resp, nil
}

func (c *BlockChainApi) TcSwapEventsByTransaction(txHash string) (*TcSwapEventResp, error) {
	resp := c.NewTcSwapEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	c.Interrupt()
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	for _, log := range receipt.Logs {
		err = c.TcSwapEventResp(resp, log)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *BlockChainApi) Interrupt() {
	if c.InterruptMili > 0 {
		time.Sleep(time.Duration(c.InterruptMili) * time.Millisecond)
	}
}

func (c *BlockChainApi) getBlock(n uint64) (*BlockResp, error) {
	if c.BlockMap == nil {
		c.BlockMap = map[uint64]*BlockResp{}
	}
	blockResp, ok := c.BlockMap[n]
	if !ok {
		var blockInfoResp struct {
			Result *struct {
				Timestamp string `json:"timestamp"`
				Hash      string `json:"hash"`
			} `json:"result"`
		}
		err := c.postJSON(
			c.BaseURL,
			map[string]string{},
			map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      1,
				"method":  "eth_getBlockByNumber",
				"params": []interface{}{
					fmt.Sprintf("0x%s", big.NewInt(int64(n)).Text(16)),
					false,
				},
			},
			&blockInfoResp,
		)
		if err != nil {
			return nil, err
		}
		bn, _ := big.NewInt(0).SetString(blockInfoResp.Result.Timestamp[2:], 16)
		c.BlockMap[n] = &BlockResp{
			time:   bn.Uint64(),
			hash:   common.HexToHash(blockInfoResp.Result.Hash),
			number: big.NewInt(int64(n)),
		}
		blockResp = c.BlockMap[n]
	}
	return blockResp, nil
}

func (b *BlockResp) Time() uint64 {
	return b.time
}
