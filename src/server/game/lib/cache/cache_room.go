package cache

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	LINEROOMCONFIGDB  		= "line_nn_room_config"
)

type LineAreaRoomList struct {
	LineAreaRooms []LineRoomConfig
}

type LineRoomConfig struct {
	RoomId			int "_id"
	AreaId			int
	AreaType		int
	IsCivilian		int
	AreaTypeName	string
	ServiceCost		int
	BaseGold		int
	LimitGold		int
	DeskNum			int
	MaxUser			int
	UserNum			int
	Status			int
}


func (cache *Cache) AddLineRoom(lineRoomConfig LineRoomConfig) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEROOMCONFIGDB).Insert(lineRoomConfig)
	return err
}

func (cache *Cache) GetLineRoomList() (lineAreaRoomList LineAreaRoomList) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	lineRoomConfig := LineRoomConfig{}

	iter := db.DB(DB).C(LINEROOMCONFIGDB).Find(nil).Iter()

	for iter.Next(&lineRoomConfig) {
		lineAreaRoomList.LineAreaRooms = append(lineAreaRoomList.LineAreaRooms, lineRoomConfig)
	}
	return lineAreaRoomList
}

func (cache *Cache) IncrRoomUserNum(roomId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEROOMCONFIGDB).Update(
		bson.M{"_id": roomId},
		bson.M{"$inc": bson.M{ "Num":  + 1}},
	)
	return err
}

func (cache *Cache) LeaveLineRoom(areaId int, roomId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEROOMCONFIGDB).Update(
		bson.M{
			"_id"		: roomId,
			"areaid"	: areaId,
		},
		bson.M{"$inc": bson.M{"Num":  - 1}},
	)
	return err
}