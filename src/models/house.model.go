package models

import "github.com/lib/pq"

type HouseModel struct {
	Id                  int            `db:"id" json:"id"`
	Address             string         `db:"address" json:"address"`
	HouseName           string         `db:"house_name" json:"house_name"`
	AdminNeedsToApprove bool           `db:"admin_needs_to_approve" json:"admin_needs_to_approve"`
	LoginImages         pq.StringArray `db:"login_images" json:"login_images"`
	HouseAdmins         pq.StringArray `db:"house_admins" json:"house_admins"`
}