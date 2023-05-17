package entity

type SwapIdoStatus string
type SwapBaseTokenSymbol string

const (
	SwapIdoStatusUpcoming SwapIdoStatus = "upcoming"
	SwapIdoStatusStated   SwapIdoStatus = "started"
	SwapIdoStatusFinished SwapIdoStatus = "finished"

	SwapBaseTokenSymbolWBTC SwapBaseTokenSymbol = "WBTC"
	SwapBaseTokenSymbolWETH SwapBaseTokenSymbol = "ETH"
)

type SwapWrapTOkenContractAddrConfig struct {
	WbtcContractAddr    string
	WethContractAddr    string
	WusdcContractAddr   string
	WpepeContractAddr   string
	WordiContractAddr   string
	RouterContractAddr  string
	FactoryContractAddr string

	WbtcToken  *Token
	WethToken  *Token
	WusdcToken *Token
	WpepeToken *Token
	WordiToken *Token
}
