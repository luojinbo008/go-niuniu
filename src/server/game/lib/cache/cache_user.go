package cache

type LineUserInfo struct {
	AccountID		string			// 用户线上看到的id
	NickName		string			// 用户的昵称
	Sex 			int				// 性别 0--未知 1--男 2--女
	HeadImgUrl		string			// 头像
	Diamond			int             // 钻石
	TotalCount		int				// 游戏总次数
	WinCount		int				// 游戏胜利次数
	Money			int				// 游戏金币
	AccessToken		string 			// token
	IsPlaying		int				// 是否在游戏
	IsOutRoom		int				// 是否在房间
}

const (
	LINEUSERDB				= "line_nn_user"
)

func (cache *Cache) AddLineUser(userInfo LineUserInfo) (err error) {

	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEUSERDB).Insert(userInfo)
	return err
}

func cache()  {
	
}