package cache

const (
	LINEDESKDB				= "line_nn_desk"
)

type LineDesk struct {
	DeskId		int "_id"
	RoomId 		int
	Waiting		int
}

func (cache *Cache) AddLineDesk(lineDesk LineDesk) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEDESKDB).Insert(lineDesk)
	return err
}