package models

import "time"

type BookingModel struct {
	Id        int       `db:"id" json:"id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
	Approved  bool      `db:"approved" json:"approved"`
	HouseId   int       `db:"house_id" json:"house_id"`
}