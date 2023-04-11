package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	es "github.com/WolffunService/theta-shared-database/database/elasticsearch"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
	"log"
)

func main() {
	es.InitES(nil)

	es.InitES(&elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})

	var buf bytes.Buffer
	query := es.Map{
		"size": 0,
		"query": es.Map{
			"match": es.Map{
				"user.id": "628740e25607ef4ec74e0b2c",
			},
		},
		"aggs": es.Map{
			"sum_death": es.Map{
				"sum": es.Map{
					"field": "death",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error getting response: %s", err)
		return
	}

	// Perform the search request.
	res, err := es.GetClient().Search(
		es.GetClient().Search.WithContext(context.Background()),
		es.GetClient().Search.WithIndex("playerstats*"),
		es.GetClient().Search.WithBody(&buf),
		es.GetClient().Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return
	}

	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()

	//Cach 1 su dung go json
	var r es.Map
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return
	}
	fmt.Print("Tong so lan chet la: ")
	fmt.Println(r["aggregations"].(map[string]interface{})["sum_death"].(map[string]interface{})["value"])

	//Cach 2 su dung gjson
	//Luu y muon test cach 2 phai comment cach 1 tai cai EOF no seek toi cuoi' roi hihihi
	var b bytes.Buffer
	b.ReadFrom(res.Body)

	fmt.Print("Tong so lan chet la: ")
	fmt.Println(gjson.GetBytes(b.Bytes(), "aggregations.sum_death.value"))

	//for _, hit := range r["aggregations"].(map[string]interface{})["sum-death"].(map[string]interface{}) {
	//	resultItem := hit.(map[string]interface{})
	//
	//	fmt.Println(resultItem)
	//
	//	//log.Printf("ID=%s, %v", resultItem["_id"], resultItem["_source"].(map[string]interface{})["ingame_mode"])
	//}

}
