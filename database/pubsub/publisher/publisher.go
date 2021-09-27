package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

var clientTopic *pubsub.Topic

func InitConfiguration(config Config) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.projectID)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	clientTopic = client.Topic(config.topicID)
	clientTopic.PublishSettings.NumGoroutines = 3
}

func PublishMessage(ctx context.Context, rawMessage []byte) {
	message := pubsub.Message{
		Data: rawMessage,
	}
	clientTopic.Publish(ctx, &message)
	log.Print("Publish message successfully")
}
