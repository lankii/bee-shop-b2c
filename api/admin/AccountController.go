package admin

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/models"
	"regexp"
	"strings"
	"time"
)

type AccountController struct {
	BaseController
}

// URLMapping ...
func (c *AccountController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
}

// Login ...
// @Title Login
// @Description Admin Login
// @Param username body string true "Login username"
// @Param password body string true "Login password"
// @Success 200 {object} models.Admin
// @Failure 500
// @router /login [post]
func (c *AccountController) Login() {
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
// @Description Admin Login
// @Param username  body string  true  "Register username"
// @Param password1 body string  true  "Register password1"
// @Param password2 body string  true  "Register password2"
// @Param nickname  body string  true  "Register nickname"
// @Param email     body string  true  "Register email"
// @Success 200 nil
// @Failure 500
// @router /register [post]
func (c *AccountController) Register() {
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
