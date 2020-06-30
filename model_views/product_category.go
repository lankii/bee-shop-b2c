package model_views

type ProductCategory struct {
	Id         int         `json:"id" description:"id"`
	Name       string      `json:"name" description:"名称"`
	Grade      int         `json:"grade" description:"层级"`
	ParentId   interface{} `json:"parent_id" description:"上级分类"`
	Orders     *int        `json:"orders" description:"排序"`
	Promotions *int        `json:"promotions" description:"促销"`
	DeleteFlag int8        `json:"delete_flag" description:"删除标记"`
}
