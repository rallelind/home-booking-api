package tables

import "database/sql"

func SetupDatabase(db *sql.DB) {

	houseTableQuery := `
		CREATE TABLE IF NOT EXISTS houses
	`

	familyTableQuery := `
		CREATE TABLE IF NOT EXISTS families
	`

	bookingTableQuery := `
		CREATE TABLE IF NOT EXISTS bookings
	`

	bookingPostTable := `
		CREATE TABLE IF NOT EXISTS booking_posts
	`

	houseRelation := `
		ALTER TABLE houses ADD COLUMN IF NOT EXISTS
	`

	familyRelation := `
		ALTER TABLE families ADD COLUMN IF NOT EXISTS
	`

	bookingRelation := `
		ALTER TABLE bookings ADD COLUMN IF NOT EXISTS
	`

	bookingPostRelation := `
		ALTER TABLE booking_posts ADD COLUMN IF NOT EXISTS
	`
	
	


}