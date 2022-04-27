package userstorage

import (
	"context"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type PlayerStat struct {
	PlayerId              string            `json:"userId"`
	PlayerDetail          usermodel.NewUser `json:"userDetail"`
	PlayerStatAccount     usermodel.PlayerStatAccount
	PlayerStatMarketplace usermodel.PlayerStatMarketplace
	PlayerStatBalance     usermodel.PlayerStatBalance
}

func FindUserRankingWithOIds(ctx context.Context, oids []primitive.ObjectID) (map[string]usermodel.UserRanking, error) {
	var users []usermodel.UserRanking
	coll := mongodb.CollRead(&usermodel.UserRanking{})
	err := coll.FindByListIDWithCtx(ctx, oids, &users)

	userMap := make(map[string]usermodel.UserRanking)
	if err != nil {
		return userMap, err
	}

	for _, userRanking := range users {
		userMap[userRanking.ID.(primitive.ObjectID).Hex()] = userRanking
	}
	return userMap, nil
}

func FindUserCurrencyWithOIds(ctx context.Context, oids []primitive.ObjectID) (map[string]usermodel.CurrencyModel, error) {
	var currencys []usermodel.CurrencyModel
	coll := mongodb.CollRead(&usermodel.CurrencyModel{})
	err := coll.FindByListIDWithCtx(ctx, oids, &currencys)

	currencyMap := make(map[string]usermodel.CurrencyModel)
	if err != nil {
		return nil, err
	}

	for _, currency := range currencys {
		currencyMap[currency.ID.(primitive.ObjectID).Hex()] = currency
	}
	return currencyMap, nil
}

func FindUserStat(skip int64, limit int64) ([]*PlayerStat, error) {
	findOption := options.Find()
	findOption.SetSkip(skip)
	findOption.SetLimit(limit)

	filter := bson.M{
		"isBot":  nil,
		"status": usermodel.ACTIVE,
	}

	users, err := findUserStat(context.Background(), filter, findOption)
	if err != nil {
		return nil, err
	}

	var oids []primitive.ObjectID
	for _, user := range users {
		oids = append(oids, user.ID.(primitive.ObjectID))
	}

	ranking, err := FindUserRankingWithOIds(context.Background(), oids)
	if err != nil {
		return nil, err
	}

	currencys, err := FindUserCurrencyWithOIds(context.Background(), oids)
	if err != nil {
		return nil, err
	}

	var playerStats []*PlayerStat
	now := time.Now().UTC()

	for _, user := range users {
		var playerStatAccount usermodel.PlayerStatAccount
		var playerStatBalance usermodel.PlayerStatBalance

		if rank, found := ranking[user.GetUserId()]; found {
			playerStatAccount.FirstOpenDate = rank.CreateAt
			playerStatAccount.AccountDate = user.CreatedAt
			playerStatAccount.AccountAge = int32(now.Sub(user.CreatedAt).Hours() / 24.0)
		}

		if currency, found := currencys[user.GetUserId()]; found {
			playerStatBalance.THGBalance = currency.GAME_THG
			playerStatBalance.THCBalance = currency.GAME_THC
			playerStatBalance.PPBalance = currency.GAME_PP
			playerStatBalance.PTBalance = currency.GAME_PT
		}

		playerStat := &PlayerStat{
			PlayerId:              user.GetUserId(),
			PlayerDetail:          user,
			PlayerStatAccount:     playerStatAccount,
			PlayerStatMarketplace: usermodel.PlayerStatMarketplace{},
			PlayerStatBalance:     playerStatBalance,
		}
		playerStats = append(playerStats, playerStat)
	}
	return playerStats, nil
}

func findUserStat(ctx context.Context, filter interface{}, findOptions ...*options.FindOptions) ([]usermodel.NewUser, error) {
	var result []usermodel.NewUser

	collection := mongodb.CollRead(&usermodel.NewUser{})
	if err := collection.SimpleFindWithCtx(ctx, &result, filter, findOptions...); err != nil {
		log.Println("FindListUser Error ", err)
		return nil, err
	}

	return result, nil
}
