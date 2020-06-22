package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"strconv"
	"strings"
)

type ProductController struct {
	BaseController
}

// URLMapping ...
func (c *ProductController) URLMapping() {
	c.Mapping("AddProduct", c.AddProduct)
	c.Mapping("GetProducts", c.GetProducts)
	c.Mapping("UpdateProduct", c.UpdateProduct)
}

// AddProduct ...
// @Title Add Product
// @Description create Product
// @Param	body	   body 	models.Product	true	"body for Product content"
// @Success 201 {int}  models.Product.Id
// @Failure 500
// @router /add [post]
func (c *ProductController) AddProduct() {
	var v models.Product

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.ServerError(err)
		return
	}

	id, err := models.AddProduct(&v)
	if err != nil {
		c.ServerError(err)
		return
	}
	c.JsonResult(common.GetHttpStatus("created"),common.ErrOK, "success", id)
}

// UpdateProduct...
// @Title Update Product
// @Description Update Product by some fields
// @Param   id			path    string          true    "The id you want to update"
// @Param	body		body 	models.Product	true	"body for Product content"
// @router /update/:id [put]
// @Success 200 nil
// @Fail    500
func (c *ProductController) UpdateProduct() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	// GetProductById
	currentProduct, err := models.GetProductById(id)
	if err != nil {
		c.ServerError(err)
		return
	}

	// Unmarshal RequestBody
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &currentProduct)
	if err != nil {
		c.ServerError(err)
		return
	}

	// UpdateProductById
	err = models.UpdateProductById(currentProduct)
	if err != nil {
		c.ServerError(err)
		return
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", nil)
}

// GetProducts ...
// @Title Get All Product List
// @Description Get Product list by some info
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	    query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @router /all [get]
// @Success 200 {object} models.Product
// @Failure 500
func (c *ProductController) GetProducts() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var pageSize int64 = 10
	var pageNumber int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
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
	l, err := models.GetAllProduct(query, fields, sortby, order, offset, pageSize)
	if err != nil {
		c.ServerError(err)
		return
	}

	var pageList []interface{}
	for _, v := range l {
		pageList = append(pageList, v)
	}

	/**
	 * 查询 Product Count
	 */
	cnt, err := models.GetProductCount()
	if err != nil {
		c.ServerError(err)
		return
	}

	if pages, err := helpers.NewPagination(pageList, int(cnt), int(pageSize), int(pageNumber)); err == nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", *pages)
	}
}
