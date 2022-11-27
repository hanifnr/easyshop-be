package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}
	return mustGetenv(key)
}
