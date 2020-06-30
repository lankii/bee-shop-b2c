package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
)

type PromotionController struct {
	BaseController
}

// URLMapping ...
func (c *PromotionController) URLMapping() {
	c.Mapping("GetAllPromotion", c.GetAllPromotion)
}

// @Title GetAllPromotion
// @Description Get all promotion by some filed
// @router /all [get]
func (c *PromotionController) GetAllPromotion() {
	var query map[string]string
	query = make(map[string]string)

	/**
	 * 查询未删除的促销
	 */
	query["delete_flag"] = "0"
	limit, err := models.GetPromotionCount(query)
	if err != nil {
		c.ServerError(err)
		return
	}

	l, err := models.GetAllPromotion(query, nil, nil, nil, 0, limit)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 只返回前端需要的字段
	 */
	var list []interface{}
	for _, v := range l {
		promotionView := model_views.Promotion{}
		promotion := v.(models.Promotion)

		promotionView.Id = promotion.Id
		promotionView.Name = promotion.Name
		promotionView.Title = promotion.Title
		promotionView.DeleteFlag = promotion.DeleteFlag
		list = append(list, promotionView)
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, list)
}
