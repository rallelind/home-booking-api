package controllers

import (
	"encoding/json"
	"home-booking-api/src/services"
	"net/http"
)

func GetElectricityPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := services.GetElectricityPrices("2023", "01", "01")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(resp)
	}
}
