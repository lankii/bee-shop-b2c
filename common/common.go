package common

import "github.com/astaxie/beego"

// session
const SessionName = "__beeShopB2C_session__"

// app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "beeShopB2C")
}

// 正则表达式
const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

/**
 * err code
 */
const (
	// 失败
	ErrError = 0
	// 成功
	ErrOK = 1
)

/**
 * err message
 */
const (
	Success = "success"
	Fail    = "fail"
)
