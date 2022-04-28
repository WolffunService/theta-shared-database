package userstorage

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"testing"
)

func connect() {
	dbConfig := &mongodb.MongoConfig{
		DbName:            "thetan",
		ConnectionUrl:     "mongodb://thetan:e54a273ab740ae44e3e9@127.0.0.1:27017/thetan?authSource=thetan&replicaSet=thetan-data-rs&readPreference=primary&directConnection=true&ssl=false",
		MaxConnectionPool: 1000,
	}

	mongodb.ConnectMongoWithConfig(dbConfig, nil)

}

func connectAudit() {
	dbConfig := &mongodb.MongoConfig{
		DbName:            "thetan",
		ConnectionUrl:     "mongodb://thetan:e8e01213532dc96be8f0@localhost:27017/thetan?replicaSet=thetan-autdit-rs&readPreference=nearest&authSource=thetan&ssl=false&appname=MongoDB%20Compass&directConnection=true&ssl=false",
		MaxConnectionPool: 1000,
	}

	mongodb.ConnectMongoWithConfig(dbConfig, nil)

}
func TestFindUserStat(t *testing.T) {
	connect()
	res, _ := FindUserRentalValueByUserId(context.Background(), "61d3f417cc3ca33c218c5c0b")
	for _, a := range res {
		fmt.Println(a)
	}
}

func TestFindHeroes(t *testing.T) {
	connect()
	res, _ := FindHeroesNFTByUserId(context.Background(), "61d3f417cc3ca33c218c5c0b")
	for _, a := range res {
		fmt.Printf("%+v\n", a)
	}
}

func TestFindHeroesUnavailable(t *testing.T) {
	connect()

	res, _ := FindHeroesNFTUnavailableByUserId(context.Background(), "61d3f417cc3ca33c218c5c0b")
	for _, a := range res {
		fmt.Printf("%+v\n", a)
	}
}

func TestFindBuyHeroEvent(t *testing.T) {
	connectAudit()

	res, _ := BuyHeroSuccessEvent(context.Background(), "61d7c7bce89e8850b84cf458")
	for _, a := range res {
		fmt.Printf("%+v\n", a)
	}
}

func TestFindBuyBoxEvent(t *testing.T) {
	connectAudit()

	res, _ := BuyBoxSuccessEvent(context.Background(), "619ba130fe775c92cca395d0")
	for _, a := range res {
		fmt.Printf("%+v\n", a)
	}
}
