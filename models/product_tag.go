package models

import "github.com/astaxie/beego/orm"

type ProductTag struct {
	Products int `orm:"column(products);"`
	Tags     int `orm:"column(tags);"`
}

// AddProductTag insert a new ProductTag into database
func AddProductTag(products int, tags int) (err error) {
	o := orm.NewOrm()
	sql := "insert into product_tag (products,tags) values (?,?)"

	_, err = o.Raw(sql, products, tags).Exec()
	return
}
