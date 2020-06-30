package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"strconv"
	"strings"
)

type OrderController struct {
	BaseController
}

// URLMapping ...
func (c *OrderController) URLMapping() {
	c.Mapping("GetAllOrder", c.GetAllOrder)
	c.Mapping("UpdateOrder", c.UpdateOrder)
}

// UpdateOrder ...
// @Title Update Order
// @Description Update Order by some fields
// @Param   id			path    string          true    "The id you want to update"
// @Param	body		body 	models.Product	true	"body for Product content"
// @router /update/:id [put]
// @Success 200 nil
// @Fail    500
func (c *OrderController) UpdateOrder() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	// GetProductById
	currentOrder, err := models.GetOrderById(id)
	if err != nil {
		c.ServerError(err)
		return
	}

	// Unmarshal RequestBody
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &currentOrder)
	if err != nil {
		c.ServerError(err)
		return
	}

	// UpdateProductById
	err = models.UpdateOrderById(currentOrder)
	if err != nil {
		c.ServerError(err)
		return
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, nil)
}

// GetAllOrder ...
// @Title Get All Order List
// @Description Get All Order List
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @Success 200 {object} model_views.Order
// @Failure 500
// @router /all [get]
func (c *OrderController) GetAllOrder() {
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var pageSize int64 = 10
	var pageNumber int64 = 1

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
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
	// pageNumber: 1 (default is 1)
	if v := c.GetString("pageNumber"); v != "" {
		pageNumber, _ = strconv.ParseInt(v, 10, 64)
	}
	// pageSize: 10 (default is 10)
	if v := c.GetString("pageSize"); v != "" {
		pageSize, _ = strconv.ParseInt(v, 10, 64)
	}
	// start position of result set
	offset := (pageNumber - 1) * pageSize
	l, err := models.GetAllOrder(query, nil, sortby, order, offset, pageSize)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 分页并只返回前端需要的字段
	 */
	var pageList []interface{}
	for _, v := range l {
		orderView := model_views.Order{}
		currentMember := &models.Member{}
		order := v.(models.Order)

		currentMember, err = models.GetMemberById(int(order.MemberId))
		if err != nil {
			c.ServerError(err)
			return
		}

		orderView.Id = order.Id
		orderView.Sn = order.Sn
		orderView.Username = currentMember.Username
		orderView.Consignee = order.Consignee
		orderView.PaymentMethodId = order.PaymentMethodId
		orderView.PaymentMethodName = order.PaymentMethodName
		orderView.ShippingMethodId = order.ShippingMethodId
		orderView.ShippingMethodName = order.ShippingMethodName
		orderView.OrderStatus = order.OrderStatus
		orderView.OrderStatusName = common.GetOrderStatusName(order.OrderStatus)
		orderView.AreaName = order.AreaName
		pageList = append(pageList, orderView)
	}

	/**
	 * 查询 Product Count
	 */
	cnt, err := models.GetOrderCount()
	if err != nil {
		c.ServerError(err)
		return
	}

	pages, err := helpers.NewPagination(pageList, int(cnt), int(pageSize), int(pageNumber))
	if err != nil {
		c.ServerError(err)
		return
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, *pages)
}
