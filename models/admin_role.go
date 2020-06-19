package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminRole struct {
	Admins int64
	Roles  int64
}

func TableName() string {
	return "admin_role"
}

// AddAdmin insert a new Admin into database and returns
// last inserted Id on success.
func AddAdminRole(admin int64, role int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO"+TableName()+"(admins, roles) VALUES(?, ?)", admin, role).Exec()
	return
}

// GetAdminRoleIdByAdmins retrieves Roles by Admins. Returns error if
// Roles or IsEnabled doesn't exist
func GetAdminRoleIdByAdmins(admins int64) (adminRole AdminRole, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT admins,roles FROM "+TableName()+" WHERE admins = ?", admins).QueryRow(&adminRole)
	return adminRole, err
}
