package pot

import (
	"time"
)

/*
Client
@Description: pot client
*/
type Client struct {
	// config
	Config *Config
	// cache
	cache *cache
	// cleaner
	cleaner *cleaner
}

/*
Config
@Description: pot config
*/
type Config struct {
	CleanerInterval   time.Duration
	DefaultExpiration time.Duration
}

/*
NewClient
@Desc:
@param: cfg
*/
func NewClient(cfg *Config) *Client {
	client := &Client{
		Config:  cfg,
		cache:   newCache(),
		cleaner: newCleaner(cfg.CleanerInterval),
	}
	// start cleaner
	client.startCleaner()
	return client
}

/*
getConfig
@Desc: pot get config
@receiver: c
@return: *Config
*/
func (c *Client) getConfig() *Config {
	if c.Config != nil {
		return c.Config
	}
	c.Config = &Config{
		CleanerInterval:   1,
		DefaultExpiration: -1,
	}
	return c.Config
}

/*
getCache
@Desc: 获取cache
@receiver: c
@return: *cache
*/
func (c *Client) getCache() *cache {
	if c.cache != nil {
		return c.cache
	}
	c.cache = newCache()
	return c.cache
}
