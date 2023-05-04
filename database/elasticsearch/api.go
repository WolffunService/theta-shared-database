package elasticsearch

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var c *elasticsearch.Client

func InitES(cfg *elasticsearch.Config) {
	if cfg == nil {
		cfg = &elasticsearch.Config{
			//TODO Get from env

			Addresses: []string{
				"http://localhost:9200",
			},
			Username: "elastic",
			Password: "talaconma",
		}
	}

	var err error
	c, err = elasticsearch.NewClient(*cfg)

	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return
	}

	Info()
}

func Info() {
	res, err := c.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return
	}

	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	log.Println(res)
}

func GetClient() *elasticsearch.Client {
	if c == nil {
		InitES(nil)

		return c
	}

	return c
}

type IndexMapping interface {
	Index() string
}

func GetIndexRequest(indexMapping IndexMapping) esapi.IndexRequest {
	res, _ := json.Marshal(indexMapping)
	m := string(res)
	return esapi.IndexRequest{
		Index: indexMapping.Index(),
		Body:  strings.NewReader(m),
	}

}

type Map map[string]interface{}
