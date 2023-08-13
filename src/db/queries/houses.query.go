package queries

// Inserts a struct from json payload
const CreateHouseQuery = `
	INSERT INTO houses (address, house_name, admin_needs_to_approve, login_images, house_admins)
	VALUES (:address, :house_name, :admin_needs_to_approve, :login_images, :house_admins)
`

const FindUserHousesQuery = `
	SELECT id, address, house_name, admin_needs_to_approve, login_images, house_admins
	FROM houses
	WHERE $1 = ANY(house_admins)

	UNION

	SELECT h.id, h.address, h.house_name, h.admin_needs_to_approve, h.login_images, h.house_admins
	FROM houses h
	INNER JOIN families f ON h.id = f.house_id
	WHERE $1 = ANY(f.members)
`


const UserHouseAdminQuery = `
	SELECT * FROM houses WHERE $1 = ANY(house_admins)
`

const UserPartOfHouseQuery = `
	SELECT * FROM houses h INNER JOIN families f ON h.id = f.house_id WHERE $1 = ANY(f.members) AND $2 = f.house_id
`

const UpdateHouseQuery = `
	UPDATE houses SET house_name = $1, admin_needs_to_approve = $2, login_images = $3
	WHERE id = $4
`

const FindHouseQuery = `
	SELECT * FROM houses WHERE id = $1
`

const RemoveHouseQuery = `
	DELETE FROM houses WHERE id = $1
`
