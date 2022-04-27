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
	for i := 1; i <= 3000; i++ {
		u := model.UniversalUserStatMapping{
			User: UserModel{
				ID:   fmt.Sprint("fake-mongo-object-id-1021"),
				Mail: fmt.Sprintf("fake_mail_%d@gmail.com", i),
			},
			StatName:  "battle_count  ",
			StatValue: 1,
			Timestamp: ranDate(),
		}

		res, err := u.GetIndexRequest().Do(context.Background(), elasticsearch.GetClient())

		fmt.Println(res)

		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		time.Sleep(10 * time.Millisecond)

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
