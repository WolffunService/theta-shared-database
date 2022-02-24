package subscriber

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/WolffunGame/theta-shared-database/database/pubsub/mpubsub"
	"google.golang.org/api/option"
)

var subscriberMap = make(map[string]*Subscriber)

func InitializeClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*pubsub.Client, error) {
	return mpubsub.InitializeClient(ctx, projectID, opts...)
}

func Subscribe(ctx context.Context, subId string, fn HandleMsg, opts ...SubscriberOption) error {
	if _, exist := subscriberMap[subId]; exist {
		log.Printf("%v have been existed in keys map", subId)
		return nil
	}

	sub := mpubsub.Client.Subscription(subId)
	if ok, err := sub.Exists(ctx); !ok || err != nil {
		fmt.Println("Cannot subscribe to subscription", subId, ok, err)
		return errors.New("something wrong with subscription")
	}

	var subscriber = &Subscriber{}
	subscriberMap[subId] = subscriber
	ctxChild, cancel := context.WithCancel(ctx)

	subscriber.Subscription = sub
	subscriber.CancelFunc = cancel

	// apply options
	conf := &dynamicConfig{}
	for _, v := range opts {
		v(conf)
	}

	go subscriber.Subscription.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		err := fn(msg.Data)

		if err != nil && !conf.AckSuccessOnly {
			fmt.Println("[pubsub] failed message", subId, err)
			msg.Ack()
		}
	})

	return nil
}

func BlockedSubscribe(ctx context.Context, subId string, fn HandleMsg, opts ...SubscriberOption) error {
	if _, exist := subscriberMap[subId]; exist {
		log.Printf("%v have been existed in keys map", subId)
		return nil
	}

	sub := mpubsub.Client.Subscription(subId)
	if ok, err := sub.Exists(ctx); !ok || err != nil {
		fmt.Println("Cannot subscribe to subscription", subId, ok, err)
		return errors.New("something wrong with subscription")
	}

	var subscriber = &Subscriber{}
	subscriberMap[subId] = subscriber
	ctxChild, cancel := context.WithCancel(ctx)

	subscriber.Subscription = sub
	subscriber.CancelFunc = cancel

	// apply options
	conf := &dynamicConfig{}
	for _, v := range opts {
		v(conf)
	}
	return subscriber.Subscription.Receive(ctxChild, func(ctx context.Context, msg *pubsub.Message) {
		err := fn(msg.Data)

		if err != nil && !conf.AckSuccessOnly {
			fmt.Println("[pubsub] failed message", subId, err)
			msg.Ack()
		}
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
