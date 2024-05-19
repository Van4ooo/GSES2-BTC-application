package third_party_api

import (
	"encoding/json"
	"net/http"
)

const (
	apiURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah"
)

type RateResponse struct {
	Bitcoin struct {
		UAH float64 `json:"uah"`
	} `json:"bitcoin"`
}

func GetRateBTC_UAH() (float64, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	var rateResponse RateResponse
	if err := json.NewDecoder(response.Body).Decode(&rateResponse); err != nil {
		return -1, err
	}

	return rateResponse.Bitcoin.UAH, nil
}
