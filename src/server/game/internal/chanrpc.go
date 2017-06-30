package internal

import (
	"github.com/name5566/leaf/gate"
	"fmt"
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
	fmt.Println("-rpclon-:",args)
	a := args[0].(gate.Agent)
	fmt.Println("get m--:",a)
	fmt.Println("len--:",len(args))
	m := args[1].(*msg.UserLoginInfo)
	err := login(m)
	if err != nil{
		a.WriteMsg(
			&msg.CodeState{MSG_STATE : msg.MSG_DB_Error},
		)
		return
	}
}