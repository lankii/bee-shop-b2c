package models

type AttributeOption struct {
	Attribute *Attribute `orm:"column(attribute);rel(fk)"`
	Options   string     `orm:"column(options);size(255);null" description:"可选项"`
}
