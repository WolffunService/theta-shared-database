package userstorage

import (
	"context"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

type UserStatResponse struct {
	UserId           string         `json:"userId"`
	UserDetail       usermodel.User `json:"userDetail"`
	AccountAge       int            `json:"accountAge"`
	HeroNFT          int            `json:"heroNFT"`
	PlayerBattle     int            `json:"playerBattle"`
	BattleFrequency  int            `json:"battleFrequency"`
	FirstOpenDate    bool           `json:"firstOpenDate"`
	GeoTier          int            `json:"geoTier"`
	CreatorViewPoint int            `json:"creatorViewPoint"`
	ConnectMKP       bool           `json:"connectMKP"`
	MKP              float64        `json:"mkp"`
	IAP              float64        `json:"iap"`
}

func FindUserStat(skip int64, limit int64) ([]UserStatResponse, error) {
	findOption := options.Find()
	findOption.SetSkip(skip)
	findOption.SetLimit(limit)

	filter := bson.M{
		"isBot": false,
	}

	users, err := findUserStat(context.Background(), filter, findOption)

	if err != nil {
		return nil, err
	}

	var userstats []UserStatResponse

	now := time.Now()
	for _, user := range users {
		userstat := UserStatResponse{
			UserId:          user.GetUserId(),
			UserDetail:      user,
			AccountAge:      int(user.CreatedAt.Sub(now).Hours() / 24.0),
			HeroNFT:         rand.Intn(10),
			PlayerBattle:    rand.Intn(210),
			BattleFrequency: rand.Intn(15),
			FirstOpenDate:   true,
			GeoTier:         1,
			ConnectMKP:      true,
			MKP:             float64(rand.Intn(6000)),
			IAP:             float64(rand.Intn(100)),
		}

		userstats = append(userstats, userstat)
	}

	return userstats, nil
}

func findUserStat(ctx context.Context, filter interface{}, findOptions ...*options.FindOptions) ([]usermodel.User, error) {
	var result []usermodel.User

	collection := mongodb.CollRead(&usermodel.User{})
	if err := collection.SimpleFindWithCtx(ctx, &result, filter, findOptions...); err != nil {
		log.Println("FindListUser Error ", err)
		return nil, err
	}

	return result, nil
}
