package publisher

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var client *pubsub.Client
var clientMapper = make(map[string]*pubsub.Topic)

func InitConfiguration(ctx context.Context, projectID string, opts ...option.ClientOption) (*pubsub.Client, error) {
	var err error
	client, err = pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		log.Printf("[InitConfiguration] [ERROR] pubsub.NewClient, err: %+v", err.Error())
		return nil, err
	}

	return client, err
}

func PullTopic(ctx context.Context, topicId string) error {
	topic := client.Topic(topicId)
	exist, err := topic.Exists(ctx)
	if exist {
		clientMapper[topicId] = client.Topic(topicId)
		clientMapper[topicId].PublishSettings.NumGoroutines = 3
	} else {
		return fmt.Errorf("not found/ err topic with cfg = %s, %+v", topicId, err)
	}

	return nil
}

func PublishMessage(ctx context.Context, topicId string, rawMessage []byte) error {
	clientTopic, exist := clientMapper[topicId]
	if !exist {
		log.Printf("[PublishMessage] [ERROR] not found client for topic, clientMapper: %+v", clientMapper)
		return fmt.Errorf("not found client for topic %v", topicId)
	}

	message := pubsub.Message{
		Data: rawMessage,
	}
	result := clientTopic.Publish(ctx, &message)
	msgId, err := result.Get(ctx)
	if err != nil {
		log.Printf("Publish message successfully. MsgID = %v | Error = %v", msgId, err)
	}
	return err
}

func ClosePublishMsgOnTopic(topic string) error {
	clientTopic, exist := clientMapper[topic]
	if !exist {
		return fmt.Errorf("not found client for topic %v", topic)
	}
	clientTopic.Stop()
	log.Printf("Stop client on topic %v", topic)
	return nil
}

func CloseAllTopic(ctx context.Context) error {
	for k, v := range clientMapper {
		v.Stop()
		log.Printf("Stop publish msg on topic %v", k)
	}
	return client.Close()
}
