package models

type ProductParameterValue struct {
	Product           *Product   `orm:"column(product);rel(fk)"`
	ParameterValue    string     `orm:"column(parameter_value);size(255);null"`
	ParameterValueKey *Parameter `orm:"column(parameter_value_key);rel(fk)"`
}
