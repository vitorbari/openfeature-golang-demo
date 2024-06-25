package main

import (
	"context"
	"fmt"

	"github.com/open-feature/go-sdk/openfeature"
)

const (
	featureKey = "my_awesome_feature"
)

func handleMyFeature(ctx context.Context, feature openfeature.IClient) error {
	evaluationCtx := openfeature.NewEvaluationContext("", map[string]interface{}{
		"UserID": "11111",
	})

	isEnabled, err := feature.BooleanValue(ctx, featureKey, false, evaluationCtx)
	if err != nil {
		return fmt.Errorf("failed to evaluate feature: %w", err)
	}

	if isEnabled {
		fmt.Println("Feature enabled ✅ ")
		return nil
	}

	fmt.Println("Feature disabled ❌ ")
	return nil
}
