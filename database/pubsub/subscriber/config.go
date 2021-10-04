package subscriber

import (
	"cloud.google.com/go/pubsub"
	"context"
)

type Config struct {
	ProjectID	string
	TopicID		string
	SubID		string
}

type Subscriber struct {
	cancelFunc	context.CancelFunc
	Topic		*pubsub.Topic
	SubID		*pubsub.Subscription
}