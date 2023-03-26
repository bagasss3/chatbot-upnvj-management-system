package config

import "time"

func parseTimeDuration(t string, defaultt time.Duration) time.Duration {
	timeDurr, err := time.ParseDuration(t)
	if err != nil {
		return defaultt
	}
	return timeDurr
}
