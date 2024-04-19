package pot

import (
	"fmt"
	"time"
)

/*
checkClientCache
@Desc: check client
@receiver: p
@return: error
*/
func (p *Pot) checkClientCache() error {
	if p.cache == nil || p.cache.elems == nil {
		return fmt.Errorf("pot: client cache is nil, chech cache init")
	}
	return nil
}

/*
Set
@Desc: set value to cache
@receiver: p
@param: key
@param: value
@param: expiration
@return: error
*/
func (p *Pot) Set(key string, value interface{}, expiration ...time.Duration) *StatusCmd {
	cmd := newStatusCmd(key)
	if err := p.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	p.cache.set(key, value, expiration...)
	cmd.setSuccess(true)
	return cmd
}

/*
Get
@Desc: get value by cache
@receiver: p
@param: key
@return: interface{}
@return: error
*/
func (p *Pot) Get(key string) *StringCmd {
	cmd := newStringCmd(key)
	if err := p.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	cmd.setVal(p.cache.get(key))
	return cmd
}

/*
Del
@Desc:
@receiver: p
@param: key
*/
func (p *Pot) Del(key string) *StatusCmd {
	cmd := newStatusCmd(key)
	if err := p.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	p.cache.del(key)
	cmd.setSuccess(true)
	return cmd
}

/*
Exists
@Desc: check key exists
@receiver: p
@param: key
@return: bool
*/
func (p *Pot) Exists(key string) *StatusCmd {
	cmd := newStatusCmd(key)
	if err := p.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	if p.cache.exists(key) {
		cmd.setResult(POT_ACTION_RESULT_EXISTS)
	}
	cmd.setSuccess(true)
	return cmd
}

/*
TTL
@Desc: check key expire ttl
@receiver: p
@param: key
@return: int64
*/
func (p *Pot) TTL(key string) *StatusCmd {
	cmd := newStatusCmd(key)
	if err := p.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	cmd.setSuccess(true)
	cmd.setResult(p.cache.ttl(key))
	return cmd
}

/*
Expire
@Desc: set key expire
@receiver: c
@param: key
@param: expire
*/
func (c *Pot) Expire(key string, expire time.Duration) *StatusCmd {
	cmd := newStatusCmd(key)
	if err := c.checkClientCache(); err != nil {
		cmd.setErr(err)
		return cmd
	}
	c.cache.expire(key, expire)
	cmd.setSuccess(true)
	return cmd
}
