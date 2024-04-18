package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Account struct {
	MarginMode          string `json:"marginMode"`
	UpdatedTime         string `json:"updatedTime"`
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	SpotHedgingStatus   string `json:"spotHedgingStatus"`
}

func (api *MyAPI) GetAccount() (*Account, error) {
	url := "https://api.bybit.com/v5/account/info"
	ts := api.Timestamp()
	signature := api.Signature(ts, api)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Failed to create request:", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-BAPI-API-KEY", api.Config.ApiKey)
	req.Header.Add("X-BAPI-TIMESTAMP", ts)
	req.Header.Add("X-BAPI-SIGN", signature)
	req.Header.Add("X-BAPI-RECV-WINDOW", api.Config.RecvWindow)

	res, err := api.Api.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type responseBody struct {
		RetCode int     `json:"retCode"`
		RetMsg  string  `json:"retMsg"`
		Result  Account `json:"result"`
	}
	var response responseBody

	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("Ошибка десериализации JSON:", err)
		return nil, err
	}

	return &response.Result, nil

}
