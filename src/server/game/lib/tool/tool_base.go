package tool

import (
	"server/game/lib/model"
	"server/game/lib/cache"
	"github.com/name5566/leaf/gate"
)

var gameModel model.Model
var gameCache cache.Cache


var Agents map[string]gate.Agent

func init()  {
	Agents = make(map[string]gate.Agent)
	gameModel.InitModel()
	gameCache.InitCache()
}
