package publisher

import (
	"context"
	"testing"
)

func TestPublishMessage(t *testing.T) {
	config := Config{
		ProjectID: "thetan-staging",
		TopicID : "test-topic",
	}
	InitConfiguration(&config)
	message := "I'm handsome"
	PublishMessage(context.TODO(), config.TopicID, []byte(message))
	CloseAllTopic(context.TODO())
}
