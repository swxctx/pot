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
func (c *Client) Set(key string, value interface{}, expiration ...time.Duration) *StatusCmd {
	cmd := NewStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	c.cache.set(key, value, expiration...)
	cmd.setSuccess(true)
	return cmd
}

/*
Get
@Desc: get value by cache
@receiver: c
@param: key
@return: interface{}
@return: error
*/
func (c *Client) Get(key string) *StringCmd {
	cmd := NewStringCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	cmd.setVal(c.cache.get(key))
	return cmd
}

/*
Del
@Desc:
@receiver: c
@param: key
*/
func (c *Client) Del(key string) *StatusCmd {
	cmd := NewStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	c.cache.del(key)
	cmd.setSuccess(true)
	return cmd
}

/*
Exists
@Desc: check key exists
@receiver: c
@param: key
@return: bool
*/
func (c *Client) Exists(key string) *StatusCmd {
	cmd := NewStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	if c.cache.exists(key) {
		cmd.setResult(POT_ACTION_RESULT_EXISTS)
	}
	cmd.setSuccess(true)
	return cmd
}

/*
TTL
@Desc: check key expire ttl
@receiver: c
@param: key
@return: int64
*/
func (c *Client) TTL(key string) *StatusCmd {
	cmd := NewStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	cmd.setSuccess(true)
	cmd.setResult(c.cache.ttl(key))
	return cmd
}

/*
Expire
@Desc: set key expire
@receiver: c
@param: key
@param: expire
*/
func (c *Client) Expire(key string, expire time.Duration) *StatusCmd {
	cmd := NewStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	c.cache.expire(key, expire)
	cmd.setSuccess(true)
	return cmd
}
