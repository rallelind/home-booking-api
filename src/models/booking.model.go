package models

import "time"

type BookingModel struct {
	Id           int       `db:"id" json:"id"`
	StartDate    time.Time `db:"start_date" json:"start_date"`
	EndDate      time.Time `db:"end_date" json:"end_date"`
	Approved     bool      `db:"approved" json:"approved"`
	HouseId      int       `db:"house_id" json:"house_id"`
	UserBooking  string    `db:"user_booking" json:"user_booking"`
	WaterUsed    int       `db:"water_used" json:"water_used"`
	ElectricUsed int       `db:"electric_used" json:"electric_used"`
}
