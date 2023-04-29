package blockchain_api

import (
	"bytes"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/redis"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type BlockChainApi struct {
	BaseURL                      string
	ChainID                      int64
	UniswapV2FactoryContractAddr string
	UniswapV2RouterContractAddr  string
}

func NewBlockChainApi(conf *config.Config, cache redis.IRedisCache) *BlockChainApi {
	return &BlockChainApi{
		BaseURL:                      conf.BlockChainApi.BaseURL,
		UniswapV2FactoryContractAddr: conf.BlockChainApi.UniswapV2FactoryContractAddr,
		UniswapV2RouterContractAddr:  conf.BlockChainApi.UniswapV2RouterContractAddr,
	}
}

func (c *BlockChainApi) buildUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.BaseURL + "/" + resourcePath
	}
	return c.BaseURL
}

func (c *BlockChainApi) doWithoutAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *BlockChainApi) methodJSON(method string, apiURL string, jsonObject interface{}, result interface{}) error {
	var buffer io.Reader
	if jsonObject != nil {
		bodyBytes, _ := json.Marshal(jsonObject)
		buffer = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, apiURL, buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.doWithoutAuth(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
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

func (c *BlockChainApi) TcSwapEventsByTransaction(txHash string) (*TcSwapEventResp, error) {
	resp := struct {
		Result *TcSwapEventResp `json:"result"`
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl(fmt.Sprintf("swap/events/%s", txHash)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (c *BlockChainApi) TcSwapEvents(numBlocks, startBlocks, endBlocks int64) (*TcSwapEventResp, error) {
	resp := struct {
		Result *TcSwapEventResp `json:"result"`
	}{}
	contractAddrs := []string{
		c.UniswapV2FactoryContractAddr,
		c.UniswapV2RouterContractAddr,
	}

	err := c.methodJSON(
		http.MethodGet,
		c.buildUrl(fmt.Sprintf("swap/events?contract_addrs=%s&num_blocks=%d&start_blocks=%d&end_blocks=%d",
			strings.Join(contractAddrs, ","), numBlocks, startBlocks, endBlocks)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
