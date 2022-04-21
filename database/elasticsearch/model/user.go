package model

import "github.com/elastic/go-elasticsearch/v8/esapi"

var UserIndicesMapping = esapi.IndicesPutMappingRequest{
	Index: []string{"user"},
	Body:  nil,
}
