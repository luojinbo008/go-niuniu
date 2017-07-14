package internal

import (
	"github.com/name5566/leaf/gate"
	"server/msg"
	"server/game/lib/model"
	"fmt"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginWechatAgent", rpcLoginWechatAgent)
	skeleton.RegisterChanRPC("LoginReAgent", rpcLoginReAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcLoginWechatAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.UserLoginByWechat)

	userInfo, err := model.WechatLogin(m)

	if err != nil {
		a.WriteMsg(
			&msg.CodeState{
				CODE 	: msg.MSG_CODE_ERROR,
				CMD		: 9000,
				MSG		: fmt.Sprintf("LOGIN ERROR:%s", err.Error()),
			},
		)
		return
	}
	a.SetUserData(userInfo)
	// 返回信息
	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_CODE_SUCCESS,
			CMD		: msg.CMD_MY_USER_INFO,
			MSG		: "LOGIN SUCCESS",
			DATA	: a.UserData(),
		},
	)
	return
}

func rpcLoginReAgent(args []interface{})  {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.UserReLogin)

	userInfo, err := model.ReLogin(m)

	if err != nil {
		a.WriteMsg(
			&msg.CodeState{
				CODE 	: msg.MSG_CODE_ERROR,
				CMD		: 9000,
				MSG		: fmt.Sprintf("RELOGIN ERROR:%s", err.Error()),
			},
		)
		return
	}

	a.SetUserData(userInfo)

	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_CODE_SUCCESS,
			CMD		: msg.CMD_MY_USER_INFO,
			MSG		: "RELOGIN SUCCESS",
			DATA	: a.UserData(),
		},
	)
}