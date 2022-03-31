package broadcast

type DataRecv struct {
	Msg string `json:"msg"`
	Dep string `json:"dep"`
}

type DataResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type DataTrade struct {
	StratName  string  `json:"strat_name"`
	Symbol     string  `json:"symbol"`
	Satisfy    float32 `json:"satisfactory"`
	MaxInvestR float32 `json:"max_invest_ratio"`
	MaxInvestM float32 `json:"max_invest_money"`
	OrderFillT int32   `json:"orderfill_time"`
	MaxiTradeT int32   `json:"max_trade_time"`
}

type MessageRecv struct {
	SignalType string   `json:"signal_type"`
	Data       DataRecv `json:"data"`
}

type MessageResp struct {
	SignalType string   `json:"signal_type"`
	Data       DataResp `json:"data"`
}

type OrderRecv struct {
	SignalType string `json:"signal_type"`
	DT         string `json:"date"`
	Trader     string `json:"trader"`
	AssetType  string `json:"asset_type"`
}
