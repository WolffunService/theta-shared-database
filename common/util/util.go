package util

import (
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromHex(s string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return objID
}

func GetBehaviorStatus(u *usermodel.User) usermodel.BehaviorStatus {
	// Behavior Point
	// 70 - 100: EXCELLENT
	// 50 - 69: GOOD
	// < 50: BAD

	bPoint := GetBehaviorStatus(u)
	if bPoint >= 70 {
		return usermodel.EXCELLENT
	}
	if bPoint >= 50 {
		return usermodel.GOOD
	}
	return usermodel.BAD
}

func GetBanMultiple(b usermodel.BehaviorStatus) int {
	switch b {
	case usermodel.BAD:
		return 3
	case usermodel.GOOD:
		return 2
	default:
		return 1
	}

}
