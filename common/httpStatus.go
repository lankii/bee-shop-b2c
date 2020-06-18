package common

const (
	ok                  = 200
	created             = 201
	internalServerError = 500
)

/**
 * 获取标准的 HTTP 状态码, -1 表示未找到
 */
func GetHttpStatus(s string) int {
	if s == "ok" {
		return ok
	} else if s == "created" {
		return created
	} else if s == "internalServerError" {
		return internalServerError
	}
	return -1
}
