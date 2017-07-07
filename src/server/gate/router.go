package gate

import (
	"server/login"
	"server/game"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{},game.ChanRPC)

	msg.Processor.SetRouter(&msg.UserLoginByWechat{},login.ChanRPC)
}
