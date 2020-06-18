package common

import "github.com/astaxie/beego"

//session
const SessionName = "__beeShopB2C_session__"

//app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "beeShopB2C")
}

//正则表达式
const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

const (
	//0,请求成功
	ErrOK = 0
	//1,请求失败
	ErrError = 1
)
