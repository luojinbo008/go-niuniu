package tool

func LoopPushDeskUser(areaId int, roomId int, deskId int) {

	lineUserInfoList := getLoopUser(areaId, roomId, deskId)
	for _, lineUserInfo := range lineUserInfoList.LineUserInfos {

	}
}

func getLoopUser(areaId int, roomId int, deskId int) (lineUserInfoList LineUserInfoList){

	seatLineList :=  gameCache.GetLineDeskSeatList(areaId, roomId, deskId)

	for _, seatLine := range seatLineList.LineSeats {

		// 真实用户
		if seatLine.SeatType == 1 {
			lineUserInfo, err := gameCache.GetLineUser(seatLine.SeatUserId)
			if err != nil {
				if lineUserInfo.IsOutRoom != 0 {
					lineUserInfoList = append(lineUserInfoList.LineUserInfos, lineUserInfo)
				}
			}
		}
	}
	return lineUserInfoList
}

func pushDeskUser(areaId int, roomId int, deskId int, userId int) {
	seatLineList :=  gameCache.GetLineDeskSeatList(areaId, roomId, deskId)

	for _, seatLine := range seatLineList.LineSeats {
		
	}
}