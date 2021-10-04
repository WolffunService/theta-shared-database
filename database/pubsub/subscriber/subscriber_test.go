package subscriber

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	cfg := &Config{
		ProjectID: "thetan-staging",
		TopicID:   "DEV_ENV_BATTLE_LOGS",
		SubID:     "DEV_ENV_BATTLE_LOGS_SUB",
	}
	SimpleSubscriber(context.Background(), cfg, func(s string) error {
		fmt.Println(s)
		return nil
	})

	http.ListenAndServe(":8888", nil)
}