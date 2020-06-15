package models

type PromotionCoupon struct {
	Promotions *Promotion `orm:"column(promotions);rel(fk)"`
	Coupons    *Coupon    `orm:"column(coupons);rel(fk)"`
}
