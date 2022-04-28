package usermodel

import "github.com/WolffunGame/theta-shared-database/database/mongodb"

func (Hero) CollectionName() string {
	return "Heroes"
}

type Hero struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `bson:",inline"`
	OwnerId              string        `json:"userId" bson:"userId"`
	OwnerAddress         string        `json:"ownerAddress" bson:"ownerAddress"`
	OwnerName            string        `json:"ownerName" bson:"ownerName"`
	UserUsingId          string        `json:"userUsingId" bson:"userUsingId"`
	HeroTypeId           int           `json:"heroTypeId" bson:"heroTypeId"`
	SkinId               int           `json:"skinId" bson:"skinId"`
	Status               HeroStatus    `json:"status" bson:"status"`
	Season               int           `json:"season" bson:"season"`
	Level                int           `json:"level" bson:"level"`
	Trophy               int           `json:"trophy" bson:"trophy"` //totoal trophie
	RefId                string        `json:"refId" bson:"refId"`
	TokenId              string        `json:"tokenId" bson:"tokenId"`
	NftId                string        `json:"nftId" bson:"nftId"`
	RentInfo             *HeroRentInfo `json:"rentInfo,omitempty" bson:"rentInfo,omitempty"`
}

type HeroStatus int

const (
	HS_DEFAULT                HeroStatus = 0
	HS_NOT_MINT               HeroStatus = 1
	HS_MINTING                HeroStatus = 2
	HS_AVAILABLE              HeroStatus = 3
	HS_RENTED                 HeroStatus = 4
	HS_ON_MARKET              HeroStatus = 10
	HS_FOR_RENT               HeroStatus = 11
	HS_IN_TRANSACTION_PROCESS HeroStatus = 20
	HS_IN_STAKING             HeroStatus = 30
	HS_LOCK                   HeroStatus = 99
)
