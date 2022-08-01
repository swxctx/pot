package pot

import (
	"sync"
	"time"
)

type Client struct {
	Config Config
	Cache  *cache
}

type Config struct {
	DefaultExpiration time.Duration
}

type cache struct {
	elems   map[string]Element
	mu      sync.RWMutex
	cleaner *cleaner
}
