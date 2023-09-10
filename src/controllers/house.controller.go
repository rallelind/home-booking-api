package controllers

import (
	"database/sql"
	"encoding/json"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"home-booking-api/src/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func CreateHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var createHousePayload models.HouseModel

		err := json.NewDecoder(r.Body).Decode(&createHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.NamedExec(queries.CreateHouseQuery, createHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("successfully created")
	}
}

func GetUserHouses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user := services.GetCurrentUser(r)

		var allHouses []models.HouseModel

		err := db.Select(&allHouses, queries.FindUserHousesQuery, user.EmailAddresses[0].EmailAddress)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "No houses were found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(allHouses)

	}
}

func UpdateAdminNeedsToApprove(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var updateHousePayload models.HouseModel

		id, ok := mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "please provide the house id", http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&updateHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.ChangeAdminNeedsToApprove, updateHousePayload.AdminNeedsToApprove, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("updated")
	}
}

func GetHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// add middleware to check if the requesting user is part of the house

		houseId, ok := mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		var house models.HouseModel

		err := db.QueryRowx(queries.FindHouseQuery, houseId).StructScan(&house)

		if err != nil {
			log.Print(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(house)

	}
}

func RemoveHouse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		houseId, ok := mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(queries.RemoveHouseQuery, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully deleted")
	}
}

func UploadHouseImages(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var houseId, ok = mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		resp, err := services.UploadImageToCloudinary(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.AddHouseImages, resp.SecureURL, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func UpdateAdminUsers(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var houseId, ok = mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		var updateHousePayload models.HouseModel

		err := json.NewDecoder(r.Body).Decode(&updateHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.AddHouseAdmin, updateHousePayload.HouseAdmins, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")
	}
}

func RemoveAdminUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var houseId, ok = mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		var updateHousePayload models.HouseModel

		err := json.NewDecoder(r.Body).Decode(&updateHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.RemoveHouseAdmin, updateHousePayload.HouseAdmins, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")
	}
}

func RemoveHouseImage(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var houseId, ok = mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing path param houseId", http.StatusBadRequest)
			return
		}

		var updateHousePayload models.HouseModel

		err := json.NewDecoder(r.Body).Decode(&updateHousePayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.RemoveHouseImage, updateHousePayload.LoginImages, houseId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")
	}
}
