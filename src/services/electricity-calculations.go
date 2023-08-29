package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ElectricityPrice struct {
	DkkPerKwh float64 `json:"DKK_per_kWh"`
	EurPerKwh float64 `json:"EUR_per_kWh"`
	EXR       float64 `json:"EXR"`
	TimeStart string  `json:"time_start"`
	TimeEnd   string  `json:"time_end"`
}

func GetElectricityPrices(year string, month string, day string) ([]ElectricityPrice, error) {

	urlTemplate := "https://www.elprisenligenu.dk/api/v1/prices/%s/%s-%s_DK1.json"

	electricityPriceUrl := fmt.Sprintf(urlTemplate, year, month, day)

	resp, err := http.Get(electricityPriceUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []ElectricityPrice

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
