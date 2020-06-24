package model_views

import "time"

type ProductView struct {
	Id                  int       `json:"id" description:"主键ID"`
	Sn                  string    `json:"sn" description:"编号"`
	Name                string    `json:"name" description:"名称"`
	ProductCategoryId   int       `json:"product_category_id" description:"商品分类ID"`
	ProductCategoryName string    `json:"product_category_name" description:"商品分类名称"`
	Price               float64   `json:"price" description:"售价"`
	Stock               int       `json:"stock" description:"库存"`
	IsMarketable        int8      `json:"is_marketable" description:"是否上架"`
	IsList              int8      `json:"is_list" description:"是否列出"`
	IsTop               int8      `json:"is_top" description:"是否置顶"`
	CreationDate        time.Time `json:"creation_date" description:"创建日期"`
}
