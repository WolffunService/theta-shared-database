package subscriber

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Config struct {
	ProjectID string
	TopicID   string
	SubID     string
}

type Subscriber struct {
	CancelFunc   context.CancelFunc
	Subscription *pubsub.Subscription
}
