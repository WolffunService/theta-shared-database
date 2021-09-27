package publisher

import (
	"context"
	"testing"
)

func TestPublishMessage(t *testing.T) {
	config := Config{
		projectID: "thetan-staging",
		topicID : "BATTLE_LOGS_TOPIC",
	}
	InitConfiguration(config)
	message := "I'm handsome"
	PublishMessage(context.TODO(), []byte(message))
}
