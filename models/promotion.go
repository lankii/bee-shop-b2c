package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Promotion struct {
	Id              int       `orm:"column(id);auto"`
	Name            string    `orm:"column(name);size(255)" description:"名称"`
	Title           string    `orm:"column(title);size(255)" description:"标题"`
	BeginDate       time.Time `orm:"column(begin_date);type(datetime);null" description:"起始日期"`
	EndDate         time.Time `orm:"column(end_date);type(datetime);null" description:"结束日期"`
	MinimumQuantity int       `orm:"column(minimum_quantity);null" description:"最小商品数量"`
	MaximumQuantity int       `orm:"column(maximum_quantity);null" description:"最大商品数量"`
	MinimumPrice    float64   `orm:"column(minimum_price);null;digits(21);decimals(6)" description:"最小商品价格"`
	MaximumPrice    float64   `orm:"column(maximum_price);null;digits(21);decimals(6)" description:"最大商品价格"`
	PriceExpression string    `orm:"column(price_expression);size(255);null" description:"价格运算表达式"`
	PointExpression string    `orm:"column(point_expression);size(255);null" description:"积分运算表达式"`
	IsFreeShipping  int8      `orm:"column(is_free_shipping)" description:"是否免运费"`
	IsCouponAllowed int8      `orm:"column(is_coupon_allowed)" description:"是否允许使用优惠券"`
	Introduction    string    `orm:"column(introduction);null" description:"介绍"`
	Orders          int       `orm:"column(orders);null" description:"排序"`
	CreateBy        string    `orm:"column(create_by);size(20);null" description:"创建人"`
	CreationDate    time.Time `orm:"column(creation_date);type(datetime);null" description:"创建日期"`
	LastUpdatedBy   string    `orm:"column(last_updated_by);size(20);null" description:"最后修改人"`
	LastUpdatedDate time.Time `orm:"column(last_updated_date);type(datetime);null" description:"最后修改日期"`
	DeleteFlag      int8      `orm:"column(delete_flag)" description:"删除标记"`
}

func (t *Promotion) TableName() string {
	return "promotion"
}

func init() {
	orm.RegisterModel(new(Promotion))
}

// AddPromotion insert a new Promotion into database and returns
// last inserted Id on success.
func AddPromotion(m *Promotion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPromotionById retrieves Promotion by Id. Returns error if
// Id doesn't exist
func GetPromotionById(id int) (v *Promotion, err error) {
	o := orm.NewOrm()
	v = &Promotion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPromotion retrieves all Promotion matches certain condition. Returns empty list if
// no records exist
func GetAllPromotion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Promotion))
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

	var l []Promotion
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

// UpdatePromotion updates Promotion by Id and returns error if
// the record to be updated doesn't exist
func UpdatePromotionById(m *Promotion) (err error) {
	o := orm.NewOrm()
	v := Promotion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePromotion deletes Promotion by Id and returns error if
// the record to be deleted doesn't exist
func DeletePromotion(id int) (err error) {
	o := orm.NewOrm()
	v := Promotion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Promotion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
