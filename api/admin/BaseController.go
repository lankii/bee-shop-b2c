package admin

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
	Admin *models.Admin //用户
}

//每个子类Controller公用方法调用前，都执行一下Prepare方法
func (c *BaseController) Prepare() {
	c.Admin = &models.Admin{} //初始化
	//从session中获取用户信息
	if admin, ok := c.GetSession(common.SessionName).(models.Admin); ok && admin.Id > 0 {
		logs.Info("admin", admin)
		c.Admin = &admin
	}
}

// Ajax接口返回Json
func (c *BaseController) JsonResult(status int, errCode int, errMsg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["err_code"] = errCode
	jsonData["message"] = errMsg

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = jsonData
	c.ServeJSON()
}

// 设置登录管理员信息
func (c *BaseController) SetAdmin(admin models.Admin) {
	if admin.Id <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(common.SessionName, admin)
		c.SetSession("uid", admin.Id)
	}
}

func (c *BaseController) ServerError(err error) {
	c.Ctx.Output.SetStatus(common.GetHttpStatus("internalServerError"))
	logs.Error(err)
}
