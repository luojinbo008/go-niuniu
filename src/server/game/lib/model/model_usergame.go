package model

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserGameData struct {
	UserID		int			// 用户id
	TotalCount	int			// 游戏总次数
	WinCount	int			// 游戏胜利次数
	Money		int			// 游戏金币
	CreatedAt	int64		// 第一次进入游戏时间
}

const (
	GAMEINFO = "nn_game"
)

func (model *Model) GetUserGameData(UserId int) (userGameData UserGameData, err error) {

	db := model.mongoDB.Ref()
	defer model.mongoDB.UnRef(db)

	err = db.DB(DB).C(GAMEINFO).Find(bson.M{
		"userid"	: UserId,
	}).One(&userGameData)

	if err != nil {
		if err.Error() == "not found" {
			userGameData.UserID = UserId
			userGameData.CreatedAt = time.Now().Unix()
			err = db.DB(DB).C(GAMEINFO).Insert(userGameData)

			if err != nil {
				log.Error("create game user add err - %v", err)
				return userGameData, err
			}
			return userGameData, nil
		}
	}
	return userGameData, err
}