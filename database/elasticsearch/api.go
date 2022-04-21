package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var c *elasticsearch.Client

func InitES() {
	cfg := elasticsearch.Config{
		//TODO Get from env

		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "talaconma",
	}

	var err error
	c, err = elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	Info()
}

func Info() {
	res, err := c.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}

func GetClient() *elasticsearch.Client {
	if c == nil {
		InitES()

		return c
	}

	return c
}
