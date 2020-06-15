package models

type PermissionRole struct {
	Permissions *Permission `orm:"column(permissions);rel(fk)"`
	Roles       *Role       `orm:"column(roles);rel(fk)" description:"角色"`
}
