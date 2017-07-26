package model

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"server/conf"
)

const DB  = "main"

func (model *Model) InitModel() {
	if conf.Server.DBMaxConnNum <= 0 {
		conf.Server.DBMaxConnNum = 100
	}
	db, err := mongodb.Dial(conf.Server.DBUrl, conf.Server.DBMaxConnNum)

	if err != nil {
		log.Fatal("dial mongodb error: %v", err)
	}

	model.mongoDB = db

	err = db.EnsureCounter(DB, "counters", "users")
	if err != nil {
		log.Fatal("ensure counter error: %v", err)
	}

	err = db.EnsureCounter(DB, "counters", "nn_area")
	if err != nil {
		log.Fatal("ensure counter error: %v", err)
	}

	err = db.EnsureCounter(DB, "counters", "nn_room")
	if err != nil {
		log.Fatal("ensure counter error: %v", err)
	}
}

type Model struct {
	mongoDB 	*mongodb.DialContext
}

func (model *Model)mongoDBDestroy() {
	model.mongoDB.Close()
	model.mongoDB = nil
}

func (model *Model) mongoDBNextSeq(id string) (int, error) {
	return model.mongoDB.NextSeq(DB, "counters", id)
}
