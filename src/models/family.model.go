package models

import (
	"github.com/lib/pq"
)

type FamilyModel struct {
	Id         int            `db:"id" json:"id"`
	FamilyName string         `db:"family_name" json:"family_name"`
	Members    pq.StringArray `db:"members" json:"members"`
	HouseId    int            `db:"house_id" json:"house_id"`
	CoverImage *string        `db:"cover_image" json:"cover_image"`
}
