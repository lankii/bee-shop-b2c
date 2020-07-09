package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"errors"
	"strings"
)

type TagController struct {
	BaseController
}

// @router / [get]
func (c *TagController) GetAllTag() {
	var query = make(map[string]string)
	var limit int64 = 10

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
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

	l, err := models.GetAllTag(query, nil, nil, nil, 0, limit)
	var list []interface{}
	for _, v := range l {
		tagView := model_views.Tag{}
		tag := v.(models.Tag)

		tagView.Id = tag.Id
		tagView.Name = tag.Name
		tagView.Type = tag.Type
		tagView.Orders = tag.Orders
		tagView.Memo = tag.Memo
		tagView.DeleteFlag = tag.DeleteFlag
		list = append(list, tagView)
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, list)
}
