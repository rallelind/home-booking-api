package tables

import (
	"database/sql"
	"log"
)

func SetupDatabase(db *sql.DB) {

	houseTableQuery := `
		CREATE TABLE IF NOT EXISTS houses(
			id SERIAL PRIMARY KEY,
			address TEXT,
			house_name TEXT,
			admin_needs_to_approve BOOLEAN,
			login_images TEXT[],
		)
	`

	familyTableQuery := `
		CREATE TABLE IF NOT EXISTS families(
			id SERIAL PRIMARY KEY,
			family_name TEXT,
			members TEXT[],
		)
	`

	bookingTableQuery := `
		CREATE TABLE IF NOT EXISTS bookings(
			id SERIAL PRIMARY KEY,
			start_date DATE,
			end_date DATE,
			approved BOOLEAN,
		)
	`

	bookingPostTable := `
		CREATE TABLE IF NOT EXISTS booking_posts(
			id SERIAL PRIMARY KEY,
			pictures TEXT[],
			description TEXT,
		)
	`

	familyRelation := `
		ALTER TABLE families
			ADD COLUMN IF NOT EXISTS house_id INT NOT NULL
			ADD CONSTRAINT fk_house FOREIGN KEY(house_id) REFERENCES houses(id)
	`

	bookingRelation := `
		ALTER TABLE bookings 
			ADD COLUMN IF NOT EXISTS house_id INT NOT NULL
			ADD CONSTRAINT fk_house FOREIGN KEY(house_id) REFERENCES houses(id)

	`

	bookingPostRelation := `
		ALTER TABLE booking_posts 
			ADD COLUMN IF NOT EXISTS booking_id INT NOT NULL
			ADD CONSTRAINT fk_booking FOREIGN KEY(booking_id) REFERENCES bookings(id)
	`

	_, err := db.Exec(houseTableQuery, familyTableQuery, bookingTableQuery, bookingPostTable, familyRelation, bookingRelation, bookingPostRelation)
	
	if err != nil {
		log.Fatal(err)
	}
}
