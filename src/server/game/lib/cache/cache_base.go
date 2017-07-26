package cache

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"server/conf"
)

const DB  = "games_cache"

func (cache *Cache) InitCache() {
	if conf.Server.DBMaxConnNum <= 0 {
		conf.Server.DBMaxConnNum = 100
	}

	db, err := mongodb.Dial(conf.Server.DBUrl, conf.Server.DBMaxConnNum)

	if err != nil {
		log.Fatal("dial mongodb error: %v", err)
	}
	cache.mongoDB = db


	db1 := cache.mongoDB.Ref()
	defer cache.mongoDB.UnRef(db1)

	db1.DB(DB).DropDatabase()
}

type Cache struct {
	mongoDB 	*mongodb.DialContext
}

func (cache *Cache) mongoDBDestroy() {
	cache.mongoDB.Close()
	cache.mongoDB = nil
}

func (cache *Cache) mongoDBNextSeq(id string) (int, error) {
	return cache.mongoDB.NextSeq(DB, "counters", id)
}
