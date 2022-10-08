package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "invalid param",
	CodeUserExist:    "用户名已存在",
	CodeUserNotExist: "用户名不存在",
	CodeServerBusy:   "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
