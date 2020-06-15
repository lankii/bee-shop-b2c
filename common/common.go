package common

import "github.com/astaxie/beego"

//session
const SessionName = "__beeShopB2C_session__"

//用户权限
const (
	//超级管理员
	RAdmin = 1
	//普通管理员
	RManager = 2
	//普通会员
	RMember = 3
	//普通用户
	RUser = 4
)

func Role(role int) string {
	if role == RAdmin {
		return "R_ADMIN"
	} else if role == RManager {
		return "R_MANAGER"
	} else if role == RMember {
		return "R_MEMBER"
	} else if role == RUser {
		return "R_USER"
	}
	return ""
}

// app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "beeShopB2C")
}

const (
	//0,请求成功
	ErrOK = 0
	//1,请求失败
	ErrError = 1
)
