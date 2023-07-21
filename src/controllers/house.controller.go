package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type HousePayload struct {
	Address             string         `db:"address" json:"address"`
	HouseName           string         `db:"house_name" json:"house_name"`
	AdminNeedsToApprove bool           `db:"admin_needs_to_approve" json:"admin_needs_to_approve"`
	LoginImages         pq.StringArray `db:"login_images" json:"login_images"`
}

func CreateHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createHousePayload HousePayload

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

func UpdateHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateHousePayload HousePayload

		vars := mux.Vars(r)
		id, ok := vars["houseId"]

		if !ok {
			http.Error(w, "please provide the house id", http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&updateHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sqlUpdate := `
			UPDATE houses SET house_name = $1, admin_needs_to_approve = $2, login_images = $3
			WHERE id = $4
		`

		_, err = db.Exec(sqlUpdate, updateHousePayload.HouseName, updateHousePayload.AdminNeedsToApprove, updateHousePayload.LoginImages, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("created")
	}	
}
