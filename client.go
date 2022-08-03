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
func (c *Client) Get(key string) *StringCmd {
	cmd := NewStringCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.SetErr(err)
		return cmd
	}
	cmd.SetVal(c.cache.get(key))
	return cmd
}

/*
  Del
  @Desc:
  @receiver: c
  @param: key
*/
func (c *Client) Del(key string) {
	if err := c.checkClientCache(); err != nil {
		return
	}
	c.cache.del(key)
}

/*
  Exists
  @Desc: check key exists
  @receiver: c
  @param: key
  @return: bool
*/
func (c *Client) Exists(key string) bool {
	if err := c.checkClientCache(); err != nil {
		return false
	}
	return c.cache.exists(key)
}

/*
  TTL
  @Desc: chech key expire ttl
  @receiver: c
  @param: key
  @return: int64
*/
func (c *Client) TTL(key string) int64 {
	if err := c.checkClientCache(); err != nil {
		return 0
	}
	return c.cache.ttl(key)
}

/*
  Expire
  @Desc: set key expire
  @receiver: c
  @param: key
  @param: expire
*/
func (c *Client) Expire(key string, expire time.Duration) {
	if err := c.checkClientCache(); err != nil {
		return
	}
	c.cache.Expire(key, expire)
}
