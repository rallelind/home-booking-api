package controllers

import (
	"database/sql"
	"encoding/json"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"home-booking-api/src/services"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type FamilyResponse struct {
	Users  []clerk.User       `json:"users"`
	Family models.FamilyModel `json:"family"`
}

func CreateFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createFamilyPayload models.FamilyModel

		err := json.NewDecoder(r.Body).Decode(&createFamilyPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var userCount int

		err = db.QueryRow(queries.UserAlreadyPartOfFamilyQuery, createFamilyPayload.Members, createFamilyPayload.HouseId).Scan(&userCount)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if userCount > 0 {
			http.Error(w, "user already part of family", http.StatusBadRequest)
			return
		}

		_, err = db.NamedExec(queries.CreateFamilyQuery, createFamilyPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("successfully created")
	}
}

func UpdateFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var updateFamilyPayload models.FamilyModel

		familyId, ok := mux.Vars(r)["familyId"]

		if !ok {
			http.Error(w, "please provide a family id", http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&updateFamilyPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.UpdateFamilyQuery, updateFamilyPayload.FamilyName, updateFamilyPayload.Members, familyId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")
	}
}

func UpdateFamilyCoverImage(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		familyId, ok := mux.Vars(r)["familyId"]

		if !ok {
			http.Error(w, "please provide a family id", http.StatusBadRequest)
			return
		}

		resp, err := services.UploadImageToCloudinary(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(queries.UpdateFamilyCoverImageQuery, resp.SecureURL, familyId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")
	}
}

func GetFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		familyId, ok := mux.Vars(r)["familyId"]

		if !ok {
			http.Error(w, "missing family id", http.StatusBadRequest)
			return
		}

		var family models.FamilyModel

		err := db.QueryRowx(queries.FindFamilyQuery, familyId).StructScan(&family)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(family)
	}
}

func GetFamilies(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		houseId, ok := mux.Vars(r)["houseId"]

		if !ok {
			http.Error(w, "missing house id", http.StatusBadRequest)
			return
		}

		var families []models.FamilyModel

		err := db.Select(&families, queries.FindFamiliesQuery, houseId)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "no families found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var familyResponse []FamilyResponse

		for i := 0; i < len(families); i++ {
			var family = families[i]
			users, err := services.GetUsersByEmails(family.Members)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			familyResponse = append(familyResponse, FamilyResponse{
				Users:  users,
				Family: family,
			})

		}

		json.NewEncoder(w).Encode(familyResponse)
	}
}

func RemoveFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		familyId, ok := vars["familyId"]

		if !ok {
			http.Error(w, "missing family id", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(queries.DeleteFamilyQuery, familyId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully deleted")

	}
}

func FindFamilyForUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		houseId, ok := vars["houseId"]

		if !ok {
			http.Error(w, "missing house id", http.StatusBadRequest)
			return
		}

		user := services.GetCurrentUser(r)

		var family models.FamilyModel

		err := db.QueryRowx(queries.FindUserFamilyQuery, user.EmailAddresses[0].EmailAddress, houseId).StructScan(&family)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users, err := services.GetUsersByEmails(family.Members)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		familyResponse := FamilyResponse{
			Users:  users,
			Family: family,
		}

		json.NewEncoder(w).Encode(familyResponse)
	}
}
