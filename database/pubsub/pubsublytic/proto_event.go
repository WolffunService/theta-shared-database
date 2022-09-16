package pubsublytic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/WolffunGame/theta-shared-database/database/pubsub/auditprotobuf"
	"github.com/WolffunGame/theta-shared-database/database/pubsub/publisher"
	"google.golang.org/protobuf/encoding/protojson"
)

type simpleEventModel struct {
	Model *auditprotobuf.SimpleEvent2
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

// simpleEventModel
func CreateSimpleAnalyticEvent(eventName string, data ...interface{}) (*simpleEventModel, error) {
	eventParams := []*auditprotobuf.KeyPair2{}
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
				eventParams = append(eventParams, &auditprotobuf.KeyPair2{
					Key:   k,
					Value: fmt.Sprintf("%v", v),
				})
			} else {
				x := v.(map[string]interface{})
				for k2, v2 := range x {
					metadata[k2] = fmt.Sprintf("%v", v2)
				}
			}
		}
	}

	return &simpleEventModel{
		Model: &auditprotobuf.SimpleEvent2{
			Event: &auditprotobuf.SimpleEventContent2{
				Timestamp:   time.Now().Unix(),
				EventName:   eventName,
				EventParams: eventParams,
			},
			Metadata: metadata,
		},
	}, nil
}

// analytic
func CreateSimpleAnalyticEventV2(eventName string, data ...interface{}) (*analytic, error) {
	eventParams := []*auditprotobuf.KeyPair2{}

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

		// loop through all field to collect analytic fields
		for k, v := range eventParamsMap {
			eventParams = append(eventParams, &auditprotobuf.KeyPair2{
				Key:   k,
				Value: fmt.Sprintf("%v", v),
			})
		}
	}

	return &analytic{
		Timestamp:    time.Now().Unix(),
		EventName:    eventName,
		AnalyticData: eventParams,
	}, nil
}

func (s *simpleEventModel) Push(ctx context.Context, topicId string) {
	if msg, err := protojson.Marshal(s.Model); err == nil {
		publisher.PublishMessage(ctx, topicId, msg)
	} else {
		log.Println("[error][analytic] cannot write event for", topicId)
	}
}

func (s *Auditlytic[T]) Push(ctx context.Context, topicId string) {
	if msg, err := json.Marshal(*s); err == nil {
		publisher.PublishMessage(ctx, topicId, msg)
	} else {
		log.Println("[error][pubsub] cannot publish event for", topicId)
	}
}
