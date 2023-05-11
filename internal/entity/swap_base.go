package entity

type SwapIdoStatus string

const (
	SwapIdoStatusUpcoming SwapIdoStatus = "upcoming"
	SwapIdoStatusStated   SwapIdoStatus = "started"
	SwapIdoStatusFinished SwapIdoStatus = "finished"
)
