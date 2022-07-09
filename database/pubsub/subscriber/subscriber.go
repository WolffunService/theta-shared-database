package subscriber

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/WolffunGame/theta-shared-database/database/mredis"
	"github.com/WolffunGame/theta-shared-database/database/mredis/thetanlock"
	"github.com/WolffunGame/theta-shared-database/database/pubsub/mpubsub"
	"google.golang.org/api/option"
)

var subscriberMap = make(map[string]*Subscriber)

func InitConfiguration(ctx context.Context, projectID string, opts ...option.ClientOption) (*pubsub.Client, error) {
	return mpubsub.InitializeClient(ctx, projectID, opts...)
}

// Subscribe to the subscription
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

	go simpleSubcribe(ctxChild, subscriber, *conf, fn, subId)

	return nil
}

// BlockedSubscribe will block the main thread (or goroutine), retry option will be ignored
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

	return simpleSubcribe(ctxChild, subscriber, *conf, fn, subId)
}

// simpleSubscribe just subscribe to an subscription
func simpleSubcribe(ctx context.Context, subscriber *Subscriber, cfg dynamicConfig, fn HandleMsg, subId string) error {
	var err error = nil
	for ok := true; ok; ok = cfg.RetrySubscribe {
		err = subscriber.Subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			key := fmt.Sprintf("%s%s", subId, msg.ID)
			if exists, err := mredis.Exists(ctx, key); err != nil {
				msg.Nack()
				fmt.Println("[pubsub] failed redis", err)
				return
			} else if exists {
				msg.Ack()
				return
			}

			err := fn(msg.Data)

			if err != nil {
				if !cfg.AckSuccessOnly {
					fmt.Println("[pubsub] failed message", subId, err)
					thetanlock.LockTimeoutDur(key, 24*3*time.Hour)
					msg.Ack()
				} else {
					msg.Nack()
				}
			} else {
				thetanlock.LockTimeoutDur(key, 24*3*time.Hour)
				msg.Ack()
			}
		})
	}

	return err
}

func SubscribeV2(ctx context.Context, subId string, fn HandleMessage, opts ...SubscriberOption) error {
	if _, exist := subscriberMap[subId]; exist {
		log.Printf("%v have been existed in keys map", subId)
		return nil
	}

	sub := mpubsub.Client.Subscription(subId)

	//Skip vì đã kiểm tra trước khi gọi hàm này rồi

	//if ok, err := sub.Exists(ctx); !ok || err != nil {
	//	fmt.Println("Cannot subscribe to subscription", subId, ok, err)
	//	return errors.New("something wrong with subscription")
	//}

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
		key := fmt.Sprintf("%s%s", subId, msg.ID)

		if exists, err := mredis.Exists(ctxChild, key); err != nil {
			msg.Nack()
			fmt.Println("[pubsub] failed redis", err)
			return
		} else if exists {
			msg.Ack()
			return
		}

		err := fn(msg)

		if err != nil {
			if !conf.AckSuccessOnly {
				fmt.Println("[pubsub] failed message", subId, err)
				thetanlock.LockTimeoutDur(key, 24*3*time.Hour)
				msg.Ack()
			} else {
				msg.Nack()
			}
		} else {
			thetanlock.LockTimeoutDur(key, 24*3*time.Hour)
			msg.Ack()
		}
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
