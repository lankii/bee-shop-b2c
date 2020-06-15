package models

type ProductImage struct {
	Title     string   `orm:"column(title);size(255);null" description:"标题"`
	ProductId *Product `orm:"column(product_id);rel(fk)" description:"商品"`
	Source    string   `orm:"column(source);size(255);null" description:"原图片"`
	Large     string   `orm:"column(large);size(255);null" description:"大图片"`
	Medium    string   `orm:"column(medium);size(255);null" description:"中图片"`
	Thumbnail string   `orm:"column(thumbnail);size(255);null" description:"缩略图"`
	Orders    int      `orm:"column(orders);null" description:"排序"`
}
