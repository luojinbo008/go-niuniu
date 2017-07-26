package tool
import (
	"server/msg"
	"github.com/name5566/leaf/gate"
	// "github.com/name5566/leaf/log"
)

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
	roomList := gameCache.GetLineRoomList()
}