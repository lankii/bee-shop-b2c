package models

type PromotionProduct struct {
	Promotions *Promotion `orm:"column(promotions);rel(fk)"`
	Products   *Product   `orm:"column(products);rel(fk)"`
}
