package subscriber

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

var subscriberMap = make(map[string]*Subscriber)

func SimpleSubscriber(cfg *Config, fn HandleMsg) error {
	_, exist := subscriberMap[cfg.SubID]
	if exist {
		log.Printf("%v have been existed in keys map", cfg.SubID)
		return nil
	}
	log.Printf("Configuration: %v", *cfg)

	ctx := context.Background()
	client, err:= pubsub.NewClient(ctx, cfg.ProjectID)

	var subscriber = &Subscriber{Client: client}

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

	err = sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		err = fn(string(msg.Data))
		// TODO more flow
		msg.Ack()
	})


	return nil
}

func CloseConnection(subID string) error {
	subscriber, exist := subscriberMap[subID]
	if exist {
		log.Printf("Prepare close connection on topic %v", subID)
		err := subscriber.Client.Close()
		log.Printf("Close connection %v successfully", subID)
		return err
	}

	log.Printf("Not found subscriber with subID %v", subID)
	return nil
}
