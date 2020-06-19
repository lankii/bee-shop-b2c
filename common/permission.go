package common

/**
 * 用户权限
 */
const (
	// 超级管理员
	RAdmin = 1
	// 普通管理员
	RManager = 2
	// 普通会员
	RMember = 3
	// 普通用户
	RUser = 4
)

func Role(role string) int64 {
	if role == "R_ADMIN" {
		return RAdmin
	} else if role == "R_MANAGER" {
		return RManager
	} else if role == "R_MEMBER" {
		return RMember
	} else if role == "R_USER" {
		return RUser
	}
	return -1
}
