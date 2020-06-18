package admin

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/models"
	"errors"
	"strconv"
	"strings"
)

type IndexController struct {
	BaseController
}

// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("GetAllAdmins", c.GetAllAdmins)
}

// @Title Get All Admins
// @Description get All Admins
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @router /getAll [get]
func (c *IndexController) GetAllAdmins() {
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
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
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

	// start position of result set
	offset := (pageNumber - 1) * pageSize
	l, err := models.GetAllAdmin(query, nil, sortby, order, int64(offset), int64(pageSize))
	if err != nil {
		c.ServerError(err)
		return
	}

	var pageList []interface{}
	for _, v := range l {
		pageList = append(pageList, v)
	}

	/**
	 * 查询 Admin Count
	 */
	cnt, err := models.GetAdminCount()
	if err != nil {
		c.ServerError(err)
		return
	}

	if pages, err := helpers.NewPagination(pageList, int(cnt), pageSize, pageNumber, 5); err == nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", *pages)
	}
}
