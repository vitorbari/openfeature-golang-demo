package main

import (
	"context"
	"time"

	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	ctx := context.Background()

	cfg := loadConfigFromEnvVars()

	// toggle between providers
	openfeature.SetProvider(buildStatsigProvider(cfg))
	//openfeature.SetProvider(buildFlagSmithProvider(cfg))

	// Create a new client
	client := openfeature.NewClient("app")

	for {
		err := handleMyFeature(ctx, client)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
