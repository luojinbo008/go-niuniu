package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	// "server/game/lib/tool"
)

func init() {

	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)
	// handler(&msg.EnterGameRoom{}, handleEnterGameRoom)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {

	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v",m.Name)
	a.WriteMsg(&msg.Hello{ Name : "Client", })
}

/*
func handleEnterGameRoom(args []interface{})  {

	m := args[0].(*msg.EnterGameRoom)
	a := args[1].(gate.Agent)

	game_tool.EnterGameRoom(a, m)

}*/
