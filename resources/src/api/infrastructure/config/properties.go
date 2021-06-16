package config

import (
	"time"
)

type Properties struct {
	In18 struct {
		Language string
	}

	Log struct {
		Level string
	}
}

func GetAsMillisecond(duration time.Duration) time.Duration {
	return duration * time.Millisecond
}
