package subscriber

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	cfg := &Config{
		ProjectID: "thetan-staging",
		TopicID:   "DEV_ENV_BATTLE_LOGS",
		SubID:     "DEV_ENV_BATTLE_LOGS_SUB",
	}
	SimpleSubscriber(cfg, func(s string) error {
		fmt.Println(s)
		return nil
	})

	time.Sleep(5_000)

	CloseConnection(cfg.SubID)

	http.ListenAndServe(":8888", nil)
}