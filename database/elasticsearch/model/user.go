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

type UserStatMapping struct {
	User      interface{} `json:"user"`
	StatName  string      `json:"stat_name"`
	StatValue float64     `json:"stat_value"`
	Timestamp time.Time   `json:"@timestamp"`
}

func (u UserStatMapping) String() string {
	res, _ := json.Marshal(u)
	return string(res)
}

func (u UserStatMapping) GetIndexRequest() esapi.IndexRequest {
	return esapi.IndexRequest{
		Index: "user-stat",
		Body:  strings.NewReader(u.String()),
	}
}
