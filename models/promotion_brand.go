package models

type PromotionBrand struct {
	Promotions *Promotion `orm:"column(promotions);rel(fk)"`
	Brands     *Brand     `orm:"column(brands);rel(fk)"`
}
