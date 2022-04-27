package userstorage

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"log"
	"testing"
)

func TestFindUserStat(t *testing.T) {
	dbConfig := &mongodb.MongoConfig{
		DbName:            "thetan",
		ConnectionUrl:     "mongodb://thetan:e54a273ab740ae44e3e9@127.0.0.1:27017/thetan?authSource=thetan&replicaSet=thetan-data-rs&readPreference=primary&directConnection=true&ssl=false",
		MaxConnectionPool: 1000,
	}

	ctxDB, client, CloseF := mongodb.ConnectMongoWithConfig(dbConfig, nil)
	defer CloseF()
	defer client.Disconnect(ctxDB)

	res, _ := FindUserStat(1000, 1000)

	for _, user := range res {
		log.Printf("%+v \n", user)
	}
}
