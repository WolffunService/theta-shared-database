package subscriber

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

var subscriberMap = make(map[string]*Subscriber)

var client *pubsub.Client


func SimpleSubscriber(ctx context.Context, cfg *Config, fn HandleMsg) error {
	var err error

	err = initializeClient(ctx, cfg.ProjectID)

	if err != nil {return err}

	_, exist := subscriberMap[cfg.SubID]
	if exist {
		log.Printf("%v have been existed in keys map", cfg.SubID)
		return nil
	}
	log.Printf("Configuration: %v", *cfg)

	var subscriber = &Subscriber{}

	// add subscriber to map
	subscriberMap[cfg.SubID] = subscriber

	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	topic := client.Topic(cfg.TopicID)

	subscriber.Topic = topic

	isExist, err := topic.Exists(ctx)

	if !isExist || err != nil {
		return fmt.Errorf("pubsub.NewClient: exist = %v | %v", isExist, err)
	}

	sub := client.Subscription(cfg.SubID)

	subscriber.SubID = sub
	ctxChild, cancel := context.WithCancel(ctx)
	subscriber.cancelFunc = cancel
	err = sub.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		err = fn(string(msg.Data))
		// TODO more flow
		msg.Ack()
	})


	return nil
}

func CloseConnection(subID string) {
	subscriber, exist := subscriberMap[subID]
	if exist {
		log.Printf("Prepare close connection on topic %v", subID)
		subscriber.cancelFunc()
		log.Printf("Close connection %v successfully", subID)
	}
	log.Printf("Not found subscriber with subID %v", subID)
}

func initializeClient(ctx context.Context, projectID string) error {
	var err error
	if client == nil {
		client, err = pubsub.NewClient(ctx, projectID)
		if err != nil {
			return fmt.Errorf("pubsub.NewClient %v", err)
		}
	}
	return err
}
