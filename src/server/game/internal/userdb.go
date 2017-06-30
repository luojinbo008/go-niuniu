package internal

import (
	"server/msg"
	"gopkg.in/mgo.v2/bson"

	"github.com/name5566/leaf/log"
)

// 数据库的数据
type UserData struct {
	UserID		int	"_id"	// 用户id 自增型的
	AccountID	string		// 用户线上看到的id
	NickName	string		// 用户的昵称
	Sex 		int			// 性别 0--女 1--男
	TotalCount	int			// 比赛总次数
	WinCount	int			// 胜利次数
	Money		int			// 账号金币
	HeadImgUrl	string		// 头像
	CreatedAt	int64		// 注册时间
	UnionId		string		// 微信id
	AccessToken	string 		// token
}

const USERDB  = "users"


func login(user  *msg.UserLoginInfo)(err error) {
	var result UserData
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		// check user
		err := db.DB(DB).C(USERDB).Find(bson.M{
			"name"	: user.Name,
			"pwd"	: user.Pwd,
		}).One(&result)

		if err != nil{
			log.Fatal("login err - %v",err)
			return
		}
	}, func() {
	})
	return
}


