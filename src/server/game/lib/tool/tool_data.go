package tool

import (
	"server/game/lib/cache"
)

type UserInfo struct {
	UserId			int				// 用户id
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

type LineUserInfoList struct {
	LineUserInfos []cache.LineUserInfo
}