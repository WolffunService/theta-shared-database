package userstorage

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"log"
	"testing"
)

func TestFindUserStat(t *testing.T) {
	dbConfig := &mongodb.MongoConfig{
		DbName:            "thetan",
		ConnectionUrl:     "mongodb://thetan:3c327f016341878ab21b@localhost:27017/?authSource=thetan&readPreference=primary&directConnection=true&ssl=false",
		MaxConnectionPool: 1000,
	}

	ctxDB, client, CloseF := mongodb.ConnectMongoWithConfig(dbConfig, nil)
	defer CloseF()
	defer client.Disconnect(ctxDB)

	res, _ := FindUserStat(1000, 1000)

	for _, user := range res {
		log.Println(user)
	}
}
