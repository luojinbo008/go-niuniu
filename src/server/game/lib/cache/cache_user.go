package cache

import (
	"gopkg.in/mgo.v2/bson"
)

type LineUserInfo struct {

	UserId			int				// 用户ID

	IsPlaying		int				// 是否在游戏
	IsWaiting		int				// 是否等待开局了

	IsCutLine		int				// 是否断线
	IsOnLine		int 			// 是否在线
	InArea			int 			// 在哪个场次
	InRoom			int 			// 在哪个房间
	InDesk			int				// 在哪个桌子
	InSeat			int				// 在哪个座位
	InDeskStart		int				// 开始进入桌子时间
	IsOutRoom		int				// 在线退出桌子

}

const (
	LINEUSERDB				= "line_nn_user"
)

func (cache *Cache) ModifyLineUser(userInfo LineUserInfo) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	_, err = db.DB(DB).C(LINEUSERDB).Upsert(
		bson.M{
			"userid" : userInfo.UserId,
		},
		bson.M{"$set": bson.M{
				"isplaying":	userInfo.IsPlaying,
				"iswaiting":	userInfo.IsWaiting,
				"iscutline":	userInfo.IsCutLine,
				"isonline":		userInfo.IsOnLine,
				"inarea":		userInfo.InArea,
				"inroom":		userInfo.InRoom,
				"indesk":		userInfo.InDesk,
				"insert":		userInfo.InSeat,
				"indeskstart":	userInfo.InDeskStart,
				"isoutroom":	userInfo.IsOutRoom,
			},
		},
	)
	return err
}

func (cache *Cache) GetLineUser(userId int) (lineUserInfo LineUserInfo, err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERDB).Find(bson.M{
		"userid" : userId,
	}).One(&lineUserInfo)
	return lineUserInfo, err
}
