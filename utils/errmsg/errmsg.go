package errmsg

const (
	SUCCSE = 200
	ERROR  = 500
	//CODE = 1000...用户模块的错误
	ERROR_USERNAME_USED        = 1001
	ERROR_PASSWORD_WRONG       = 1002
	ERROR_USER_NOT_EXIST       = 1003
	ERROR_TOKEN_EXIST          = 1004
	ERROR_TOKEN_RUNTIME        = 1005
	ERROR_TOKEN_WRONG          = 1006
	ERROR_TOKEN_TYPE_WRONG     = 1007
	ERROR_USERS_PAGEINFO_ERROR = 1008
	//CODE = 2000...文章模块的错误

	//CODE = 3000...分类模块的错误
)

var codeMsg = map[int]string{
	SUCCSE:                     "OK",
	ERROR:                      "FAIL",
	ERROR_USERNAME_USED:        "用户名已存在",
	ERROR_PASSWORD_WRONG:       "密码错误",
	ERROR_USER_NOT_EXIST:       "用户不存在",
	ERROR_TOKEN_EXIST:          "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:        "TOKEN过期",
	ERROR_TOKEN_WRONG:          "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:     "TOKEN格式错误",
	ERROR_USERS_PAGEINFO_ERROR: "暂无数据",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
