package global

//错误码对应的意思
var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	LOGIN_ERROR : "登录信息错误或已过期",
	ERROR_EXIST_TAG : "已存在该标签名称",
	ERROR_NOT_EXIST_TAG : "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}


const (
	WEB_ADMINS_LOGIN_EXPIRE = 86400 * 7
	WEB_ADMINS_MAX_UPLOAD_SIZE = 1024 * 1024 * 4
)