package model

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
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
	ID           string    `json:"id"`
	Mail         string    `json:"mail"`
	Country      string    `json:"country"`
	Created      time.Time `json:"created"`
	Opened       time.Time `json:"opened"`
	WalletLinked time.Time `json:"wallet_linked"`
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
}

type WalletStatMapping struct {
	User UserModel `json:"user"`
	THC  int64     `json:"thc"`
	THG  int64     `json:"thg"`
}

type IAPStatMapping struct {
	User  UserModel `json:"user"`
	Value int64     `json:"value"`
}

type UniversalUserStatMapping struct {
	User      interface{} `json:"user"`
	StatName  string      `json:"stat_name"`
	StatValue float64     `json:"stat_value"`
	Timestamp time.Time   `json:"@timestamp"`
}

func (u UniversalUserStatMapping) String() string {
	res, _ := json.Marshal(u)
	return string(res)
}

func (u UniversalUserStatMapping) GetIndexRequest() esapi.IndexRequest {
	return esapi.IndexRequest{
		Index: "user-stat",
		Body:  strings.NewReader(u.String()),
	}
}

func (u BattleStatMapping) Index() string {
	return "playerstats-battle"
}
