package internal

import (
	"github.com/name5566/leaf/gate"
	"server/msg"
	"fmt"
	"server/game/lib/tool"
)


var Agents map[string]gate.Agent

func init() {
	Agents = make(map[string]gate.Agent)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginWechatAgent", rpcLoginWechatAgent)
	skeleton.RegisterChanRPC("LoginReAgent", rpcLoginReAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	Agents[fmt.Sprint("%s", &a)] = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	Agents[fmt.Sprint("%s", &a)] = nil
	userData := a.UserData()
	var data tool.UserInfo
	if userData != nil {
		data = userData.(tool.UserInfo)
		tool.LineUserCut(fmt.Sprint("%s", &a), data.UserId)
	}

}

func rpcLoginWechatAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.UserLoginByWechat)
	userInfo, err := tool.WechatLogin(m)

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
	tool.LineUserModify(a)


	// 登陆返回
	tool.PushUserLoginInfo(a)
	return
}

func rpcLoginReAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.UserReLogin)

	userInfo, err := tool.ReLogin(m)

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
	tool.LineUserModify(a)

	// 登陆返回
	tool.PushUserLoginInfo(a)
}
