package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminRole struct {
	Admins int64 `orm:"column(admins);" description:"管理员 Id"`
	Roles  int64 `orm:"column(roles);" description:"Role Id"`
}

// AddAdminRole insert a new AdminRole into database and returns
// last inserted Id on success.
func AddAdminRole(admins int64, roles int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO admin_role (admins, roles) VALUES(?, ?)", admins, roles).Exec()
	return
}

// GetAdminRoleIdByAdmins retrieves Roles by Admins. Returns error if
// Roles or IsEnabled doesn't exist
func GetAdminRoleIdByAdmins(admins int64) (adminRole AdminRole, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT admins,roles FROM admin_role WHERE admins = ?", admins).QueryRow(&adminRole)
	return adminRole, err
}
