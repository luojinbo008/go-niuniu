package init


import (
	"server/game/lib/cache"
	"github.com/name5566/leaf/log"
)

func InitAreaAndRoom() {

	areaList := gameModel.GetAreaList()
	roomList := gameModel.GetRoomList()

	roomMaxUser := 100
	for _, area := range areaList.Areas {
		roomNum := 0
		for _, room := range roomList.Rooms {
			if area.AreaId == room.AreaId {
				// 房间 配置
				err := gameCache.AddLineRoom(cache.LineRoomConfig{
					room.RoomId,
					room.AreaId,
					area.AreaType,
					area.IsCivilian,
					getAreaTypeName(area.AreaType),
					area.ServiceCost,
					area.BaseGold,
					area.LimitGold,
					room.DeskNum,
					roomMaxUser,
					0,				// 房间在线人数
					room.Status,
				})

				if err != nil {
					log.Error("init room err - %v", err)
				}
				roomNum ++
			}
		}

		err := gameCache.AddLineArea(cache.LineAreaConfig{
			area.AreaId,
			area.AreaType,
			area.IsCivilian,
			getAreaTypeName(area.AreaType),
			area.ServiceCost,
			area.BaseGold,
			area.LimitGold,
			0,
			roomMaxUser * roomNum,
		})
		if err != nil {
			log.Error("init area err - %v", err)
		}
	}
}


func InitRoomDesk() {
	lineAreaRoomList := gameCache.GetLineRoomList()
	for _, roomConfig := range lineAreaRoomList.LineAreaRooms {
		for i := 1; i <= roomConfig.DeskNum; i++ {

			// 设置每房间桌子
			err := gameCache.AddLineDesk(cache.LineDesk{
				i,
				roomConfig.RoomId,
				0,
			})

			if err != nil {
				log.Error("init Desk err - %v", err)
			}
		}
	}
}

func getAreaTypeName(AreaType int) string {

	switch AreaType {
		case 1:
			return "随机庄"
		case 2:
			return "炸金牛"
		default:
			return "未知"
	}
}
