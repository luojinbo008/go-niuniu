package tool

import (
	"github.com/name5566/leaf/log"
	"errors"
)

func GetRoomByAreaId(userInfo interface{}, areaId int) (x bool, err error)  {

	userData := userInfo.(UserInfo)

	userMoney := userData.Money

	// 如果 已在房间  待补充

	config, err := gameCache.GetAreaConfig(areaId)
	if err != nil {
		log.Error("Enter Room Get Area Config Error: %v" , err)
	}

	// 进入下线不符合
	if userMoney < config.BaseGold {
		log.Debug("Enter Room userMoney Limit")
		return false, errors.New("金币数量不满足进入条件")
	}

	// 场次最大容量
	areaMaxUserNum := config.MaxNum

	// 场次目前人数
	areaCurrentUserNum := config.Num

	if int(areaMaxUserNum) - 5 < int(areaCurrentUserNum) {
		log.Debug("Enter Room MaxUserNum Limit")
		return false, errors.New("房间人数已满")
	}

	lineAreaRoomList := gameCache.GetLineRoomList()

	for _, roomConfig := range lineAreaRoomList.LineAreaRooms {

		if roomConfig.Status != 1 {
			log.Debug("Enter Room And Room Status Stop")
			return false, errors.New("房间被禁用")
		}

		err = IncrUserNum(areaId, roomConfig.RoomId)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func IncrUserNum(areaId int, roomId int) (err error) {
	gameCache.IncrRoomUserNum(roomId)
	if err != nil {
		log.Debug("IncrUserNum Incr Room User Num Error")
		return errors.New("房间用户数量失败")
	}

	gameCache.IncrAreaUserNum(areaId)
	if err != nil {
		log.Debug("IncrUserNum Incr Area User Num Error")
		return errors.New("场次用户数量失败")
	}

	return nil
}