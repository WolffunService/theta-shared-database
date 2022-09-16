package pubsublytic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/WolffunGame/theta-shared-database/database/pubsub/publisher"
)

type keyPair struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type EventJSONModel struct {
	EventName   string    `json:"event_name,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	UserId      string    `json:"user_id,omitempty"`
	Country     string    `json:"country,omitempty"`
	EventParams []keyPair `json:"event_params,omitempty"`
}

func CreateJSONEvent(eventName string, country string, userId string, data ...interface{}) (*EventJSONModel, error) {
	eventParams := []keyPair{}

	for i := 0; i < len(data); i++ {
		eventParamsMap := map[string]interface{}{}

		// marshal fk struct to json
		bytes, err := json.Marshal(data[i])
		if err != nil {
			return nil, err
		}

		// unmarshal json to map[string]string
		err = json.Unmarshal(bytes, &eventParamsMap)
		if err != nil {
			return nil, err
		}

		// loop through all field to tách metadata với analytic fields
		for k, v := range eventParamsMap {
			eventParams = append(eventParams, keyPair{
				Key:   k,
				Value: fmt.Sprintf("%v", v),
			})
		}
	}

	return &EventJSONModel{
		Timestamp:   time.Now(),
		EventName:   eventName,
		UserId:      userId,
		Country:     country,
		EventParams: eventParams,
	}, nil
}

func (s *EventJSONModel) Push(ctx context.Context, topicId string) error {
	msg, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return publisher.PublishMessage(ctx, topicId, msg)
}

func PushJSONAnalytic(topicId string, eventName string, country string, userId string, data ...interface{}) error {
	defer Recover()
	event, err := CreateJSONEvent(eventName, country, userId, data...)
	if err != nil {
		log.Println("[error][analytic] cannot marshal event special event reward")
		return err
	}

	return event.Push(context.TODO(), topicId)
}
