package entity

type SwapIdoStatus string

const (
	SwapIdoStatusUpcoming SwapIdoStatus = "upcoming"
	SwapIdoStatusStated   SwapIdoStatus = "started"
	SwapIdoStatusFinished SwapIdoStatus = "finished"
)

type SwapWrapTOkenContractAddrConfig struct {
	WbtcContractAddr  string
	WethContractAddr  string
	WusdcContractAddr string
	WpepeContractAddr string
	WordiContractAddr string

	WbtcToken  *Token
	WethToken  *Token
	WusdcToken *Token
	WpepeToken *Token
	WordiToken *Token
}
