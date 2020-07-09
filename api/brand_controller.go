package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"strings"
)

type BrandController struct {
	BaseController
}

// @router / [get]
func (c *BrandController) GetAllBrand() {
	var query = make(map[string]string)
	var limit int64

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
	/**
	 * 默认查询未删除
	 */
	query["DeleteFlag"] = "0"
	limit, err := models.GetBrandtCount(query)
	if err != nil {
		c.ServerError(err)
		return
	}

	l, err := models.GetAllBrand(query, nil, nil, nil, 0, limit)
	if err != nil {
		c.ServerError(err)
		return
	}

	var list []interface{}
	for _, v := range l {
		brandView := model_views.Brand{}
		brand := v.(models.Brand)

		brandView.Id = brand.Id
		brandView.Name = brand.Name
		brandView.Type = brand.Type
		brandView.Logo = brand.Logo
		brandView.Orders = brand.Orders
		brandView.DeleteFlag = brand.DeleteFlag
		list = append(list, brandView)
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, list)
}
