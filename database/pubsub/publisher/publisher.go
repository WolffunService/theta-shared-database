package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

var client *pubsub.Client
var clientMapper = make(map[string]*pubsub.Topic)

func InitConfiguration(config *Config) error {
	ctx := context.Background()
	if client == nil {
		var err error
		client, err = pubsub.NewClient(ctx, config.ProjectID)
		if err != nil {
			log.Fatalf("pubsub.NewClient: %v", err)
		}
	}
	topic := client.Topic(config.TopicID)
	exist, err := topic.Exists(ctx)
	if exist {
		clientMapper[config.TopicID] = client.Topic(config.TopicID)
		clientMapper[config.TopicID].PublishSettings.NumGoroutines = 3
	} else {
		return fmt.Errorf("not found topic with cfg = %v", config)
	}
	return err
}

func PublishMessage(ctx context.Context, topic string, rawMessage []byte) error {

	clientTopic, exist := clientMapper[topic]
	if !exist {
		return fmt.Errorf("not found client for topic %v", topic)
	}
	message := pubsub.Message{
		Data: rawMessage,
	}
	result := clientTopic.Publish(ctx, &message)
	msgId, err := result.Get(ctx)
	log.Printf("Publish message successfully. MsgID = %v | Error = %v", msgId, err)
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