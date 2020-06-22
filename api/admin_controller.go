package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AdminController struct {
	BaseController
}

// URLMapping ...
func (c *AdminController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
	c.Mapping("GetAllAdmins", c.GetAdmins)
	c.Mapping("UpdateAdmin", c.UpdateAdmin)
}

// Login ...
// @Title Login
// @Description Admin Login
// @Param username body string true "Login username"
// @Param password body string true "Login password"
// @Success 200 {object} models.Admin
// @Failure 500
// @router /login [post]
func (c *AdminController) Login() {
	var username string
	var password string

	if v := c.GetString("username"); v != "" {
		username = v
	}
	if v := c.GetString("password"); v != "" {
		password = v
	}

	currentAdmin, err := models.GetAdminByUsernameAndIsEnabled(username, 1)
	if err != nil {
		c.ServerError(err)
		return
	}
	if currentAdmin == nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, "用户不存在", nil)
		return
	}

	ok, err := helpers.PasswordVerify(currentAdmin.Password, password)
	if err != nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, "密码不正确", nil)
		return
	}
	// TODO 需要过滤信息，只返回必须的信息
	if ok {
		if adminRole, err := models.GetAdminRoleIdByAdmins(currentAdmin.Id); err == nil {
			currentAdmin.RoleId = adminRole.Roles
			currentAdmin.LoginDate = time.Now()
			c.SetAdmin(*currentAdmin)
			c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", *currentAdmin)
		}
	}
}

// Register ...
// @Title Register
// @Description Admin Register
// @Param username  body string  true  "Register username"
// @Param password1 body string  true  "Register password1"
// @Param password2 body string  true  "Register password2"
// @Param nickname  body string  true  "Register nickname"
// @Param email     body string  true  "Register email"
// @Success 200 nil
// @Failure 500
// @router /register [post]
func (c *AdminController) Register() {
	var username string
	var password1 string
	var password2 string
	var nickname string
	var email string

	if v := c.GetString("username"); v != "" {
		username = v
	}
	if v := c.GetString("password1"); v != "" {
		password1 = v
	}
	if v := c.GetString("password2"); v != "" {
		password2 = v
	}
	if v := c.GetString("nickname"); v != "" {
		nickname = v
	}
	if v := c.GetString("email"); v != "" {
		email = v
	}

	if password1 != password2 {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "登录密码与确认密码不一致")
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

	if l := strings.Count(nickname, "") - 1; l < 2 || l > 20 {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "用户昵称限制在2-20个字符")
		return
	}

	currentAdmin, _ := models.GetAdminByUsername(username)
	if currentAdmin != nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, "该用户已注册", nil)
		return
	}
	currentAdmin = &models.Admin{}
	currentAdmin.Username = username
	currentAdmin.Name = nickname
	currentAdmin.Email = email
	currentAdmin.LoginIp = c.Ctx.Request.Host
	currentAdmin.IsEnabled = 1

	hash, err := helpers.PasswordHash(password1)
	if err != nil {
		c.ServerError(err)
	}
	currentAdmin.Password = hash

	id, err := models.AddAdmin(currentAdmin)

	if err != nil {
		c.ServerError(err)
		return
	}

	if err = models.AddAdminRole(id, common.RUser); err == nil {
		c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, "success", nil)
	} else {
		c.ServerError(err)
	}
}

// GetAdmins ...
// @Title Get Admin List
// @Description Get Admin list by some info
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @router /all [get]
func (c *AdminController) GetAdmins() {
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
	l, err := models.GetAllAdmin(query, fields, sortby, order, offset, pageSize)
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

	if pages, err := helpers.NewPagination(pageList, int(cnt), int(pageSize), int(pageNumber), 5); err == nil {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", *pages)
	}
}

// UpdateAdmin ...
// @Title Update Admin
// @Description update Admin by some info
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
func (c *AdminController) UpdateAdmin() {
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
