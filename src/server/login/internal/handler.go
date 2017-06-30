package internal
import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"server/game"
)

func init() {
	handler(&msg.UserLoginInfo{},handlLoginUser)
}

func handler(m interface{}, h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlLoginUser(args []interface{}) {

	// 收到登陆信息
	m := args[0].(*msg.UserLoginInfo)

	// 获取发送者
	a := args[1].(gate.Agent)


	// 交给 game
	game.ChanRPC.Go("LoginAgent", a, m)

	// 返回信息
	a.WriteMsg(
		&msg.CodeState{MSG_STATE : msg.MSG_Login_OK},
	)
}