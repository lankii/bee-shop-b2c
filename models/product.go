package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id                int              `orm:"column(id);auto" description:"id"`
	Sn                string           `orm:"column(sn);size(100)" description:"编号"`
	Name              string           `orm:"column(name);size(255)" description:"名称"`
	FullName          string           `orm:"column(full_name);size(255)" description:"全称"`
	Price             float64          `orm:"column(price);digits(21);decimals(6)" description:"销售价"`
	Cost              float64          `orm:"column(cost);null;digits(21);decimals(6)" description:"成本价"`
	MarketPrice       float64          `orm:"column(market_price);digits(21);decimals(6)" description:"市场价"`
	Image             string           `orm:"column(image);size(255);null" description:"展示图片"`
	Unit              string           `orm:"column(unit);size(255);null" description:"单位"`
	Weight            int              `orm:"column(weight);null" description:"重量"`
	Stock             int              `orm:"column(stock);null" description:"库存"`
	AllocatedStock    int              `orm:"column(allocated_stock)" description:"已分配库存"`
	StockMemo         string           `orm:"column(stock_memo);size(255);null" description:"库存备注"`
	Point             int64            `orm:"column(point)" description:"赠送积分"`
	IsMarketable      int8             `orm:"column(is_marketable)" description:"是否上架"`
	IsList            int8             `orm:"column(is_list)" description:"是否列出"`
	IsTop             int8             `orm:"column(is_top)" description:"是否置顶"`
	IsGift            int8             `orm:"column(is_gift)" description:"是否为赠品"`
	Introduction      string           `orm:"column(introduction);null" description:"介绍"`
	Memo              string           `orm:"column(memo);size(255);null" description:"备注"`
	Keyword           string           `orm:"column(keyword);size(255);null" description:"搜索关键词"`
	SeoTitle          string           `orm:"column(seo_title);size(255);null" description:"页面标题"`
	SeoKeywords       string           `orm:"column(seo_keywords);size(255);null" description:"页面关键词"`
	SeoDescription    string           `orm:"column(seo_description);size(255);null" description:"页面描述"`
	Score             float32          `orm:"column(score)" description:"评分"`
	TotalScore        int64            `orm:"column(total_score)" description:"总评分"`
	ScoreCount        int64            `orm:"column(score_count)" description:"评分数"`
	Hits              int64            `orm:"column(hits)" description:"点击数"`
	WeekHits          int64            `orm:"column(week_hits)" description:"周点击数"`
	MonthHits         int64            `orm:"column(month_hits)" description:"月点击数"`
	Sales             int64            `orm:"column(sales)" description:"销量"`
	WeekSales         int64            `orm:"column(week_sales)" description:"周销量"`
	MonthSales        int64            `orm:"column(month_sales)" description:"月销量"`
	WeekHitsDate      time.Time        `orm:"column(week_hits_date);auto_now_add;type(datetime)" description:"周点击数更新日期"`
	MonthHitsDate     time.Time        `orm:"column(month_hits_date);auto_now_add;type(datetime)" description:"月点击数更新日期"`
	WeekSalesDate     time.Time        `orm:"column(week_sales_date);auto_now_add;type(datetime)" description:"周销量更新日期"`
	MonthSalesDate    time.Time        `orm:"column(month_sales_date);auto_now_add;type(datetime)" description:"月销量更新日期"`
	AttributeValue0   string           `orm:"column(attribute_value0);size(255);null" description:"商品属性值0"`
	AttributeValue1   string           `orm:"column(attribute_value1);size(255);null" description:"商品属性值1"`
	AttributeValue2   string           `orm:"column(attribute_value2);size(255);null" description:"商品属性值2"`
	AttributeValue3   string           `orm:"column(attribute_value3);size(255);null" description:"商品属性值3"`
	AttributeValue4   string           `orm:"column(attribute_value4);size(255);null" description:"商品属性值4"`
	AttributeValue5   string           `orm:"column(attribute_value5);size(255);null" description:"商品属性值5"`
	AttributeValue6   string           `orm:"column(attribute_value6);size(255);null" description:"商品属性值6"`
	AttributeValue7   string           `orm:"column(attribute_value7);size(255);null" description:"商品属性值7"`
	AttributeValue8   string           `orm:"column(attribute_value8);size(255);null" description:"商品属性值8"`
	AttributeValue9   string           `orm:"column(attribute_value9);size(255);null" description:"商品属性值9"`
	AttributeValue10  string           `orm:"column(attribute_value10);size(255);null" description:"商品属性值10"`
	AttributeValue11  string           `orm:"column(attribute_value11);size(255);null" description:"商品属性值11"`
	AttributeValue12  string           `orm:"column(attribute_value12);size(255);null" description:"商品属性值12"`
	AttributeValue13  string           `orm:"column(attribute_value13);size(255);null" description:"商品属性值13"`
	AttributeValue14  string           `orm:"column(attribute_value14);size(255);null" description:"商品属性值14"`
	AttributeValue15  string           `orm:"column(attribute_value15);size(255);null" description:"商品属性值15"`
	AttributeValue16  string           `orm:"column(attribute_value16);size(255);null" description:"商品属性值16"`
	AttributeValue17  string           `orm:"column(attribute_value17);size(255);null" description:"商品属性值17"`
	AttributeValue18  string           `orm:"column(attribute_value18);size(255);null" description:"商品属性值18"`
	AttributeValue19  string           `orm:"column(attribute_value19);size(255);null" description:"商品属性19"`
	BrandId           *Brand           `orm:"column(brand_id);rel(fk)" description:"品牌"`
	GoodsId           *Goods           `orm:"column(goods_id);rel(fk)" description:"货品"`
	ProductCategoryId *ProductCategory `orm:"column(product_category_id);rel(fk)" description:"商品分类"`
	CreateBy          string           `orm:"column(create_by);size(20);null" description:"创建人"`
	CreationDate      time.Time        `orm:"column(creation_date);auto_now_add;type(datetime);null" description:"创建日期"`
	LastUpdatedBy     string           `orm:"column(last_updated_by);size(20);null" description:"最后修改人"`
	LastUpdatedDate   time.Time        `orm:"column(last_updated_date);auto_now;type(datetime);null" description:"最后修改日期"`
	DeleteFlag        int8             `orm:"column(delete_flag)" description:"删除标记"`
}

func (t *Product) TableName() string {
	return "product"
}

func init() {
	orm.RegisterModel(new(Product))
}

// AddProduct insert a new Product into database and returns
// last inserted Id on success.
func AddProduct(m *Product) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductById retrieves Product by Id. Returns error if
// Id doesn't exist
func GetProductById(id int) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetProductCount calculate Product Count. Returns error if
// Table doesn't exist
func GetProductCount() (cnt int64, err error) {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable(new(Product)).Count(); err == nil {
		return cnt, nil
	}
	return 0, err
}

// GetAllProduct retrieves all Product matches certain condition. Returns empty list if
// no records exist
func GetAllProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
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

	var l []Product
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

// UpdateProduct updates Product by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductById(m *Product) (err error) {
	o := orm.NewOrm()
	v := Product{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProduct deletes Product by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProduct(id int) (err error) {
	o := orm.NewOrm()
	v := Product{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Product{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
