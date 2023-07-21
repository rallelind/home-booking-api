package controllers

import "github.com/lib/pq"

type FamilyPayload struct {
	Id         int            `db:"id" json:"id"`
	FamilyName string         `db:"family_name" json:"family_name"`
	Members    pq.StringArray `db:"members" json:"members"`
}

func CreateFamily() {}

func UpdateFamily() {}

func GetFamily() {}

func RemoveFamily() {}