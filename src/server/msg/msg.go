package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

// 状态 常量标记
const (
	MSG_CODE_SUCCESS	= 0 	// 成功
	MSG_CODE_ERROR		= 1 	// 失败
)

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&UserLoginByWechat{})
	Processor.Register(&CodeState{})
	Processor.Register(&UserReLogin{})
}

type CodeState struct {
	CODE 		int 			// const
	CMD			int				// const
	MSG 		string 			// 警告信息
	DATA		interface{}	//
}

type Hello struct {
	Name 	string
}

// 微信端登录
type UserLoginByWechat struct {
	Code 		string
}

// 根据AccessToken 登陆
type UserReLogin struct {
	AccountID string
	AccessToken string
}
