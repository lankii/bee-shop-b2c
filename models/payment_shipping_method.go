package models

type PaymentShippingMethod struct {
	PaymentMethods  *PaymentMethod  `orm:"column(payment_methods);rel(fk)"`
	ShippingMethods *ShippingMethod `orm:"column(shipping_methods);rel(fk)"`
}
