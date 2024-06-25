package main

import "os"

type config struct {
	statsigKey       string
	flaghsmithApiKey string
}

func loadConfigFromEnvVars() config {
	return config{
		statsigKey:       os.Getenv("STATSIG_KEY"),
		flaghsmithApiKey: os.Getenv("FLAGSMITH_KEY"),
	}
}
