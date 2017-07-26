package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	THIRDINFODB = "thirdinfos"
)

// 数据库的数据
type ThirdInfoData struct {
	UserID		int			// 用户id
	Openid		string		// 微信id
	Unionid		string		// 平台id
}

// 根据openId 获得三方账号信息
func (model *Model) GetThirdInfoUserInfoByThirdId(openid string) (userThirdInfo ThirdInfoData, err error) {
	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	err = db.DB(DB).C(THIRDINFODB).Find(bson.M{
		"openid"	: openid,
	}).One(&userThirdInfo)
	return userThirdInfo, err
}

// 新增第三方账号信息
func (model *Model) AddThirdInfoUserInfo(userThirdInfo ThirdInfoData) (userThirdData ThirdInfoData, err error) {
	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	err = db.DB(DB).C(THIRDINFODB).Insert(userThirdInfo)
	return userThirdInfo, err
}