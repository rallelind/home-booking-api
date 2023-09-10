
CREATE TABLE IF NOT EXISTS houses(
	id SERIAL PRIMARY KEY,
	address TEXT,
	house_name TEXT,
	admin_needs_to_approve BOOLEAN,
	login_images TEXT[],
	house_admins TEXT[]
)
	

CREATE TABLE IF NOT EXISTS families(
	id SERIAL PRIMARY KEY,
	family_name TEXT,
	members TEXT[],
	cover_image TEXT
)
	

CREATE TABLE IF NOT EXISTS bookings(
	id SERIAL PRIMARY KEY,
	start_date DATE,
	end_date DATE,
	approved BOOLEAN,
	user_booking TEXT
	electricity_used INT
	water_used INT
)
	

CREATE TABLE IF NOT EXISTS booking_posts(
	id SERIAL PRIMARY KEY,
	pictures TEXT[],
	description TEXT
)
	

ALTER TABLE families
	ADD COLUMN IF NOT EXISTS house_id INT NOT NULL,
	ADD CONSTRAINT fk_house FOREIGN KEY(house_id) REFERENCES houses(id)
	

ALTER TABLE bookings 
	ADD COLUMN IF NOT EXISTS house_id INT NOT NULL,
	ADD CONSTRAINT fk_house FOREIGN KEY(house_id) REFERENCES houses(id)

ALTER TABLE booking_posts 
	ADD COLUMN IF NOT EXISTS booking_id INT NOT NULL,
	ADD CONSTRAINT fk_booking FOREIGN KEY(booking_id) REFERENCES bookings(id)


