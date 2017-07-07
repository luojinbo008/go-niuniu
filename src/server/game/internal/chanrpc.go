package internal

import (
	"github.com/name5566/leaf/gate"
	"server/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginAgent", rpcLoginAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}


func rpcLoginAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.UserLoginByWechat)

	userInfo, err := wechatLogin(m)
	if err != nil {
		a.WriteMsg(
			&msg.CodeState{
				CODE 	: msg.MSG_DB_Error,
				MSG		: "DB ERROR",
			},
		)
		return
	}

	// 返回信息
	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_Login_OK,
			MSG		: "LOGIN SUCCESS",
			DATA	: userInfo,
		},
	)
	return
}