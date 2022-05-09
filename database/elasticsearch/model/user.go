package model

import (
	"time"
)

//var UserIndicesMapping = esapi.IndicesPutMappingRequest{
//	Index: []string{"user"},
//	Body:  nil,
//}

//Nhóm stat battle, earning được track sau mỗi battle, loại stat này có dạng append-only time series nên sẽ sử dụng
//Data stream để index

//

type StatName string

// UserModel Các stat cố định của player sẽ define tại struct này để phục vụ filtering
type UserModel struct {
	ID              string    `json:"id"`
	Mail            string    `json:"mail"`
	Country         string    `json:"country"`
	Created         time.Time `json:"created"`
	GameOpened      time.Time `json:"game_opened"`
	WalletConnected time.Time `json:"wallet_connected"`
}

type BattleStatMapping struct {
	User       UserModel `json:"user"`
	IngameMode string    `json:"ingame_mode"`
	Mode       string    `json:"mode"`
	Hero       string    `json:"hero"`
	Skill1     string    `json:"skill_1"`
	Skill2     string    `json:"skill_2"`
	Result     string    `json:"result"`
	Trophy     int       `json:"trophy"`
	Region     string    `json:"region"`
	Timestamp  time.Time `json:"@timestamp"`
}

func (BattleStatMapping) Index() string {
	return "playerstats-battle"
}

type IAPStatMapping struct {
	User       UserModel `json:"user"`
	PackID     string    `json:"pack_id"` //enum
	Source     string    `json:"source"`  //enum
	Price      int64     `json:"price"`   //
	PriceType  string    `json:"price_type"`
	PowerPoint int       `json:"power_point"`
	HeroID     string    `json:"hero_id"` //object id, ko trùng
	SkinID     string    `json:"skin_id"` //enum, có trùng
	Timestamp  time.Time `json:"@timestamp"`
}

func (IAPStatMapping) Index() string {
	return "playerstats-iap"
}

type MarketPlaceStatMapping struct {
	User       UserModel `json:"user"`
	MarketType string    `json:"market_type"`
	Action     string    `json:"action"`
	RefType    string    `json:"ref_type"`
	TokenID    string    `json:"token_id"`
	ItemID     string    `json:"item_id"`
	Price      int       `json:"price"`
	Currency   string    `json:"currency"`
	Country    string    `json:"country"`

	HeroTypeID int    `json:"hero_type_id"`
	SkinRarity string `json:"skin_rarity"`
	HeroRarity string `json:"hero_rarity"`
	HeroLevel  string `json:"hero_level"`
	Trophy     string `json:"trophy"`

	ItemRarity string    `json:"item_rarity"`
	ItemTypeID int       `json:"item_type_id"` //có trùng
	ItemType   string    `json:"item_type"`    //avatar_frame, cosmetic vv, có trùng
	Timestamp  time.Time `json:"@timestamp"`
}

func (MarketPlaceStatMapping) Index() string {
	return "playerstats-marketplace"
}

type WalletStatMapping struct {
	User UserModel `json:"user"`
	THC  int64     `json:"thc"`
	THG  int64     `json:"thg"`
}

type UniversalUserStatMapping struct {
	User      interface{} `json:"user"`
	StatName  string      `json:"stat_name"`
	StatValue float64     `json:"stat_value"`
	Timestamp time.Time   `json:"@timestamp"`
}
