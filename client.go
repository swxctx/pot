package pot

import (
	"fmt"
	"time"
)

/*
  checkClientCache
  @Desc: check client
  @receiver: c
  @return: error
*/
func (c *Client) checkClientCache() error {
	if c.cache == nil || c.cache.elems == nil {
		return fmt.Errorf("pot: client cache is nil, chech cache init")
	}
	return nil
}

/*
  Set
  @Desc: set value to cache
  @receiver: c
  @param: key
  @param: value
  @param: expiration
  @return: error
*/
func (c *Client) Set(key string, value interface{}, expiration ...time.Duration) error {
	if err := c.checkClientCache(); err != nil {
		return err
	}
	c.cache.set(key, value, expiration...)
	return nil
}

/*
  Get
  @Desc: get vavlue by cache
  @receiver: c
  @param: key
  @return: interface{}
  @return: error
*/
func (c *Client) Get(key string) (interface{}, error) {
	if err := c.checkClientCache(); err != nil {
		return nil, err
	}
	return c.cache.get(key), nil
}
