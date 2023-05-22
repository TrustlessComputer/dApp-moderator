package blockchain_api

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/erc20"
	"dapp-moderator/utils/gmpayment"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"dapp-moderator/utils/uniswapfactory"
	"dapp-moderator/utils/uniswappair"
	"dapp-moderator/utils/uniswaprouter"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockResp struct {
	time   uint64
	hash   common.Hash
	number *big.Int
}

type BlockChainApi struct {
	BaseURL              string
	ChainID              int64
	client               *ethclient.Client
	InterruptMili        int
	BlockMap             map[uint64]*BlockResp
	ScanLimitBlockNumber int64
}

func NewBlockChainApi(conf *config.Config, cache redis.IRedisCache) *BlockChainApi {
	return &BlockChainApi{
		BaseURL:              conf.Swap.BaseURL,
		InterruptMili:        100,
		ScanLimitBlockNumber: 50,
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

type TcGmPaymentPaidEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	AmountGM        *big.Int `json:"amount_gm"`
	User            string   `json:"user"`
	Index           uint     `json:"log_index"`
}

type TcSwapEventResp struct {
	PairCreated     []*TcSwapPairCreatedEventResp `json:"pair_created"`
	Swap            []*TcSwapSwapEventResp        `json:"swap"`
	PairMint        []*TcSwapMintBurnEventResp    `json:"mint"`
	PairBurn        []*TcSwapMintBurnEventResp    `json:"burn"`
	PairSync        []*TcSwapSyncEventResp        `json:"sync"`
	GmPaymentPaid   []*TcGmPaymentPaidEventResp   `json:"gm_paid"`
	LastBlockNumber int64                         `json:"last_block_number"`
}

type TcTmTokenTransferEventResp struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	From            string   `json:"from"`
	To              string   `json:"to"`
	Value           *big.Int `json:"value"`
	Index           uint     `json:"log_index"`
}

type TcTmTokenEventResp struct {
	Transfer        []*TcTmTokenTransferEventResp `json:"transfer"`
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

func (c *BlockChainApi) parsePrkAuth(prkHex string) (common.Address, *ecdsa.PrivateKey, error) {
	prk, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return common.Address{}, nil, err
	}
	pbk, ok := prk.Public().(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, errors.New("error casting public key to ECDSA")
	}
	pbkHex := crypto.PubkeyToAddress(*pbk)
	return pbkHex, prk, nil
}

func (c *BlockChainApi) getGasPrice() (*big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	c.Interrupt()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(12))
	gasPrice = gasPrice.Div(gasPrice, big.NewInt(10))
	return gasPrice, nil
}

func (c *BlockChainApi) WaitMined(hash string) error {
	time.Sleep(5 * time.Second)
	client, err := c.getClient()
	if err != nil {
		return err
	}
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	r, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return errors.New("transaction is not Successful")
	}
	return nil
}

