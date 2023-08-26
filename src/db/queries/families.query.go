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

const FindUserFamilyQuery = `
	SELECT * FROM families WHERE $1 = ANY(members) AND house_id = $2 
`

const DeleteFamilyQuery = `
	DELETE FROM families WHERE id = $1
`

const FindFamiliesQuery = `
	SELECT * FROM families WHERE house_id = $1
`

const UpdateFamilyCoverImageQuery = `
	UPDATE families SET cover_image = $1 WHERE id = $2
`

const UserAlreadyPartOfFamilyQuery = `
	SELECT COUNT(*) FROM families            
	WHERE EXISTS (
		SELECT 1 FROM unnest(members) AS m
		WHERE m = ANY($1::TEXT[])
		AND house_id = $2
	);
`
