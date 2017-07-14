package internal
import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"server/game"
)

func init() {
	handler(&msg.UserLoginByWechat{},handlWechatLoginUser)
	handler(&msg.UserReLogin{},handlReLoginUser)
}

func handler(m interface{}, h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlWechatLoginUser(args []interface{}) {

	// 收到登陆信息
	m := args[0].(*msg.UserLoginByWechat)

	// 获取发送者
	a := args[1].(gate.Agent)

	// 交给 game
	game.ChanRPC.Go("LoginWechatAgent", a, m)
}

func handlReLoginUser(args []interface{})  {

	// 收到登陆信息
	m := args[0].(*msg.UserReLogin)

	// 获取发送者
	a := args[1].(gate.Agent)

	// 交给 game
	game.ChanRPC.Go("LoginReAgent", a, m)

}