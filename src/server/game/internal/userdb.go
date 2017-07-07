package internal

import (
	"server/msg"
	"gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
	"time"
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

// 数据库的数据
type ThirdInfoData struct {
	UserID		int			// 用户id 自增型的
	Openid		string		// 微信id
	Unionid		string		// 平台id
}

type UserData struct {
	UserID		int	"_id"	// 用户id 自增型的
	AccountID	string					// 用户线上看到的id
	NickName	string					// 用户的昵称
	Sex 		int						// 性别 0--女 1--男
	TotalCount	int						// 比赛总次数
	WinCount	int						// 胜利次数
	Money		int						// 账号金币
	HeadImgUrl	string					// 头像
	CreatedAt	int64					// 注册时间
	AccessToken	string 					// token
}

const (
	THIRDINFO 	= "thirdinfos"
	USERDB  	= "users"
	MD5KEY		= "UvogCy0Y18UODPB4"
)

func (data *UserData) initValue() error {
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
func wechatLogin(user  *msg.UserLoginByWechat) (userInfo UserData, err error) {


	var userThirdInfo 	ThirdInfoData
	db := mongoDB.Ref()
	defer mongoDB.UnRef(db)
	err = db.DB(DB).C(THIRDINFO).Find(bson.M{
		"openid"	: user.Code,
	}).One(&userThirdInfo)

	if err != nil {
		log.Debug(err.Error())
		if err.Error() == "not found" {
			err = userInfo.initValue()
			if err != nil {
				log.Error("wechat login err - %v", err)
				return userInfo, err
			}
			h := md5.New()

			h.Write([]byte(fmt.Sprintf("%s%s", userInfo.AccountID, MD5KEY)))
			userInfo.AccessToken = hex.EncodeToString(h.Sum(nil))

			err = db.DB(DB).C(USERDB).Insert(userInfo)
			if err != nil {
				log.Error("wechat login user add err - %v", err)
				return userInfo, err
			}
			err = db.DB(DB).C(THIRDINFO).Insert(&ThirdInfoData{
				userInfo.UserID,
				user.Code,
				"",
			})
			if err != nil {
				log.Error("wechat login thirdinfo add err - %v", err)
				return userInfo, err
			}
			return userInfo, nil
		}
	}

	err = db.DB(DB).C(USERDB).Find(bson.M{
		"_id"	: userThirdInfo.UserID,
	}).One(&userInfo)

	if err != nil {
		return userInfo, err
	}
	log.Debug("login info - %v is login in", userInfo.AccountID)
	return userInfo, err

}

