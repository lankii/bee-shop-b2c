package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id                 int       `orm:"column(id);auto" description:"主键_id"`
	Sn                 string    `orm:"column(sn);size(100)" description:"订单编号"`
	OrderStatus        int       `orm:"column(order_status)" description:"订单状态"`
	PaymentStatus      int       `orm:"column(payment_status)" description:"支付状态"`
	ShippingStatus     int       `orm:"column(shipping_status)" description:"配送状态"`
	Fee                float64   `orm:"column(fee);digits(21);decimals(6)" description:"支付手续费"`
	Freight            float64   `orm:"column(freight);digits(21);decimals(6)" description:"运费"`
	PromotionDiscount  float64   `orm:"column(promotion_discount);digits(21);decimals(6)" description:"促销折扣"`
	CouponDiscount     float64   `orm:"column(coupon_discount);digits(21);decimals(6)" description:"优惠券折扣"`
	OffsetAmount       float64   `orm:"column(offset_amount);digits(21);decimals(6)" description:"调整金额"`
	AmountPaid         float64   `orm:"column(amount_paid);digits(21);decimals(6)" description:"已付金额"`
	Point              int64     `orm:"column(point)" description:"赠送积分"`
	Consignee          string    `orm:"column(consignee);size(255)" description:"收货人"`
	AreaName           string    `orm:"column(area_name);size(255)" description:"地区名称"`
	Address            string    `orm:"column(address);size(255)" description:"地址"`
	ZipCode            string    `orm:"column(zip_code);size(255)" description:"邮编"`
	Phone              string    `orm:"column(phone);size(255)" description:"电话"`
	IsInvoice          int8      `orm:"column(is_invoice)" description:"是否开据发票"`
	InvoiceTitle       string    `orm:"column(invoice_title);size(255);null" description:"发票标题"`
	Tax                float64   `orm:"column(tax);digits(21);decimals(6)" description:"税金"`
	Memo               string    `orm:"column(memo);size(255);null" description:"附言"`
	Promotion          string    `orm:"column(promotion);size(255);null" description:"促销"`
	Expire             time.Time `orm:"column(expire);type(datetime);null" description:"到期时间"`
	LockExpire         time.Time `orm:"column(lock_expire);type(datetime);null" description:"锁定到期时间"`
	IsAllocatedStock   int8      `orm:"column(is_allocated_stock)" description:"是否已分配库存"`
	PaymentMethodName  string    `orm:"column(payment_method_name);size(255)" description:"支付方式名称"`
	ShippingMethodName string    `orm:"column(shipping_method_name);size(255)" description:"配送方式名称"`
	AreaId             int64     `orm:"column(area_id);null" description:"地区"`
	PaymentMethodId    int64     `orm:"column(payment_method_id);null" description:"支付方式"`
	ShippingMethodId   int64     `orm:"column(shipping_method_id);null" description:"配送方式"`
	OperatorId         int64     `orm:"column(operator_id);null" description:"操作员"`
	MemberId           int64     `orm:"column(member_id)" description:"会员"`
	CouponCode         int64     `orm:"column(coupon_code);null" description:"优惠码"`
	ServiceDate        time.Time `orm:"column(service_date);type(datetime);null" description:"服务时间"`
	CreateBy           string    `orm:"column(create_by);size(20);null" description:"创建人"`
	CreationDate       time.Time `orm:"column(creation_date);auto_now_add;type(datetime);null" description:"创建日期"`
	LastUpdatedBy      string    `orm:"column(last_updated_by);size(20);null" description:"最后修改人"`
	LastUpdatedDate    time.Time `orm:"column(last_updated_date);auto_now;type(datetime);null" description:"最后修改日期"`
	DeleteFlag         int8      `orm:"column(delete_flag)" description:"删除标记"`
}

func (t *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}

// AddOrder insert a new Order into database and returns
// last inserted Id on success.
func AddOrder(m *Order) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderById retrieves Order by Id. Returns error if
// Id doesn't exist
func GetOrderById(id int) (v *Order, err error) {
	o := orm.NewOrm()
	v = &Order{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetProductCount calculate Product Count. Returns error if
// Table doesn't exist
func GetOrderCount() (cnt int64, err error) {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable(new(Order)).Count(); err == nil {
		return cnt, nil
	}
	return 0, err
}

// GetAllOrder retrieves all Order matches certain condition. Returns empty list if
// no records exist
func GetAllOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Order))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Order
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateOrder updates Order by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderById(m *Order) (err error) {
	o := orm.NewOrm()
	v := Order{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrder deletes Order by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrder(id int) (err error) {
	o := orm.NewOrm()
	v := Order{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Order{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
