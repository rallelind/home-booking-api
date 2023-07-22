package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type FamilyPayload struct {
	Id         int            `db:"id" json:"id"`
	FamilyName string         `db:"family_name" json:"family_name"`
	Members    pq.StringArray `db:"members" json:"members"`
}

func CreateFamily(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createFamilyPayload FamilyPayload

		err := json.NewDecoder(r.Body).Decode(&createFamilyPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertFamily := `
			INSERT INTO families (family_name, members)
			VALUES  (:family_name, :members)
		`

		_, err = db.NamedExec(insertFamily, createFamilyPayload)

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

		var updateFamilyPayload FamilyPayload

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

		sqlUpdate := `
			UPDATE families SET family_name = $1, members = $2
			WHERE id = $3
		`

		_, err = db.Exec(sqlUpdate, updateFamilyPayload.FamilyName, updateFamilyPayload.Members, familyId)

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

		sqlFindFamily := `
			SELECT * FROM families WHERE id = $1
		`

		var family FamilyPayload

		err := db.QueryRowx(sqlFindFamily, familyId).StructScan(&family)

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

		sqlRemoveFamily := `
			DELETE FROM families WHERE id = $1
		`

		_, err := db.Exec(sqlRemoveFamily, familyId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode("successfully deleted")

	}
}