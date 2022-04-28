package userstorage

import (
	"context"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"thetansm/enums/heroenum"
	"thetansm/models/auditmodel"
	"time"
)

// FindUserRankingWithOIds tìm thông tin user ranking dựa vào objectId
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

// FindUserCurrencyWithOIds tìm thông tin user currency dựa vào objectId
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

// FindUserRentalValueByUserId tìm thông tin rental của user dựa vào ID
func FindUserRentalValueByUserId(ctx context.Context, userId string) ([]usermodel.HeroRentInfo, error) {
	collection := mongodb.CollRead(&usermodel.HeroRentInfo{})
	var result []usermodel.HeroRentInfo

	filter := bson.M{
		"ownerId": userId,
	}

	if err := collection.SimpleFindWithCtx(ctx, &result, filter); err != nil {
		return nil, err
	}

	return result, nil
}

// FindUserStat tìm thông tin player stat
func FindUserStat(skip int64, limit int64) ([]*usermodel.PlayerStat, error) {
	findOption := options.Find()
	findOption.SetSkip(skip)
	findOption.SetLimit(limit)

	filter := bson.M{
		"isBot":  nil,
		"status": usermodel.ACTIVE,
	}

	users, err := FindUsers(context.Background(), filter, findOption)
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

	var playerStats []*usermodel.PlayerStat
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

		playerStat := &usermodel.PlayerStat{
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

// FindHeroesNFTByUserId tìm toàn bộ user nft dựa vào userid
func FindHeroesNFTByUserId(ctx context.Context, userId string) ([]usermodel.Hero, error) {
	collection := mongodb.CollRead(&usermodel.Hero{})
	var result []usermodel.Hero
	//var heroesNFT []usermodel.Hero
	filter := bson.M{
		"userId": userId,
		"status": bson.M{"$ne": 0},
	}

	if err := collection.SimpleFindWithCtx(ctx, &result, filter); err != nil {
		return nil, err
	}

	return result, nil
}

// FindHeroesNFTUnavailableByUserId tìm toàn bộ user nft nhưng không available dựa theo userid
func FindHeroesNFTUnavailableByUserId(ctx context.Context, userId string) ([]usermodel.Hero, error) {
	collection := mongodb.CollRead(&usermodel.Hero{})
	var result []usermodel.Hero
	filter := bson.M{
		"userId": userId,
		"status": bson.M{
			"$in": []int32{4, 10, 11}},
	}

	if err := collection.SimpleFindWithCtx(ctx, &result, filter); err != nil {
		return nil, err
	}

	return result, nil
}

// FindUsers tìm thông tin user dựa vào filter
func FindUsers(ctx context.Context, filter interface{}, findOptions ...*options.FindOptions) ([]usermodel.NewUser, error) {
	var result []usermodel.NewUser

	collection := mongodb.CollRead(&usermodel.NewUser{})
	if err := collection.SimpleFindWithCtx(ctx, &result, filter, findOptions...); err != nil {
		log.Println("FindListUser Error ", err)
		return nil, err
	}

	return result, nil
}

// BuyBoxSuccessEvent tìm toàn bộ thông tin transaction mua box của user
func BuyBoxSuccessEvent(ctx context.Context, userId string) ([]auditmodel.BoxEvent, error) {
	collection := mongodb.CollectionByName("BoxAudits")

	var result []auditmodel.BoxEvent
	filter := bson.M{
		"userId":           userId,
		"boxPurchaseEvent": bson.M{"$ne": nil},
	}

	if err := collection.SimpleFindWithCtx(ctx, &result, filter); err != nil {
		return nil, err
	}

	return result, nil
}

// BuyHeroSuccessEvent tìm toàn bộ thông tin transaction mua hero của user
func BuyHeroSuccessEvent(ctx context.Context, userId string) ([]auditmodel.HeroEvent, error) {
	collection := mongodb.CollectionByName("HeroAudits")

	var result []auditmodel.HeroEvent
	filter := bson.M{
		"userId":    userId,
		"tradeInfo": bson.M{"$ne": nil},
		"status":    heroenum.HESt_Succeeded,
	}

	if err := collection.SimpleFindWithCtx(ctx, &result, filter); err != nil {
		return nil, err
	}

	return result, nil
}
