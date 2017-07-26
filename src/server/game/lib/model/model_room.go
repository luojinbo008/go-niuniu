package model

import (
	// "gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
	"fmt"
)


type RoomList struct{
	Rooms []RoomDate
}

// 数据库的数据
type RoomDate struct {
	RoomId			int	"_id" 	// 房间id 自增型的
	AreaId			int			// 场次id
	Name			string		// 房间名
	DeskNum 		int			// 开桌数
	Status			int     	// 茶水费
}

const (
	ROOMDB  	= "nn_room"
)

func (data *RoomDate) initRoomInfo(model *Model) (error) {

	roomId, err := model.mongoDBNextSeq("nn_room")

	if err != nil {
		log.Debug("get next rooms id error: %v", err)
		return fmt.Errorf("get next rooms id error: %v", err)
	}
	data.RoomId = roomId
	return nil
}

func (model *Model) GetRoomList() (roomList RoomList) {

	roomDate := RoomDate{}

	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	iter := db.DB(DB).C(ROOMDB).Find(nil).Iter()

	for iter.Next(&roomDate) {
		roomList.Rooms = append(roomList.Rooms, roomDate)
	}

	return roomList
 }

func (model *Model) AddRoom(roomData RoomDate) (roomInfo RoomDate, err error) {
	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	roomData.initRoomInfo(model)

	err = db.DB(DB).C(ROOMDB).Insert(roomData)
	return roomData, err
}