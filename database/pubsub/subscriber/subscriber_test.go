package subscriber

// func TestName(t *testing.T) {
// 	cfg := &Config{
// 		ProjectID: "thetan-staging",
// 		TopicID:   "DEV_ENV_BATTLE_LOGS",
// 		SubID:     "DEV_ENV_BATTLE_LOGS_SUB",
// 	}

// 	// Connect to server
// 	ctx := context.Background()
// 	client, err := InitializeClient(ctx, cfg.ProjectID)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	defer client.Close()

// 	// Subscribe
// 	Subscribe(context.Background(), cfg.SubID, func(s []byte) error {
// 		fmt.Println(s)
// 		return nil
// 	})

// 	// Alive
// 	http.ListenAndServe(":8888", nil)
// }
