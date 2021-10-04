package subscriber

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

var subscriberMap = make(map[string]*Subscriber)

var client *pubsub.Client

func InitializeClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	var err error
	if client == nil {
		client, err = pubsub.NewClient(ctx, projectID)
		if err != nil {
			return nil, fmt.Errorf("pubsub.NewClient %v", err)
		}
	}
	return client, nil
}

func Subscribe(ctx context.Context, cfg *Config, fn HandleMsg) error {
	if _, exist := subscriberMap[cfg.SubID]; exist {
		log.Printf("%v have been existed in keys map", cfg.SubID)
		return nil
	}

	log.Printf("Configuration: %v", *cfg)

	// Check topic
	topic := client.Topic(cfg.TopicID)
	isExist, err := topic.Exists(ctx)
	if !isExist || err != nil {
		return fmt.Errorf("pubsub.NewClient: exist = %v | %v", isExist, err)
	}

	// Add susbcriber
	var subscriber = &Subscriber{}
	subscriber.Topic = topic
	subscriberMap[cfg.SubID] = subscriber

	sub := client.Subscription(cfg.SubID)

	subscriber.Subscription = sub
	ctxChild, cancel := context.WithCancel(ctx)
	subscriber.CancelFunc = cancel
	err = sub.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		err = fn(msg.Data)
		// TODO more flow
		msg.Ack()
	})

	return nil
}

func CloseConnection(subID string) {
	subscriber, exist := subscriberMap[subID]
	if exist {
		subscriber.CancelFunc()
		delete(subscriberMap, subID)
		log.Printf("Close connection %v successfully\n", subID)
	} else {
		log.Printf("Not found subscriber with subID %v\n", subID)
	}
}
