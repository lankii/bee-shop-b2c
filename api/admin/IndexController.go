package admin

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/models"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type IndexController struct {
	BaseController
}

// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("GetAllAdmins", c.GetAllAdmins)
	c.Mapping("UpdateAdmin", c.UpdateAdmin)
}

// GetAllAdmins ...
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

// UpdateAdmin ...
// @Title Update Admin
// @Description update Admin
// @Param id         body string  true  "Admin Id"
// @Param password1  body string  true  "Should update Admin password1"
// @Param password2  body string  true  "Should update Admin password2"
// @Param email      body string  true  "Admin email"
// @Param role_name  body string  true  "Admin role name"
// @Param is_enabled body int     true  "If enabled"
// @Param department body string  false "Admin Department"
// @Param nickname   boyd string  false "Admin Nickname"
// @Success 200 nil
// @Failure 500
// @router /update [post]
func (c *IndexController) UpdateAdmin() {
	var id int64
	var password1 string
	var password2 string
	var email string
	var roleName string
	var isEnabled int

	// id
	if v, err := c.GetInt64("id"); v > -1 && err == nil {
		id = v
	}

	currentAdmin, err := models.GetAdminById(id)
	if err != nil {
		c.ServerError(err)
		return
	}

	// password1
	if v := c.GetString("password1"); v != "" {
		password1 = v
	}
	// password2
	if v := c.GetString("password2"); v != "" {
		password2 = v
	}
	// email
	if v := c.GetString("email"); v != "" {
		email = v
	}
	// roleName
	if v := c.GetString("role_name"); v != "" {
		roleName = v
	}
	// isEnabled
	if v, err := c.GetInt("is_enabled"); err == nil {
		isEnabled = v
	}
	// department
	if v := c.GetString("department"); v != "" {
		currentAdmin.Department = v
	}
	// nickname
	if v := c.GetString("nickname"); v != "" {
		currentAdmin.Name = v
	}

	if password1 != password2 {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "两次密码不一致")
		return
	}

	if l := strings.Count(password1, ""); password1 == "" || l > 20 || l < 6 {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "密码必须在6-20个字符之间")
		return
	}

	if ok, err := regexp.MatchString(common.RegexpEmail, email); !ok || err != nil || email == "" {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "邮箱格式错误")
		return
	}

	hash, err := helpers.PasswordHash(password1)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 赋值并更新 Admin
	 */
	currentAdmin.Password = hash
	currentAdmin.Email = email
	currentAdmin.RoleId = common.Role(roleName)
	currentAdmin.IsEnabled = int8(isEnabled)

	err = models.UpdateAdminById(currentAdmin)
	if err != nil {
		c.ServerError(err)
		return
	}
	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", nil)
}
