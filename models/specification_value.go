package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SpecificationValue struct {
	Id              int            `orm:"column(id);auto" description:"id"`
	Name            string         `orm:"column(name);size(255)" description:"名称"`
	Image           string         `orm:"column(image);size(255);null" description:"图片"`
	SpecificationId *Specification `orm:"column(specification_id);rel(fk)" description:"规格"`
	Orders          int            `orm:"column(orders);null" description:"排序"`
	CreationDate    time.Time      `orm:"column(creation_date);type(datetime);null" description:"创建日期"`
	CreateBy        string         `orm:"column(create_by);size(20);null" description:"创建人"`
	LastUpdatedBy   string         `orm:"column(last_updated_by);size(20);null" description:"最后修改人"`
	LastUpdatedDate time.Time      `orm:"column(last_updated_date);type(datetime);null" description:"最后修改日期"`
	DeleteFlag      int8           `orm:"column(delete_flag)" description:"删除标记"`
}

func (t *SpecificationValue) TableName() string {
	return "specification_value"
}

func init() {
	orm.RegisterModel(new(SpecificationValue))
}

// AddSpecificationValue insert a new SpecificationValue into database and returns
// last inserted Id on success.
func AddSpecificationValue(m *SpecificationValue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSpecificationValueById retrieves SpecificationValue by Id. Returns error if
// Id doesn't exist
func GetSpecificationValueById(id int) (v *SpecificationValue, err error) {
	o := orm.NewOrm()
	v = &SpecificationValue{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSpecificationValue retrieves all SpecificationValue matches certain condition. Returns empty list if
// no records exist
func GetAllSpecificationValue(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SpecificationValue))
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

	var l []SpecificationValue
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

// UpdateSpecificationValue updates SpecificationValue by Id and returns error if
// the record to be updated doesn't exist
func UpdateSpecificationValueById(m *SpecificationValue) (err error) {
	o := orm.NewOrm()
	v := SpecificationValue{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSpecificationValue deletes SpecificationValue by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSpecificationValue(id int) (err error) {
	o := orm.NewOrm()
	v := SpecificationValue{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SpecificationValue{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
