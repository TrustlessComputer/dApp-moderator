package structure

import "time"

type ReportErc20 struct {
	Result interface{} `json:"result"`
	Data   []struct {
		Address           string `json:"address"`
		AddressCheck      string `json:"address_check"`
		TotalSupply       string `json:"total_supply"`
		TotalSupplyNumber string `json:"total_supply_number"`
		Owner             string `json:"owner"`
		Decimal           int    `json:"decimal"`
		DeployedAtBlock   int    `json:"deployed_at_block"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Thumbnail         string `json:"thumbnail"`
		Description       string `json:"description"`
		Social            struct {
			Website   string `json:"website"`
			Discord   string `json:"discord"`
			Twitter   string `json:"twitter"`
			Telegram  string `json:"telegram"`
			Medium    string `json:"medium"`
			Instagram string `json:"instagram"`
		} `json:"social"`
		Index           int     `json:"index"`
		Volume          string  `json:"volume"`
		TotalVolume     string  `json:"total_volume"`
		BtcVolume       string  `json:"btc_volume"`
		UsdVolume       string  `json:"usd_volume"`
		BtcTotalVolume  string  `json:"btc_total_volume"`
		UsdTotalVolume  string  `json:"usd_total_volume"`
		MarketCap       string  `json:"market_cap"`
		UsdMarketCap    string  `json:"usd_market_cap"`
		Price           string  `json:"price"`
		BtcPrice        string  `json:"btc_price"`
		UsdPrice        string  `json:"usd_price"`
		Percent         float64 `json:"percent"`
		Percent7Day     float64 `json:"percent_7day"`
		Network         string  `json:"network"`
		Priority        int     `json:"priority"`
		BaseTokenSymbol string  `json:"base_token_symbol"`
		Status          string  `json:"status"`
		Chart           []struct {
			Time             time.Time `json:"time"`
			Timestamp        int       `json:"timestamp"`
			VolumeFrom       string    `json:"volume_from"`
			VolumeTo         string    `json:"volume_to"`
			BtcPrice         string    `json:"btc_price"`
			UsdPrice         string    `json:"usd_price"`
			Low              string    `json:"low"`
			Open             string    `json:"open"`
			Close            string    `json:"close"`
			High             string    `json:"high"`
			VolumeFromUsd    string    `json:"volume_from_usd"`
			VolumeToUsd      string    `json:"volume_to_usd"`
			TotalVolumeUsd   string    `json:"total_volume_usd"`
			TotalVolume      string    `json:"total_volume"`
			LowUsd           string    `json:"low_usd"`
			OpenUsd          string    `json:"open_usd"`
			CloseUsd         string    `json:"close_usd"`
			HighUsd          string    `json:"high_usd"`
			ConversionType   string    `json:"conversion_type"`
			ConversionSymbol string    `json:"conversion_symbol"`
		} `json:"chart"`
	} `json:"data"`
	Error error `json:"error"`
}
