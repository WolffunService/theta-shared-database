package main

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-database/database/elasticsearch"
	"github.com/WolffunGame/theta-shared-database/database/elasticsearch/model"
	"log"
	"math/rand"
	"time"
)

type UserModel struct {
	ID   string `json:"id"`
	Mail string `json:"mail"`
}

func main() {
	for i := 1; i <= 10000; i++ {
		if i%2 == 0 {
			continue
		}

		u := model.UserStatMapping{
			User: UserModel{
				ID:   fmt.Sprintf("fake-mongo-object-id-%d", i),
				Mail: fmt.Sprintf("fake_mail_%d@gmail.com", i),
			},
			StatName:  "battle_count",
			StatValue: 1,
			Timestamp: ranDate(),
		}

		res, err := u.GetIndexRequest().Do(context.Background(), elasticsearch.GetClient())

		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		res.Body.Close()
	}
}

func ranDate() time.Time {
	min := time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 4, 25, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
