package internal
import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"server/game"
)

func init() {
	handler(&msg.UserLoginByWechat{},handlLoginUser)
}

func handler(m interface{}, h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlLoginUser(args []interface{}) {

	// 收到登陆信息
	m := args[0].(*msg.UserLoginByWechat)

	// 获取发送者
	a := args[1].(gate.Agent)

	// 交给 game
	game.ChanRPC.Go("LoginAgent", a, m)

}