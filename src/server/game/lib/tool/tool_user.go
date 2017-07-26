package tool

import (
	"server/msg"
	"go.lib/wechat/oauth2"
	"github.com/name5566/leaf/log"
	"fmt"
	"server/game/lib/model"
)

type UserInfo struct {
	AccountID		string			// 用户线上看到的id
	NickName		string			// 用户的昵称
	Sex 			int				// 性别 0--未知 1--男 2--女
	HeadImgUrl		string			// 头像
	Diamond			int             // 钻石
	TotalCount		int				// 游戏总次数
	WinCount		int				// 游戏胜利次数
	Money			int				// 游戏金币
	AccessToken		string 			// token
}

// 微信端登陆
func WechatLogin(wechatLogin *msg.UserLoginByWechat) (userInfo UserInfo, err error) {

	var userData  		model.UserData
	var userThirdInfo  	model.ThirdInfoData
	var userGameData  	model.UserGameData

	wechatUserInfo, err := oauth2.GetAutoUserInfo(wechatLogin.Code)

	if err != nil {
		log.Error("login get thirdinfo - %s", err)
		return userInfo, err
	}

	if wechatUserInfo.Errcode != 0 {
		log.Error("login get thirdinfo - %s, %s", wechatUserInfo.Errcode, wechatUserInfo.Errmsg)
		return userInfo, fmt.Errorf(wechatUserInfo.Errmsg)
	}

	userThirdInfo, err = gameModel.GetThirdInfoUserInfoByThirdId(wechatUserInfo.Openid)

	if err != nil {
		if err.Error() == "not found" {

			userData.NickName = wechatUserInfo.Nickname
			userData.HeadImgUrl = wechatUserInfo.Headimgurl
			userData.Sex = wechatUserInfo.Sex

			userData, err = gameModel.AddUserInfo(userData)

			if err != nil {
				log.Error("login user add err - %v", err)
				return userInfo, err
			}

			userThirdInfo, err = gameModel.AddThirdInfoUserInfo(model.ThirdInfoData{
				userData.UserID,
				wechatUserInfo.Openid,
				"",
			})

			if err != nil {
				log.Error("login thirdinfo add err - %v", err)
				return userInfo, err
			}
		}
	} else {
		userData, err = gameModel.GetUserInfoByUserId(userThirdInfo.UserID)
		if err != nil {
			return userInfo, err
		}
	}

	userGameData, err = gameModel.GetUserGameData(userData.UserID)

	if err != nil {
		return userInfo, err
	}

	userInfo = UserInfo{
		userData.AccountID,
		userData.NickName,
		userData.Sex,
		userData.HeadImgUrl,
		userData.Diamond,
		userGameData.TotalCount,
		userGameData.WinCount,
		userGameData.Money,
		userData.AccessToken,
	}
	return userInfo, err
}

func ReLogin(userReLogin *msg.UserReLogin) (userInfo UserInfo, err error) {

	var userData		model.UserData
	var userGameData  	model.UserGameData

	userData, err = gameModel.GetUserInfoByAccessToken(userReLogin.AccountID, userReLogin.AccessToken)

	if err != nil {
		return userInfo, err
	}
	userGameData, err = gameModel.GetUserGameData(userData.UserID)

	if err != nil {
		return userInfo, err
	}

	userInfo = UserInfo{
		userData.AccountID,
		userData.NickName,
		userData.Sex,
		userData.HeadImgUrl,
		userData.Diamond,
		userGameData.TotalCount,
		userGameData.WinCount,
		userGameData.Money,
		userData.AccessToken,
	}

	return userInfo, err
}