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
	Id                  int            `db:"id" json:"id"`
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
		json.NewEncoder(w).Encode("successfully created")
	}
}

func UpdateHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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

		json.NewEncoder(w).Encode("updated")
	}
}

func GetHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// add middleware to check if the requesting user is part of the house

		vars := mux.Vars(r)
		houseId, ok := vars["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		sqlFindHouse := `
			SELECT * FROM houses WHERE id = $1
		`

		var house HousePayload

		err := db.QueryRowx(sqlFindHouse, houseId).StructScan(&house)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(house)

	}
}

func RemoveHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		houseId, ok := vars["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		sqlRemoveHouse := `
			DELETE FROM houses WHERE id = $1
		`

		_, err := db.Exec(sqlRemoveHouse, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully deleted")
	}
}