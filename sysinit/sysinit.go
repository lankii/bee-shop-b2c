package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func sysInit() {
	/**
     * 开发环境才打印SQL
	 */
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	_ = logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
}
