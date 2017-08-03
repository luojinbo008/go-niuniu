package cache

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	LINEDESKDB				= "line_nn_desk"
)

type LineDesk struct {
	DeskId				int "_id"
	AreaId				int
	RoomId 				int
	Waiting				int

	SeatNum				int
	IsSeatNum			int
	IsPlaying			int
	InCompute			int
	CustomerNum			int

	SeatCardsAll		int
	StartTime			int
	EndTime				int
	Status				int
	StatusEndTime		int
	CurrentGameRound	int
	Banker				int	// 庄家 所在座位
	IsLock				int	// 锁定状态不允许用户进入
}

func (cache *Cache) AddLineDesk(lineDesk LineDesk) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEDESKDB).Insert(lineDesk)
	return err
}

func (cache *Cache) LeaveLineDesk(areaId int, roomId int, deskId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEDESKDB).Update(
		bson.M{
			"areaid": areaId,
			"roomid": roomId,
			"deskid": deskId,
		},
		bson.M{"$inc": bson.M{
				"isseatnum" 	: -1,
				"customernum"	: -1,
			},
		},
	)
	return err
}

func (cache *Cache) GetLineDesk(areaId int, roomId int, deskId int) (err error, lineDesk LineDesk) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEDESKDB).Find(
		bson.M{
			"areaid": areaId,
			"roomid": roomId,
			"deskid": deskId,
		},
	).One(&lineDesk)
	return err, lineDesk
}