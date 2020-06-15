package models

type PromotionProductCategory struct {
	Promotions        *Promotion       `orm:"column(promotions);rel(fk)" description:"促销Id"`
	ProductCategories *ProductCategory `orm:"column(product_categories);rel(fk)" description:"商品分类"`
}
