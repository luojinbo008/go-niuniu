package model

import (
	"server/msg"
	"gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
	"time"
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"go.lib/wechat/oauth2"
)

// 数据库的数据
type ThirdInfoData struct {
	UserID		int			// 用户id 自增型的
	Openid		string		// 微信id
	Unionid		string		// 平台id
}

type UserData struct {
	UserID			int	"_id"		// 用户id 自增型的
	AccountID		string			// 用户线上看到的id
	NickName		string			// 用户的昵称
	Sex 			int				// 性别 0--未知 1--男 2--女
	HeadImgUrl		string			// 头像
	CreatedAt		int64			// 注册时间
	AccessToken		string 			// token
	UserGameInfo	UserGameData	// 游戏数据
}

const (
	THIRDINFODB = "thirdinfos"
	USERDB  	= "users"
	MD5KEY		= "UvogCy0Y18UODPB4"
)

func (data *UserData) initUserInfo() error {
	userID, err := mongoDBNextSeq("users")

	if err != nil {
		log.Debug("get next users id error: %v", err)
		return fmt.Errorf("get next users id error: %v", err)
	}

	data.UserID = userID
	data.AccountID = time.Now().Format("0102") + strconv.Itoa(data.UserID)
	data.CreatedAt = time.Now().Unix()
	return nil
}

// 微信端登陆
func WechatLogin(wechatLogin *msg.UserLoginByWechat) (userInfo UserData, err error) {
	var userThirdInfo 	ThirdInfoData
	db := mongoDB.Ref()
	defer mongoDB.UnRef(db)

	wechatUserInfo, err := oauth2.GetAutoUserInfo(wechatLogin.Code)

	if err != nil {
		log.Error("login get thirdinfo - %s", err)
		return userInfo, err
	}

	if wechatUserInfo.Errcode != 0 {
		log.Error("login get thirdinfo - %s, %s", wechatUserInfo.Errcode, wechatUserInfo.Errmsg)

		return userInfo, fmt.Errorf(wechatUserInfo.Errmsg)
	}

	err = db.DB(DB).C(THIRDINFODB).Find(bson.M{
		"openid"	: wechatUserInfo.Openid,
	}).One(&userThirdInfo)


	if err != nil {
		if err.Error() == "not found" {
			err = userInfo.initUserInfo()
			if err != nil {
				log.Error("go.lib login err - %v", err)
				return userInfo, err
			}
			h := md5.New()

			h.Write([]byte(fmt.Sprintf("%s%s", userInfo.AccountID, MD5KEY)))

			userInfo.AccessToken = hex.EncodeToString(h.Sum(nil))
			userInfo.NickName = wechatUserInfo.Nickname
			userInfo.HeadImgUrl = wechatUserInfo.Headimgurl
			userInfo.Sex = wechatUserInfo.Sex

			err = db.DB(DB).C(USERDB).Insert(userInfo)
			if err != nil {
				log.Error("login user add err - %v", err)
				return userInfo, err
			}
			err = db.DB(DB).C(THIRDINFODB).Insert(&ThirdInfoData{
				userInfo.UserID,
				wechatUserInfo.Openid,
				"",
			})
			if err != nil {
				log.Error("login thirdinfo add err - %v", err)
				return userInfo, err
			}
		}
	} else {
		err = db.DB(DB).C(USERDB).Find(bson.M{
			"_id"	: userThirdInfo.UserID,
		}).One(&userInfo)

		if err != nil {
			return userInfo, err
		}
	}

	userGameData, err := getUserGameData(userInfo.UserID)
	if err != nil {
		return userInfo, err
	}
	userInfo.UserGameInfo = userGameData
	log.Debug("login info - %v is login in", userInfo.AccountID)
	return userInfo, err
}


// 重连登陆
func ReLogin(userReLogin *msg.UserReLogin) (userInfo UserData, err error) {

	db := mongoDB.Ref()
	defer mongoDB.UnRef(db)

	err = db.DB(DB).C(USERDB).Find(bson.M{
		"accountid"		: userReLogin.AccountID,
		"accesstoken"	: userReLogin.AccessToken,
	}).One(&userInfo)

	if err != nil {
		return userInfo, err
	}

	userGameData, err := getUserGameData(userInfo.UserID)
	if err != nil {
		return userInfo, err
	}

	userInfo.UserGameInfo = userGameData
	log.Debug("login info - %v is login in", userInfo.AccountID)
	return userInfo, err
}
