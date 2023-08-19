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

const FindHouseQuery = `
	SELECT * FROM houses WHERE id = $1
`

const RemoveHouseQuery = `
	DELETE FROM houses WHERE id = $1
`

const RemoveHouseImage = `
	UPDATE houses SET login_images = array_remove(login_images, $1) WHERE id = $2
`

const RemoveHouseAdmin = `
	UPDATE houses SET house_admins = array_remove(house_admins, $1) WHERE id = $2
`

const AddHouseImages = `
	UPDATE houses SET login_images = array_append(login_images, $1) WHERE id = $2
`

const AddHouseAdmin = `
	UPDATE houses SET house_admins = $1 WHERE id = $2
`

const ChangeAdminNeedsToApprove = `
	UPDATE houses SET admin_needs_to_approve = $1 WHERE id = $2
`
