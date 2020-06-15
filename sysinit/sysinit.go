package sysinit

import "github.com/astaxie/beego/logs"

func sysInit() {
	_ = logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
}
