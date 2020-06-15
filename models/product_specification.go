package models

type ProductSpecification struct {
	Products       *Product       `orm:"column(products);rel(fk)"`
	Specifications *Specification `orm:"column(specifications);rel(fk)"`
}
