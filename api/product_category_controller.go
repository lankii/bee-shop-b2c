package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"github.com/astaxie/beego/utils"
	"strconv"
	"strings"
)

type ProductCategoryController struct {
	BaseController
}

// URLMapping ...
func (c *ProductCategoryController) URLMapping() {
	c.Mapping("GetProductCategoryAll", c.GetAllProductCategory)
	c.Mapping("AddProductCategory", c.AddProductCategory)
	c.Mapping("GetOneProductCategory", c.GetOneProductCategory)
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
	var promotions []string
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
	if v := c.GetString("parentId"); v != "" {
		parentId, _ = strconv.Atoi(v)
	}
	// orders
	if v := c.GetString("orders"); v != "" {
		orders, _ = strconv.Atoi(v)
	}
	// promotions
	if v := c.GetString("promotions"); v != "" {
		promotions = strings.Split(v, ",")
	}
	// is_top
	if v := c.GetString("isTop"); v != "" {
		isTop, _ = strconv.Atoi(v)
	}
	// is_marketable
	if v := c.GetString("isMarketable"); v != "" {
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
	if len(promotions) > 0 {
		for _, p := range promotions {
			p, _ := strconv.Atoi(p)
			err := models.AddPromotionProductCategory(p, int(id))
			if err != nil {
				c.ServerError(err)
				return
			}
		}

	}

	c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, common.Success, nil)
}

// @Title GetOne
// @Description get ProductCategory by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProductCategory
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductCategoryController) GetOneProductCategory() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProductCategoryById(id)

	if err != nil {
		c.JsonResult(common.GetHttpStatus("forbidden"), common.ErrError, common.Fail, ":id is empty")
		return
	}

	productCategoryView := model_views.ProductCategory{}
	productCategoryView.Id = v.Id
	productCategoryView.Name = v.Name
	productCategoryView.Orders = v.Orders
	productCategoryView.Grade = v.Grade
	if v.ParentId != nil {
		parentId := *v.ParentId
		productCategoryView.ParentId = parentId
		parentProductCategory, err := models.GetProductCategoryById(parentId)
		if err != nil {
			c.ServerError(err)
			return
		}
		productCategoryView.ParentName = parentProductCategory.Name
	} else {
		productCategoryView.ParentId = nil
	}
	productCategoryView.IsMarketable = v.IsMarketable
	productCategoryView.IsTop = v.IsTop
	productCategoryView.IsShow = v.IsShow

	/**
	 * 查询对应的 Promotion 的结果集
	 */
	if l, err := models.GetPromotionProductCategoryByProductCategory(v.Id); err == nil {
		productCategoryView.Promotions = l
	}

	productCategoryView.DeleteFlag = v.DeleteFlag

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, productCategoryView)
}

// @Title GetProductCategoryAll
// @Description Get all product category by some filed
// @router /all [get]
// @Success 200 {object} model_view.ProductCategory
// @Failure 500
func (c *ProductCategoryController) GetAllProductCategory() {
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

// @Title UpdateProductCategory
// @Description Update product's category
// @Success 204
// @Failure 500
// @router /update [put]
func (c *ProductCategoryController) UpdateProductCategory() {
	var id int
	var name string
	var parentId int
	var orders int
	var promotions []string
	var isTop int
	var isMarketable int
	var isShow int

	// parentId
	if v := c.GetString("id"); v != "" {
		id, _ = strconv.Atoi(v)
	}
	// name
	if v := c.GetString("name"); v != "" {
		name = v

	} else {
		c.JsonResult(common.GetHttpStatus("internalServerError"), common.ErrError, common.Fail, nil)
		return
	}
	// parentId
	if v := c.GetString("parentId"); v != "" {
		parentId, _ = strconv.Atoi(v)
	}
	// orders
	if v := c.GetString("orders"); v != "" {
		orders, _ = strconv.Atoi(v)
	}
	// promotions
	if v := c.GetString("promotions"); v != "" {
		promotions = strings.Split(v, ",")
	}
	// is_top
	if v := c.GetString("isTop"); v != "" {
		isTop, _ = strconv.Atoi(v)
	}
	// is_marketable
	if v := c.GetString("isMarketable"); v != "" {
		isMarketable, _ = strconv.Atoi(v)
	}
	// is_show
	if v := c.GetString("isShow"); v != "" {
		isShow, _ = strconv.Atoi(v)
	}

	var productCategory models.ProductCategory

	productCategory.Id = id
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

	err := models.UpdateProductCategoryById(&productCategory)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 促销
	 */
	var productCategoryId = productCategory.Id
	if promotionsArr, err := models.GetPromotionProductCategoryByProductCategory(productCategoryId); err == nil {
		// 当 promotionsArr > promotions参数时，需要删除
		if len(promotionsArr) > len(promotions) {
			for _, p := range promotionsArr {
				if in := utils.InSlice(p, promotions); in == false {
					p, _ := strconv.Atoi(p)
					_ = models.DeletePromotionProductCategory(p, productCategoryId)
				}
			}
		} else if len(promotionsArr) < len(promotions) { // 当 promotionsArr > promotions参数时，需要新增
			for _, p := range promotions {
				if in := utils.InSlice(p, promotionsArr); in == false {
					p, _ := strconv.Atoi(p)
					_ = models.AddPromotionProductCategory(p, productCategoryId)
				}
			}
		}
	}

	c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, common.Success, nil)
}
