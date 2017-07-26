package tool

import (
	"server/game/lib/model"
	"server/game/lib/cache"
)

var gameModel model.Model
var gameCache cache.Cache

func init()  {
	gameModel.InitModel()
	gameCache.InitModel()
}
