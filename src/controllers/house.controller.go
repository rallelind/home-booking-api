package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CreateHousePayload struct {
	Address             string         `db:"address" json:"address"`
	HouseName           string         `db:"house_name" json:"house_name"`
	AdminNeedsToApprove bool           `db:"admin_needs_to_approve" json:"admin_needs_to_approve"`
	LoginImages         pq.StringArray `db:"login_images" json:"login_images"`
}

func CreateHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createHousePayload CreateHousePayload

		err := json.NewDecoder(r.Body).Decode(&createHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertHouse := `
			INSERT INTO houses (address, house_name, admin_needs_to_approve, login_images)
			VALUES (:address, :house_name, :admin_needs_to_approve, :login_images)
		`

		_, err = db.NamedExec(insertHouse, createHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createHousePayload)
	}
}
