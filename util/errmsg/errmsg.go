package errmsg

const (
	// 成功
	Success = 200
	// 失败
	Error = 500

	// 用户相关 code = 10001
	ERROR_USER_PASSWORD_ERROR = 10001 // 密码错误
	ERROR_USER_NOT_EXIST      = 10002 // 用户不存在
	ERROR_USER_ADD            = 10003 // 用户添加失败
	ERROR_USER_ADD_EXIST      = 10004 // 用户已存在
)

var CodeMsg = map[int]string{
	Success:                   "success",
	Error:                     "error",
	ERROR_USER_PASSWORD_ERROR: "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_USER_ADD:            "用户添加失败",
	ERROR_USER_ADD_EXIST:      "用户已存在",
}

func GetMsg(code int) string {
	return CodeMsg[code]
}
