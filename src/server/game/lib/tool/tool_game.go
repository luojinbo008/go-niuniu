package tool
import (
	"server/msg"
	"github.com/name5566/leaf/gate"
	// "github.com/name5566/leaf/log"
)

type RoomConfigSlice struct {
	RoomConfigs []RoomConfig
}

type RoomConfig struct {
	Condition 		int
	EndPoints		int
	Online 			int
	AreaId			int
	AreaType		int
	AreaTypeName	string
	IsCivilian		int
}

func PushUserLoginInfo(a gate.Agent)  {

	// 推送用户基本数据
	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_CODE_SUCCESS,
			CMD		: msg.CMD_MY_USER_INFO,
			MSG		: "LOGIN SUCCESS",
			DATA	: a.UserData(),
		},
	)

	// 推送房间数据
	lineRoomList := gameCache.GetLineRoomList()

	var roomConfigSlice RoomConfigSlice

	for _, roomConfig := range lineRoomList.LineAreaRooms {

		roomConfigSlice.RoomConfigs = append(
			roomConfigSlice.RoomConfigs,
			RoomConfig{
				roomConfig.LimitGold,
				roomConfig.BaseGold,
				0,
				roomConfig.AreaId,
				roomConfig.AreaType,
				roomConfig.AreaTypeName,
				roomConfig.IsCivilian,
			},
		)
	}

	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_CODE_SUCCESS,
			CMD		: msg.CMD_ROOM_INFO,
			MSG		: "GET ROOM LIST SUCCESS",
			DATA	: roomConfigSlice.RoomConfigs,
		},
	)

}

// 进入房间
func EnterGameRoom(a gate.Agent, data *msg.EnterGameRoom) {

	if data.AreaId != 0 {
		data.Type = "chose"
	} else {
		a.WriteMsg(
			&msg.CodeState {
				CODE 	: msg.MSG_CODE_ERROR,
				CMD		: msg.CMD_ERROR,
				MSG		: "GAME IN DATA ERROR",
			},
		)
	}

	userData := a.UserData()

	if userData == nil {
		a.WriteMsg(
			&msg.CodeState {
				CODE 	: msg.MSG_CODE_ERROR,
				CMD		: msg.CMD_ERROR,
				MSG		: "ENTER ROOM GET USERDATA ERROR",
			},
		)
	}
	// 正在游戏 后补逻辑



	if data.Type == "chose" {
		GetRoomByAreaId(userData, data.AreaId)
	}
	a.WriteMsg(
		&msg.CodeState {
			CODE 	: msg.MSG_CODE_ERROR,
			CMD		: msg.CMD_ERROR,
			MSG		: "ENTER ROOM GET USERDATA ERROR",
		},
	)
}

// 用户离开房间专门处理
// action 1 断开连接 2 用户主动退出 3 用户被踢出房间
func UserLeaveDesk(userId int, action int) {

	// 获取用户信息
	lineUserInfo, _ := gameCache.GetLineUser(userId)

	// 判断用户是否在游戏
	if lineUserInfo.IsPlaying == 1 {

		// 正在游戏断线
		lineUserInfo.IsOutRoom = 1
	} else {

		// 不在游戏断线，数据清除
		if lineUserInfo.InDesk != 0 {

			// 在桌子上了，退出桌子
			gameCache.LeaveLineSeat(lineUserInfo.InArea, lineUserInfo.InRoom, lineUserInfo.InSeat, lineUserInfo.InDesk)

			// 桌子人数减一
			gameCache.LeaveLineDesk(lineUserInfo.InArea, lineUserInfo.InRoom, lineUserInfo.InDesk)

			// 场次人数减一
			gameCache.LeaveLineArea(lineUserInfo.InArea)

			// 房间人数减一
			gameCache.LeaveLineRoom(lineUserInfo.InArea, lineUserInfo.InRoom)

		}

		lineUserInfo.InArea 		= 0
		lineUserInfo.InRoom 		= 0
		lineUserInfo.InDesk 		= 0
		lineUserInfo.IsPlaying 		= 0
		lineUserInfo.IsWaiting 		= 0
		lineUserInfo.IsOutRoom 		= 0
		lineUserInfo.InSeat 		= 0
		lineUserInfo.InDeskStart	= 0

		if action == 1 {
			lineUserInfo.IsOnLine 	= 0
			lineUserInfo.IsCutLine 	= 1
		}
	}

	// 更新用户数据
	gameCache.ModifyLineUser(lineUserInfo)

	// 删除用户对应连接
	if action == 1 {
		gameCache.RemoveLineUserFdByUserId(userId)
	}

	// 用户主动退出
	if action == 2 {

		// 推送桌子数据
		LoopPushDeskUser(lineUserInfo.InDesk)

		// 推动数据给用户，带上用户基本数据
	}
}