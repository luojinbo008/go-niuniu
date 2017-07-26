package model

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
	"time"
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

// 数据库的数据
type UserData struct {
	UserID      int    "_id" // 用户id 自增型的
	AccountID   string       // 用户线上看到的id
	NickName    string       // 用户的昵称
	Sex         int          // 性别 0--未知 1--男 2--女
	HeadImgUrl  string       // 头像
	Diamond     int          // 钻石
	AccessToken string       // token
	CreatedAt   int64        // 注册时间
}

const (
	USERDB  	= "users"
	MD5KEY		= "UvogCy0Y18UODPB4"
)

func (data *UserData) initUserInfo(model *Model) (error) {
	userID, err := model.mongoDBNextSeq("users")

	if err != nil {
		log.Debug("get next users id error: %v", err)
		return fmt.Errorf("get next users id error: %v", err)
	}

	data.UserID = userID
	data.AccountID = time.Now().Format("0102") + strconv.Itoa(data.UserID)
	data.CreatedAt = time.Now().Unix()
	return nil
}

func (model *Model) GetUserInfoByUserId(userId int) (userInfo UserData, err error) {
	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	err = db.DB(DB).C(USERDB).Find(bson.M{
		"_id"	: userId,
	}).One(&userInfo)

	return userInfo, err
}

func (model *Model) AddUserInfo(userData UserData) (userInfo UserData, err error) {

	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	userData.initUserInfo(model)

	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%s", userData.AccountID, MD5KEY)))
	userData.AccessToken = hex.EncodeToString(h.Sum(nil))

	err = db.DB(DB).C(USERDB).Insert(userData)
	return userData, err
}

func (model *Model) GetUserInfoByAccessToken(accountID string, accessToken string) (userInfo UserData, err error) {

	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	err = db.DB(DB).C(USERDB).Find(bson.M{
		"accountid"		: accountID,
		"accesstoken"	: accessToken,
	}).One(&userInfo)
	return userInfo, err
}