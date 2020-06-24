package model_views

import (
	"time"
)

type OrderView struct {
	Id                 int       `json:"id" description:"订单ID"`
	Sn                 string    `json:"sn" description:"订单编号"`
	Amount             float64   `json:"amount" description:"订单金额"`
	MemberId           int64     `json:"member_id" description:"会员ID"`
	Username           string    `json:"username" description:"会员用户名"`
	Consignee          string    `json:"consignee" description:"收货人"`
	PaymentMethodId    int64     `json:"payment_method_id" description:"支付方式"`
	PaymentMethodName  string    `json:"payment_method_name" description:"支付方式名称"`
	OrderStatus        int       `json:"order_status" description:"订单状态"`
	OrderStatusName    string    `json:"order_status_name" description:"订单状态名称"`
	ShippingMethodId   int64     `json:"shipping_method_id" description:"配送方式"`
	ShippingMethodName string    `json:"shipping_method_name" description:"配送方式名称"`
	AreaName           string    `json:"area_name" description:"地区名称"`
	Address            string    `json:"address" description:"地址"`
	ZipCode            string    `json:"zip_code" description:"邮编"`
	Phone              string    `json:"phone" description:"电话"`
	CreationDate       time.Time `json:"creation_date" description:"创建日期"`
}
