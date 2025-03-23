package clickwebsocket

import "encoding/json"

type Message struct {
	TypeMessage string          `json:"type_message"`
	Data        json.RawMessage `json:"data"`
}

type Validate struct {
	Valid float64 `json:"valid"`
	Nonce float64 `json:"nonce"`
}

type ClickBatch struct {
	ClicksInfo []click `json:"clicks_info"`
	SendTime   int64   `json:"send_time"`
}

type click struct {
	ClickValue float64 `json:"click_value"`
	ClickTime  int64   `json:"click_time"`
}

type UserInfo struct {
	ValidClicks float64 `json:"valid_clicks"`
	Clicks      float64 `json:"all_clicks"`
}
