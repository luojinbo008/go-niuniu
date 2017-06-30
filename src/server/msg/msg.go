package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

// 状态 常量标记
const (
	MSG_Register_Existed	= 0 	// 注册用户已存在
	MSG_Register_OK			= 1 	// 注册成功
	MSG_Login_Error			= 2 	// 登录失败 信息错误
	MSG_Login_OK			= 3 	// 登录成功
	MSG_DB_Error			= 111 	// 数据库出错

	//房间信息 1000开始标记
	MSG_ROOM_NOTAUTH		= 1001 	//没有权限
	MSG_ROOM_ERRORPWD		= 1002 	//密码错误
	MSG_ROOM_OVERVOLUME		= 1003 	//你已经在其他房间了 拒绝加入其他房间
	MSG_ROOM_NOMONEY		= 1004 	//起始资金不够
	MSG_ROOM_NOTEMPTY		= 1005 	//房子不空
	MSG_ROOM_NOROOM			= 1006 	//没有该房子记录
)

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&UserLoginInfo{})
}

type CodeState struct {
	MSG_STATE 	int 	// const
	Message 	string 	// 警告信息
}

type Hello struct {
	Name 	string
}

// 登录
type UserLoginInfo struct {
	Name 	string
	Pwd		string
}

type LoginError struct {
	State 	int
	Message	string
}



