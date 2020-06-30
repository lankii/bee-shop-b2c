package models

import "github.com/astaxie/beego/orm"

type PromotionProductCategory struct {
	Promotions        int   `orm:"column(promotions);" description:"促销Id"`
	ProductCategories int64 `orm:"column(product_categories);" description:"商品分类Id"`
}

// AddPromotionProductCategory insert a new Admin into database and returns
// last inserted Id on success.
func AddPromotionProductCategory(promotions int, productCategories int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO promotion_product_category (promotions, product_categories) VALUES(?, ?)", promotions, productCategories).Exec()
	return
}
