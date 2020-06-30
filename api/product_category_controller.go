package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"strconv"
	"strings"
)

type ProductCategoryController struct {
	BaseController
}

// URLMapping ...
func (c *ProductCategoryController) URLMapping() {
	c.Mapping("GetProductCategoryAll", c.GetProductCategoryAll)
	c.Mapping("AddProductCategory", c.AddProductCategory)
}

// @Title AddProductCategory
// @Description create ProductCategory
// @Param name      body string  true   "Category name"
// @Param parent_id body int     false  "Category parent is id"
// @Param orders    body int     false  "Orders"
// @Success 201
// @Failure 500
// @router /add [post]
func (c *ProductCategoryController) AddProductCategory() {
	var name string
	var parentId int
	var orders int
	var promotions int
	var isTop int
	var isMarketable int
	var isShow int

	// name
	if v := c.GetString("name"); v != "" {
		name = v
		if productCategory, _ := models.GetProductCategoryByName(name); productCategory != nil {
			c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, common.Fail, "分类名称不能重复")
			return
		}
	} else {
		c.JsonResult(common.GetHttpStatus("internalServerError"), common.ErrError, common.Fail, nil)
		return
	}
	// parentId
	if v := c.GetString("parent_id"); v != "" {
		parentId, _ = strconv.Atoi(v)
	}
	// orders
	if v := c.GetString("orders"); v != "" {
		orders, _ = strconv.Atoi(v)
	}
	// promotions
	if v := c.GetString("promotions"); v != "" {
		promotions, _ = strconv.Atoi(v)
	}
	// is_top
	if v := c.GetString("is_top"); v != "" {
		isTop, _ = strconv.Atoi(v)
	}
	// is_marketable
	if v := c.GetString("is_marketable"); v != "" {
		isMarketable, _ = strconv.Atoi(v)
	}
	// is_show
	if v := c.GetString("is_show"); v != "" {
		isShow, _ = strconv.Atoi(v)
	}

	var productCategory models.ProductCategory

	/**
	 * 如果存在上级分类，层级则在上级分类之后
	 * 否则为顶级分类
	 */
	if parentId > 0 {
		parentProductCategory, err := models.GetProductCategoryById(parentId)
		if err != nil {
			c.ServerError(err)
			return
		}
		productCategory.ParentId = &parentProductCategory.Id
		productCategory.Grade = parentProductCategory.Grade + 1
		productCategory.TreePath = parentProductCategory.TreePath + strconv.Itoa(parentProductCategory.Id) + ","
	} else {
		productCategory.Grade = 0
		productCategory.TreePath = ","
	}
	productCategory.Name = name
	productCategory.IsTop = int8(isTop)
	productCategory.IsMarketable = int8(isMarketable)
	productCategory.IsShow = int8(isShow)

	/**
	 * 排序
	 */
	if orders > 0 {
		productCategory.Orders = &orders
	}

	id, err := models.AddProductCategory(&productCategory)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 促销
	 */
	if promotions > 0 {
		err := models.AddPromotionProductCategory(promotions, id)
		if err != nil {
			c.ServerError(err)
			return
		}
	}

	c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, common.Success, nil)
}

// @Title UpdateProductCategory
// @Description Update product's category
// @Success 200
// @Failure 500
// @router /update [put]
func (c *ProductCategoryController) UpdateProductCategory() {

}

// @Title GetProductCategoryAll
// @Description Get all product category by some filed
// @router /all [get]
// @Success 200 {object} model_view.ProductCategory
// @Failure 500
func (c *ProductCategoryController) GetProductCategoryAll() {
	/**
	 * 默认查询所有未删除分类
	 */
	var query = make(map[string]string)

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, "fail", "Error: invalid query key/value pair")
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	query["DeleteFlag"] = "0"

	limit, err := models.GetProductCategoryCount(query)
	if err != nil {
		c.ServerError(err)
		return
	}

	l, err := models.GetAllProductCategory(query, nil, nil, nil, 0, limit)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 只返回前端需要的字段
	 */
	var list []interface{}
	for _, v := range l {
		productCategoryView := model_views.ProductCategory{}
		productCategory := v.(models.ProductCategory)

		productCategoryView.Id = productCategory.Id
		productCategoryView.Name = productCategory.Name
		productCategoryView.Orders = productCategory.Orders
		productCategoryView.Grade = productCategory.Grade
		if productCategory.ParentId != nil {
			productCategoryView.ParentId = productCategory.ParentId
		} else {
			productCategoryView.ParentId = nil
		}
		productCategoryView.DeleteFlag = productCategory.DeleteFlag

		list = append(list, productCategoryView)
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, list)
}
