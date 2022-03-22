package subscriber

import "cloud.google.com/go/pubsub"

type HandleMsg func([]byte) error
type HandleMessage func(*pubsub.Message) error
