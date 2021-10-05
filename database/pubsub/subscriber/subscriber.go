package subscriber

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var subscriberMap = make(map[string]*Subscriber)

var client *pubsub.Client

func InitializeClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*pubsub.Client, error) {
	var err error
	if client == nil {
		client, err = pubsub.NewClient(ctx, projectID, opts...)
		if err != nil {
			return nil, fmt.Errorf("pubsub.NewClient %v", err)
		}
	}
	return client, nil
}

func Subscribe(ctx context.Context, subId string, fn HandleMsg) error {
	if _, exist := subscriberMap[subId]; exist {
		log.Printf("%v have been existed in keys map", subId)
		return nil
	}

	sub := client.Subscription(subId)
	if ok, err := sub.Exists(ctx); !ok || err != nil {
		fmt.Println("Cannot subscribe to subscription", subId, ok, err)
		return errors.New("something wrong with subscription")
	}

	var subscriber = &Subscriber{}
	subscriberMap[subId] = subscriber
	ctxChild, cancel := context.WithCancel(ctx)

	subscriber.Subscription = sub
	subscriber.CancelFunc = cancel
	go subscriber.Subscription.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		_ = fn(msg.Data)
		// TODO more flow
		msg.Ack()
	})

	return nil
}

func BlockedSubscribe(ctx context.Context, subId string, fn HandleMsg) error {
	if _, exist := subscriberMap[subId]; exist {
		log.Printf("%v have been existed in keys map", subId)
		return nil
	}

	sub := client.Subscription(subId)
	if ok, err := sub.Exists(ctx); !ok || err != nil {
		fmt.Println("Cannot subscribe to subscription", subId, ok, err)
		return errors.New("something wrong with subscription")
	}

	var subscriber = &Subscriber{}
	subscriberMap[subId] = subscriber
	ctxChild, cancel := context.WithCancel(ctx)

	subscriber.Subscription = sub
	subscriber.CancelFunc = cancel
	return subscriber.Subscription.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		_ = fn(msg.Data)
		// TODO more flow
		msg.Ack()
	})
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
