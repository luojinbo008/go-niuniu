package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	game_init "server/game/lib/init"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {

	m.Skeleton = skeleton
	game_init.InitAreaAndRoom()
	game_init.InitRoomDesk()
}

func (m *Module) OnDestroy() {

}
