package common

const (
	ok                  = 200
	created             = 201
	noContent           = 204
	forbidden           = 403
	internalServerError = 500
)

/**
 * 获取标准的 HTTP 状态码, -1 表示未找到
 */
func GetHttpStatus(s string) int {
	switch s {
	case "ok":
		return ok
	case "created":
		return created
	case "noContent":
		return noContent
	case "forbidden":
		return forbidden
	case "internalServerError":
		return internalServerError
	default:
		return -1
	}
}
