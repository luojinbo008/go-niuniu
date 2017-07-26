package cache

type LineAreaConfig struct {
	AreaId 			int "_id"
	AreaType		int
	IsCivilian		int
	AreaTypeName	string
	ServiceCost		int
	BaseGold		int
	LimitGold		int
	Num				int
	MaxNum 			int
}

const (
	LINEAREACONFIGDB 	= "line_nn_area_config"
)

func (cache *Cache) AddLineArea(lineAreaConfig LineAreaConfig) (err error) {
	db := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db)

	err = db.DB(DB).C(LINEAREACONFIGDB).Insert(lineAreaConfig)
	return err
}