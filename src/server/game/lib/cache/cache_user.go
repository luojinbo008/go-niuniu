package cache

import (
	"gopkg.in/mgo.v2/bson"
)

type LineUserInfo struct {
	Ip				string			// 服务器IP
	Fd				string			// 内存地址
	UserId			int				// 用户ID
	IsPlaying		int				// 是否在游戏
	IsOutRoom		int				// 是否在房间
}

const (
	LINEUSERDB				= "line_nn_user"
)

func (cache *Cache) ModifyLineUser(userInfo LineUserInfo) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERDB).Update(
		bson.M{
			"ip" : userInfo.Ip,
			"fd" : userInfo.Fd,
		},
		bson.M{"$unset": bson.M{
				"userid":		userInfo.UserId,
				"isplaying":	userInfo.IsPlaying,
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