package tool

import (
	"server/msg"
	"github.com/name5566/leaf/gate"
	"go.lib/wechat/oauth2"
	"github.com/name5566/leaf/log"
	"fmt"
	"server/game/lib/model"
	"server/game/lib/cache"
)

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
		userData.UserID,
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
		userData.UserID,
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


func LineUserCut(fd string, userId int) (err error) {
	lineUserInfo, err := gameCache.GetLineUser(userId)
	if err != nil {
		return err
	}
	lineUserInfo.Fd = ""
	lineUserInfo.Ip = ""
	err = gameCache.ModifyLineUser(lineUserInfo)
	return err
}

func LineUserModify(a gate.Agent) (err error) {
	userData := a.UserData()
	data := userData.(UserInfo)
	var lineUserInfo cache.LineUserInfo

	lineUserInfo.Ip = "192.168.56.101"
	lineUserInfo.Fd = fmt.Sprint("%s", &a)
	lineUserInfo.UserId = data.UserId

	err = gameCache.ModifyLineUser(lineUserInfo)
	log.Debug("%v", err)
	return err
}