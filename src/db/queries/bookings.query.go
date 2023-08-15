package queries

const FindBookingForSpecificDateQuery = `
	SELECT COUNT(*) FROM bookings
	WHERE (start_date, end_date) OVERLAPS ($1, $2) 
	AND approved = true
`

const CreateBookingQuery = `
	INSERT INTO bookings (start_date, end_date, approved, house_id, user_booking)
	VALUES (:start_date, :end_date,
		CASE
			WHEN (SELECT admin_needs_to_approve FROM houses WHERE id = :house_id) = TRUE THEN FALSE
			ELSE TRUE
		END,
	:house_id, :user_booking
	)
`

const DeleteBookingQuery = `
	DELETE FROM bookings WHERE id = $1
`

const UpdateBookingQuery = `
	UPDATE bookings SET approved = $1
	WHERE id = $2
	RETURNING start_date, end_date, approved, id
`

const DeleteOverlappingBookingsQuery = `
	DELETE FROM bookings WHERE (start_date, end_date) OVERLAPS ($1, $2) AND id <> $3
`

const GetBookingsForHouse = `
	SELECT * FROM bookings WHERE house_id = $1
	ORDER BY start_date ASC
`

const GetBookingForCurrentDate = `
	SELECT * FROM bookings WHERE house_id = $1 AND start_date <= $2 AND end_date >= $2
`
