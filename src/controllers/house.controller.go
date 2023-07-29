package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
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
	HouseAdmins         pq.StringArray `db:"house_admins" json:"house_admins"`
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
			INSERT INTO houses (address, house_name, admin_needs_to_approve, login_images, house_admins)
			VALUES (:address, :house_name, :admin_needs_to_approve, :login_images, :house_admins)
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

func GetUserHouses(db *sqlx.DB, clerkClient clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		// ignore the ok because a user will be there as of require session middleware will return 403 else
		sessionClaims, _ := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		user, err := clerkClient.Users().Read(sessionClaims.Claims.Subject)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Did not find user"))
			return
		}

		var housePayload HousePayload

		housesQuery := `

		`

		err = db.QueryRowx(housesQuery, user.EmailAddresses).StructScan(pq.Array(&housePayload))

		if err != nil {
			log.Print(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(housePayload)

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

		w.Header().Set("Content-Type", "application/json")
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
