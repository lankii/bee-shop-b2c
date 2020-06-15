package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io"
)

type BaseController struct {
	beego.Controller
	Member          *models.Member    //用户
	Option          map[string]string //全局设置
	EnableAnonymous bool              //开启匿名访问
}

/**
 * 每个子类Controller公用方法调用前，都执行一下Prepare方法
 */
func (c *BaseController) Prepare() {
	c.Member = models.NewMember()
	c.EnableAnonymous = false //初始化
	//从session中获取用户信息
	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.Id > 0 {
		c.Member = &member
	} else {

	}
}

/**
 * Ajax接口返回Json
 */
func (c *BaseController) JsonResult(errCode int, errMessage string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["errCode"] = errCode
	jsonData["message"] = errMessage
	jsonData["data"] = nil

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		logs.Info(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	c.StopRun()
}

func (c *BaseController) ServerError(err error) {
	c.Ctx.Output.SetStatus(500)
	logs.Error(err)
}
