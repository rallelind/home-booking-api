package queries

// Inserts a struct from json payload
const CreateFamilyQuery = `
	INSERT INTO families (family_name, members, house_id)
	VALUES (:family_name, :members, :house_id)
`

const UpdateFamilyQuery = `
	UPDATE families SET family_name = $1, members = $2
	WHERE id = $3
`

const FindFamilyQuery = `
	SELECT * FROM families WHERE id = $1
`

const DeleteFamilyQuery = `
	DELETE FROM families WHERE id = $1
`

const FindFamiliesQuery = `
	SELECT * FROM families WHERE house_id = $1
`