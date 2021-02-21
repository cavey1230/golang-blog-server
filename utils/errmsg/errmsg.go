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
	ERROR_USERS_ROLE_ERROR     = 1009
	//CODE = 2000...文章模块的错误
	ERROR_CATEGORY_USED           = 2001
	ERROR_CATEGORY_PAGEINFO_ERROR = 2002
	//CODE = 3000...分类模块的错误
	ERROR_ARTICLE_TITLE_USED = 3001
	ERROR_ARTICLE_NOT_DEFINE = 3002
)

var codeMsg = map[int]string{
	//CODE = 1000...用户模块的错误
	SUCCSE:                     "OK",
	ERROR:                      "FAIL",
	ERROR_USERNAME_USED:        "用户名已存在",
	ERROR_PASSWORD_WRONG:       "密码错误",
	ERROR_USER_NOT_EXIST:       "用户不存在",
	ERROR_TOKEN_EXIST:          "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:        "TOKEN过期",
	ERROR_TOKEN_WRONG:          "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:     "TOKEN格式错误",
	ERROR_USERS_PAGEINFO_ERROR: "暂无用户数据",
	ERROR_USERS_ROLE_ERROR:     "角色状态异常",
	//CODE = 2000...文章模块的错误
	ERROR_CATEGORY_USED:           "文章分类已占用",
	ERROR_CATEGORY_PAGEINFO_ERROR: "暂无分类数据",
	//CODE = 3000...分类模块的错误
	ERROR_ARTICLE_TITLE_USED: "文章标题已占用",
	ERROR_ARTICLE_NOT_DEFINE: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
