package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000+iota



	CodeServerBusy
)


var CodeMsgMap = map[ResCode]string{
	CodeSuccess: "success",


	CodeServerBusy: "服务繁忙",
}

func (c ResCode) Msg()string{
	msg,ok:=CodeMsgMap[c]
	if !ok{
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}

