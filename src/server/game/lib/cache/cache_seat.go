package cache

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	LINESEATDB				= "line_nn_seat"
)

type LineSeatList struct{
	LineSeats []LineSeat
}

type LineSeat struct {
	AreaId			int
	DeskId			int
	RoomId 			int
	SeatNo			int		// 座位号
	SeatType		int		// 座位上类型 1普通用户 9机器人
	SeatUserId		int		// 座位上用户uid 如果是机器人就是aid

	SeatFinish		int
	SeatCards		int
	SeatCardType	int
	SeatBigCard		int
	SeatBet			int
	SeatAutoFinish	int
}

func (cache *Cache) AddLineSeat(lineSeat LineSeat) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINESEATDB).Insert(lineSeat)
	return err
}

func (cache *Cache) GetLineDeskSeatList(areaId int, roomId int, deskId int) (lineSeatList LineSeatList) {
	lineSeat := LineSeat{}

	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	iter := db.DB(DB).C(LINESEATDB).Find(
		bson.M{
			"areaid": areaId,
			"roomid": roomId,
			"deskid": deskId,
		},
	).Iter()
	for iter.Next(&lineSeat) {
		lineSeatList.LineSeats = append(lineSeatList.LineSeats, lineSeat)
	}
	return lineSeatList
}

func (cache *Cache) LeaveLineSeat(areaId int, roomId int, deskId int, seatNo int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINESEATDB).Update(
		bson.M{
			"areaid": areaId,
			"roomid": roomId,
			"deskid": deskId,
			"seatno": seatNo,
		},
		bson.M{"$set": bson.M{
				"seattype" 		: 0,
				"seatuserid"	: 0,
			},
		},
	)
	return err
}