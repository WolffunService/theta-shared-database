package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

var clientTopic *pubsub.Topic

func InitConfiguration(config *Config) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	clientTopic = client.Topic(config.TopicID)
	clientTopic.PublishSettings.NumGoroutines = 3
}

func PublishMessage(ctx context.Context, rawMessage []byte) {
	message := pubsub.Message{
		Data: rawMessage,
	}
	result := clientTopic.Publish(ctx, &message)
	log.Print(result.Get(ctx))
}

func GetTopic() *pubsub.Topic {
	return clientTopic
}