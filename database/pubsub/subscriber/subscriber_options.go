package subscriber

type dynamicConfig struct {
	AckSuccessOnly bool
}

type SubscriberOption func(*dynamicConfig)

func AckSuccessOnly() SubscriberOption {
	return func(dc *dynamicConfig) {
		dc.AckSuccessOnly = true
	}
}