func (c *BlockChainApi) DefaultSignerFn(prk *ecdsa.PrivateKey) bind.SignerFn {
	return func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		client, err := c.getClient()
		if err != nil {
			return nil, err
		}
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}
		signedTx, err := types.SignTx(t, types.NewEIP155Signer(chainID), prk)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}
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

	instance, err := gmpayment.NewGmpayment(log.Address, client)
	if err != nil {
		return err
	}

	// ParsePaid
	{
		logParsed, err := instance.ParsePaid(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.GmPaymentPaid = append(
				resp.GmPaymentPaid,
				&TcGmPaymentPaidEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					AmountGM:        logParsed.AmountGM,
					User:            logParsed.User.Hex(),
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

func (c *BlockChainApi) Erc20TotalSupply(erc20Addr string) (*big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(common.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	c.Interrupt()
	balance, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *BlockChainApi) Erc20GetCoinBalance(erc20Addr, accountAddress string) (*big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(common.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	c.Interrupt()
	balance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(accountAddress))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *BlockChainApi) TcSwapEvents(contracts []string, numBlocks, startBlock, endBlock int64) (*TcSwapEventResp, error) {
	resp := c.NewTcSwapEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	contractAddresses := []common.Address{}
	for _, item := range contracts {
		contractAddresses = append(contractAddresses, common.HexToAddress(item))
	}

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
		lastNumber -= c.ScanLimitBlockNumber
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

func (c *BlockChainApi) GetBitcoinPrice() (float64, error) {
	headers := make(map[string]string)
	data, _, _, err := helpers.JsonRequest("https://api.coingecko.com/api/v3/simple/price?ids=BITCOIN&vs_currencies=USD", "GET", headers, nil)
	if err != nil {
		return 0, err
	}

	type USDResp struct {
		Usd float64 `json:"usd"`
	}

	type CoingeckoResp struct {
		Bitcoin *USDResp `json:"bitcoin"`
	}

	resp := &CoingeckoResp{}
	err = helpers.ParseData(data, resp)
	if err != nil {
		return 0, err
	}

	return resp.Bitcoin.Usd, nil
}

func (c *BlockChainApi) GetEthereumPrice() (float64, error) {
	headers := make(map[string]string)
	data, _, _, err := helpers.JsonRequest("https://api.coingecko.com/api/v3/simple/price?ids=ETHEREUM&vs_currencies=USD", "GET", headers, nil)
	if err != nil {
		return 0, err
	}

	type USDResp struct {
		Usd float64 `json:"usd"`
	}

	type CoingeckoResp struct {
		Bitcoin *USDResp `json:"ethereum"`
	}

	resp := &CoingeckoResp{}
	err = helpers.ParseData(data, resp)
	if err != nil {
		return 0, err
	}

	return resp.Bitcoin.Usd, nil
}

func (c *BlockChainApi) TcTmTokenEvents(contracts []string, numBlocks, startBlock, endBlock int64) (*TcTmTokenEventResp, error) {
	resp := c.NewTcTmTokenEventResp()
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	contractAddresses := []common.Address{}
	for _, item := range contracts {
		contractAddresses = append(contractAddresses, common.HexToAddress(item))
	}

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
			err = c.TcTmTokenEventResp(resp, &log)
			if err != nil {
				return nil, err
			}
		}
		lastNumber -= c.ScanLimitBlockNumber
	}
	return resp, nil
}

func (c *BlockChainApi) NewTcTmTokenEventResp() *TcTmTokenEventResp {
	return &TcTmTokenEventResp{}
}

func (c *BlockChainApi) TcTmTokenEventResp(resp *TcTmTokenEventResp, log *types.Log) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}

	erc20, err := erc20.NewErc20(log.Address, client)
	if err != nil {
		return err
	}

	// ParseTransfer
	{
		logParsed, err := erc20.ParseTransfer(*log)
		if err == nil {
			block, err := c.getBlock(log.BlockNumber)
			if err != nil {
				return err
			}
			resp.Transfer = append(
				resp.Transfer,
				&TcTmTokenTransferEventResp{
					TxHash:          log.TxHash.Hex(),
					ContractAddress: logParsed.Raw.Address.Hex(),
					Timestamp:       block.Time(),
					From:            logParsed.From.Hex(),
					To:              logParsed.To.Hex(),
					Value:           logParsed.Value,
					Index:           log.Index,
				},
			)
		}
	}

	return nil
}

func (c *BlockChainApi) TcSwapGetReserves(pairAddress string) (*big.Int, *big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, nil, err
	}
	instance, err := uniswappair.NewUniswappair(common.HexToAddress(pairAddress), client)
	if err != nil {
		return nil, nil, err
	}
	c.Interrupt()
	reserve, err := instance.GetReserves(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}
	return reserve.Reserve0, reserve.Reserve1, nil
}

func (c *BlockChainApi) TcSwapGetAmountsOut(routerAddress string, amountIn *big.Int, pathStr []string) ([]*big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := uniswaprouter.NewUniswaprouter(common.HexToAddress(routerAddress), client)
	if err != nil {
		return nil, err
	}
	c.Interrupt()
	pathAddr := []common.Address{}
	for _, item := range pathStr {
		pathAddr = append(pathAddr, common.HexToAddress(item))
	}
	amountOut, err := instance.GetAmountsOut(&bind.CallOpts{}, amountIn, pathAddr)
	if err != nil {
		return nil, err
	}
	return amountOut, nil
}

func (c *BlockChainApi) TcSwapGetAmountsIn(routerAddress string, amountOut *big.Int, pathStr []string) ([]*big.Int, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := uniswaprouter.NewUniswaprouter(common.HexToAddress(routerAddress), client)
	if err != nil {
		return nil, err
	}
	c.Interrupt()
	pathAddr := []common.Address{}
	for _, item := range pathStr {
		pathAddr = append(pathAddr, common.HexToAddress(item))
	}
	amountIn, err := instance.GetAmountsIn(&bind.CallOpts{}, amountOut, pathAddr)
	if err != nil {
		return nil, err
	}
	return amountIn, nil
}

