package usermodel

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"time"
)

type CurrencyModel struct {
	mongodb.IDField `bson:",inline"`
	Balances        `bson:",inline"`
	UpdatedAt       time.Time `json:"lastModified" bson:"lastModified"`
}

type Balances struct {
	GAME_THC int64 `json:"gTHC" bson:"gTHC"`
	GAME_THG int64 `json:"gTHG" bson:"gTHG"`
	GAME_PP  int64 `json:"gPP" bson:"gPP"`
	GAME_PT  int64 `json:"gPT" bson:"gPT"`
}

func (model CurrencyModel) CollectionName() string {
	return "Currency"
}
