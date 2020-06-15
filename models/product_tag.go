package models

type ProductTag struct {
	Products *Product `orm:"column(products);rel(fk)"`
	Tags     *Tag     `orm:"column(tags);rel(fk)"`
}
