package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/models"
	common2 "cleverbamboo.com/bee-shop-b2c/utils/common"
)

type AccountController struct {
	BaseController
}

// URLMapping ...
func (c *AccountController) URLMapping() {
	c.Mapping("Register", c.Register)
}

/**
 * 用户注册
 * username是唯一判断
 */
// @router /register [post]
func (c *AccountController) Register() {
	name := c.GetString("name")
	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" {
		c.JsonResult(common.ErrError, "账号不能为空", nil)
		return
	}

	if password == "" {
		c.JsonResult(common.ErrError, "密码不能为空", nil)
		return
	}

	if name == "" {
		c.JsonResult(common.ErrError, "昵称不能为空", nil)
		return
	}

	member, _ := models.GetMemberByUsername(username)

	if member == nil {
		member := models.NewMember()
		member.Username = username
		member.Name = name
		member.Password = common2.Md5Crypt(password)
		id, err := models.AddMember(member)
		if err != nil {
			c.ServerError(err)
			return
		} else {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = id
		}
		c.JsonResult(common.ErrOK, "ok", c.Data["json"])
	} else {
		c.JsonResult(common.ErrError, "该用户已经注册", nil)
	}
}

func (c *AccountController) Login(username string, password string) {

}
