package pubsublytic

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/WolffunGame/theta-shared-database/database/pubsub/auditproto"
	"github.com/WolffunGame/theta-shared-database/database/pubsub/publisher"
	"google.golang.org/protobuf/encoding/protojson"
)

type simpleEventModel struct {
	Model *auditproto.SimpleEvent
}

func Recover() {
	if r := recover(); r != nil {
		log.Println("recovered from ", r)
	}
}

func PushCustomAnalytic(topicId string, eventName string, data ...interface{}) {
	go func() {
		defer Recover()
		event, err := CreateSimpleAnalyticEvent(eventName, data...)
		if err != nil {
			log.Println("[error][analytic] cannot marshal event special event reward")
			return
		}

		event.Push(context.TODO(), topicId)
	}()
}

func CreateSimpleAnalyticEvent(eventName string, data ...interface{}) (*simpleEventModel, error) {
	eventParams := []*auditproto.KeyPair{}
	metadata := map[string]string{}

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
			if k != "metadata" {
				eventParams = append(eventParams, &auditproto.KeyPair{
					Key:   k,
					Value: v.(string),
				})
			} else {
				x := v.(map[string]interface{})
				for k2, v2 := range x {
					metadata[k2] = v2.(string)
				}
			}
		}
	}

	return &simpleEventModel{
		Model: &auditproto.SimpleEvent{
			Event: &auditproto.SimpleEventContent{
				Timestamp:   time.Now().Unix(),
				EventName:   eventName,
				EventParams: eventParams,
			},
			Metadata: metadata,
		},
	}, nil
}

func (s *simpleEventModel) Push(ctx context.Context, topicId string) {
	if msg, err := protojson.Marshal(s.Model); err == nil {
		publisher.PublishMessage(ctx, topicId, msg)
	} else {
		log.Println("[error][analytic] cannot write event for", topicId)
	}
}
