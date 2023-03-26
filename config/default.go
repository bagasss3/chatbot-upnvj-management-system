package config

import "time"

// default const
const (
	DefaultConnMaxLifeTime time.Duration = 1 * time.Hour
	DefaultConnMaxIdleTime time.Duration = 15 * time.Minute
)
