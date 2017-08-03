package cache

import (
	"gopkg.in/mgo.v2/bson"
)

type LineAreaConfig struct {
	AreaId 			int "_id"
	AreaType		int
	IsCivilian		int
	AreaTypeName	string
	ServiceCost		int
	BaseGold		int
	LimitGold		int
	Num				int
	MaxNum 			int
}

const (
	LINEAREACONFIGDB 	= "line_nn_area_config"
)

func (cache *Cache) AddLineArea(lineAreaConfig LineAreaConfig) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEAREACONFIGDB).Insert(lineAreaConfig)
	return err
}

func (cache *Cache) GetAreaConfig(areaId int) (lineAreaConfig LineAreaConfig, err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEAREACONFIGDB).Find(bson.M{
		"_id"	: areaId,
	}).One(&lineAreaConfig)

	return lineAreaConfig, err
}

func (cache *Cache) IncrAreaUserNum(areaId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEAREACONFIGDB).Update(
		bson.M{"_id": areaId},
		bson.M{"$inc": bson.M{ "Num":  + 1}},
	)
	return err
}

func (cache *Cache) LeaveLineArea(areaId int) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEAREACONFIGDB).Update(
		bson.M{"_id": areaId},
		bson.M{"$inc": bson.M{"Num":  - 1}},
	)
	return err
}