func (c *BlockChainApi) TcSwapExactTokensForTokens(routerAddress string, amountIn, amountOutMin *big.Int, trader, prkHex, fromToken, toToken string) (string, error) {
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	routerContractAddress := common.HexToAddress(routerAddress)
	instance, err := uniswaprouter.NewUniswaprouter(routerContractAddress, client)
	if err != nil {
		return "", err
	}
	c.Interrupt()

	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}

	nonceAuth, err := client.NonceAt(context.Background(), pbkHex, nil)
	if err != nil {
		return "", err
	}
	auth := bind.NewKeyedTransactor(prk)
	auth.Nonce = big.NewInt(int64(nonceAuth))
	auth.Value = big.NewInt(0)         // in wei
	auth.GasLimit = uint64(2500000000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = c.DefaultSignerFn(prk)

	path := []common.Address{
		common.HexToAddress(fromToken),
		common.HexToAddress(toToken),
	}
	deadline := big.NewInt(time.Now().Unix() + 60*30)
	traderAddres := common.HexToAddress(trader)
	// estimate tx
	fmt.Printf(`amountIn=%s`, amountIn.String())
	fmt.Printf(`amountOutMin=%s`, amountOutMin.String())

	{
		pAbi, err := abi.JSON(strings.NewReader(uniswaprouter.UniswaprouterABI))
		if err != nil {
			return "", err
		}
		data, err := pAbi.Pack(
			"swapExactTokensForTokens",
			amountIn, amountOutMin, path, traderAddres, deadline,
		)
		if err != nil {
			return "", err
		}
		c.Interrupt()
		gasUsed, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From: pbkHex,
			To:   &routerContractAddress,
			Data: data,
		})
		if err != nil {
			return "", err
		}
		// // check fee <= 0.0003 ETH
		// limitGas := big.NewInt(int64(configs.GetHighGasConfig() * 1e18))
		limitGas := big.NewInt(int64(0.003 * 1e18))
		if new(big.Int).Mul(big.NewInt(int64(gasUsed)), gasPrice).Cmp(limitGas) > 0 {
			estGas := new(big.Int).Mul(big.NewInt(int64(gasUsed)), gasPrice)
			return estGas.String(), errors.New("gas fee is too expensive")
		}
		auth.GasLimit = gasUsed * 12 / 10 // more 20%
	}
	c.Interrupt()

	tnx, err := instance.SwapExactTokensForTokens(auth, amountIn, amountOutMin, path, traderAddres, deadline)
	if err != nil {
		return "", err
	}
	err = c.WaitMined(tnx.Hash().Hex())
	if err != nil {
		return "", err
	}
	return tnx.Hash().Hex(), nil
}

/////////////GM PAYMENT
func (c *BlockChainApi) SignWithEthereum(privateKey string, dataBytes []byte) (string, error) {
	signBytes := append([]byte("\x19Ethereum Signed Message:\n32"), dataBytes...)
	hash := crypto.Keccak256Hash(signBytes)
	prk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hash.Bytes(), prk)
	if err != nil {
		return "", err
	}
	signature[crypto.RecoveryIDOffset] += 27
	sigHex := hexutil.Encode(signature)
	sigHex = sigHex[2:]
	return sigHex, nil
}

func (c *BlockChainApi) GmPaymentSignMessage(contractAddr, adminAddrr, adminPrk, userAddrr, tokenAddr string, chainID, amount *big.Int) (string, error) {
	datas := []byte{}
	datas = append(datas, common.HexToHash(contractAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(chainID.Bytes()).Bytes()...)
	datas = append(datas, common.HexToHash(adminAddrr).Bytes()...)
	datas = append(datas, common.HexToHash(userAddrr).Bytes()...)
	datas = append(datas, common.HexToHash(tokenAddr).Bytes()...)
	datas = append(datas, common.BytesToHash(amount.Bytes()).Bytes()...)

	dataByteHash := crypto.Keccak256Hash(
		datas,
	)
	signature, err := c.SignWithEthereum(adminPrk, dataByteHash.Bytes())
	if err != nil {
		return "", err
	}

	return signature, nil
}
