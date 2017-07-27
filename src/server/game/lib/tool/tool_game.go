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
