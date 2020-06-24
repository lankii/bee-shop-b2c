package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
)

type ProductCategoryController struct {
	BaseController
}

// URLMapping ...
func (c *ProductCategoryController) URLMapping() {
	c.Mapping("GetProductCategoryAll", c.GetProductCategoryAll)
}

// GetProductCategoryAll ...
// @Title Get All Product Category List
// @Description Get All Product Category List
// @router /all [get]
// @Success 200 {object} models.ProductCategory
// @Failure 500
func (c *ProductCategoryController) GetProductCategoryAll() {
	/**
	 * 查询所有未删除分类
	 */
	var query = make(map[string]string)

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
		if productCategory.ParentId != nil {
			productCategoryView.ParentId = productCategory.ParentId.Id
		} else {
			productCategoryView.ParentId = nil
		}
		productCategoryView.DeleteFlag = productCategory.DeleteFlag

		list = append(list, productCategoryView)
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", list)
}
