package cache

import (
	"gopkg.in/mgo.v2/bson"
)

type LineUserFd struct {
	Ip 		string 	// 服务器IP
	Fd 		string 	// 内存地址
	UserId 	int		// 用户id
}

const (
	LINEUSERFDDB				= "line_nn_user_fd"
)

func (cache *Cache) AddLineUserFd(lineUserFd LineUserFd) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERFDDB).Insert(lineUserFd)
	return err
}

func (cache *Cache)RemoveLineUserFd (lineUserFd LineUserFd) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERFDDB).Remove(bson.M{
		"ip"	: lineUserFd.Ip,
		"fd"	: lineUserFd.Fd,
	})
	return err
}

func (cache *Cache)RemoveLineUserFdByUserId (userId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERFDDB).Remove(bson.M{
		"userid"	: userId,
	})
	return err
}