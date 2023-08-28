package services

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetElectricityPrices() ([]map[string]interface{}, error) {
	resp, err := http.Get("https://www.elprisenligenu.dk/api/v1/prices/2023/08-24_DK1.json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
