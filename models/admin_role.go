package models

type AdminRole struct {
	Admins *Admin `orm:"column(admins);rel(fk)"`
	Roles  *Role  `orm:"column(roles);rel(fk)" description:"角色"`
}
