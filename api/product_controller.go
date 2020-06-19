package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

type ProductController struct {
	BaseController
}

// URLMapping ...
func (c *ProductController) URLMapping() {
	c.Mapping("GetProducts", c.GetProducts)
}

// GetProducts ...
// @Title Get Product List
// @Description Get Product list by some info
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @router /all [get]
func (c *ProductController) GetProducts() {
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var pageSize int = 10
	var pageNumber int

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

	// pageNumber: integer
	if v := c.GetString("pageNumber"); v != "" {
		pageNumber, _ = strconv.Atoi(v)
	}

	// pageSize: integer
	if v := c.GetString("pageSize"); v != "" {
		pageSize, _ = strconv.Atoi(v)
	}

	logs.Info("sortby", sortby, pageNumber, pageSize, order)
}
