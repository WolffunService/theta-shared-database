package usermodel

import "github.com/WolffunGame/theta-shared-database/database/mongodb"

type RentState int

const (
	RS_FOR_RENT      RentState = 1
	RS_RENTED        RentState = 2
	RS_EXPIRED       RentState = 3
	RS_STOP_FOR_RENT RentState = 4
)

type HeroRentInfo struct {
	mongodb.DefaultModel `bson:",inline"`
	TokenId              string    `json:"tokenId" bson:"tokenId"`
	RefId                string    `json:"refId" bson:"refId"`
	TokenName            string    `json:"tokenName" bson:"tokenName"`
	SkinId               int       `json:"skinId" bson:"skinId"`
	DONT_USE_OwnerId     string    `json:"ownerId" bson:"ownerId"` // DONT USE
	State                RentState `json:"state" bson:"state"`
	LastModified         int64     `json:"lastModified" bson:"lastModified"`
	CreatedTime          int64     `json:"createdTime" bson:"createdTime"`
	RentedTime           int64     `json:"rentedTime" bson:"rentedTime"`
	ExpiredTime          int64     `json:"expiredTime" bson:"expiredTime"`
	RenterAddress        string    `json:"renterAddress" bson:"renterAddress"`
	RentBattles          int       `json:"rentBattles" bson:"rentBattles"`
	TransactionHash      string    `json:"transactionHash" bson:"transactionHash"`
}

func (HeroRentInfo) CollectionName() string {
	return "HeroRentInfos"
}
