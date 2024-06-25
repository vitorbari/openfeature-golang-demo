package main

import (
	"github.com/open-feature/go-sdk/openfeature"

	flagsmithClient "github.com/Flagsmith/flagsmith-go-client/v3"
	flagsmith "github.com/open-feature/go-sdk-contrib/providers/flagsmith/pkg"
	statsig "github.com/open-feature/go-sdk-contrib/providers/statsig/pkg"
)

func buildStatsigProvider(cfg config) openfeature.FeatureProvider {
	provider, err := statsig.NewProvider(
		statsig.ProviderConfig{
			SdkKey: cfg.statsigKey,
		},
	)
	if err != nil {
		panic(err)
	}

	provider.Init(openfeature.EvaluationContext{})

	return provider
}

func buildFlagSmithProvider(cfg config) openfeature.FeatureProvider {
	client := flagsmithClient.NewClient(cfg.flaghsmithApiKey)
	return flagsmith.NewProvider(client, flagsmith.WithUsingBooleanConfigValue())
}
