package controllers

import (
	"encoding/json"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func CreateFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createFamilyPayload models.FamilyModel

		err := json.NewDecoder(r.Body).Decode(&createFamilyPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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

		vars := mux.Vars(r)
		familyId, ok := vars["familyId"]

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

func GetFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		familyId, ok := vars["familyId"]

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
