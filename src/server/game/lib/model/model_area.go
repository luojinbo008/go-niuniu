package model

import (
	// "gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
	"fmt"
)


type AreaList struct{
	Areas []AreaDate
}

// 数据库的数据
type AreaDate struct {
	AreaId			int    "_id" 	// 场次id 自增型的
	AreaType		int				// 场次类型
	IsCivilian 		int				// 是否平民房间
	ServiceCost		int             // 茶水费
	BaseGold		int				// 最低 进场费
	LimitGold		int 			//
}

const (
	AREADB  	= "nn_area"
)

func (data *AreaDate) initAreaInfo(model *Model) (error) {

	areaId, err := model.mongoDBNextSeq("nn_area")

	if err != nil {
		log.Debug("get next area id error: %v", err)
		return fmt.Errorf("get next area id error: %v", err)
	}
	data.AreaId = areaId
	return nil
}

func (model *Model) GetAreaList() (areaList AreaList) {

	areaDate := AreaDate{}

	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	iter := db.DB(DB).C(AREADB).Find(nil).Iter()
	for iter.Next(&areaDate) {
		areaList.Areas = append(areaList.Areas, areaDate)
	}
	return areaList
}

func (model *Model) AddArea(areaData AreaDate) (areaInfo AreaDate, err error) {
	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	areaData.initAreaInfo(model)

	err = db.DB(DB).C(AREADB).Insert(areaData)
	return areaData, err
}