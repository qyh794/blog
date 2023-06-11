package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvaildParam
	CodeUserExist
	CodeUserNotExist
	CodeInvaildPassword
	CodeSeverBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeDeletePost
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvaildParam:    "参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvaildPassword: "用户名或密码错误",
	CodeSeverBusy:       "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效token",
	CodeDeletePost:      "删除失败",
}

func (res ResCode) Msg() string {
	msg, ok := codeMsgMap[res]
	if !ok {
		msg = codeMsgMap[CodeSeverBusy]
	}
	return msg
}

