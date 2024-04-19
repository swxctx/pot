package pot

import "time"

/*
Config
@Description: pot config
*/
type Config struct {
	// cache key clean interval(default: 1s)
	CleanerInterval time.Duration
	// show trace level log
	ShowTrace bool
}
