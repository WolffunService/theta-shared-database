package subscriber

type dynamicConfig struct {
	AckSuccessOnly bool
}

type SubscriberOption func(*dynamicConfig)

func AckFailedMessage() SubscriberOption {
	return func(dc *dynamicConfig) {
		dc.AckSuccessOnly = false
	}
}

func NotAckFailedMessage() SubscriberOption {
	return func(dc *dynamicConfig) {
		dc.AckSuccessOnly = true
	}
}
