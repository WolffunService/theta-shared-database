package subscriber

type dynamicConfig struct {
	AckSuccessOnly bool
	RetrySubscribe bool
}

type SubscriberOption func(*dynamicConfig)

func AckSuccessOnly() SubscriberOption {
	return func(dc *dynamicConfig) {
		dc.AckSuccessOnly = true
	}
}

func RetrySubscribe() SubscriberOption {
	return func(dc *dynamicConfig) {
		dc.RetrySubscribe = true
	}
}
