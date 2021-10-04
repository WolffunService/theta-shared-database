package subscriber

import "cloud.google.com/go/pubsub"

type Config struct {
	ProjectID	string
	TopicID		string
	SubID		string
}

type Subscriber struct {
	Client 	*pubsub.Client
	Topic	*pubsub.Topic
	SubID	*pubsub.Subscription
}