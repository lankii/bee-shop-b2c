package models

type ProductSpecificationValue struct {
	Products            *Product            `orm:"column(products);rel(fk)"`
	SpecificationValues *SpecificationValue `orm:"column(specification_values);rel(fk)"`
}
