package pot

import (
	"sync"
	"time"
)

/*
cache
@Description: cache element
*/
type cache struct {
	// cache value
	elems map[string]Element
	mu    sync.RWMutex
}

/*
newCache
@Desc: new cache
@return: *cache
*/
func newCache() *cache {
	elems := make(map[string]Element)
	cache := &cache{
		elems: elems,
	}
	return cache
}

/*
Set
@Desc: set value to cache
@receiver: c
@param: key
@param: value
@param: expiration
*/
func (c *cache) set(key string, value interface{}, expiration ...time.Duration) {
	c.mu.Lock()
	c.elems[key] = Element{
		Value:      value,
		Expiration: paramsCacheExpiration(expiration...),
	}
	c.mu.Unlock()
}

/*
Get
@Desc: get value by cache
@receiver: c
@param: key
@return: interface{}
@return: bool
*/
func (c *cache) get(key string) interface{} {
	c.mu.RLock()
	elem, exists := c.elems[key]
	if !exists {
		c.mu.RUnlock()
		return nil
	}
	// expired
	if elem.expired() {
		c.mu.RUnlock()
		return nil
	}
	c.mu.RUnlock()
	return elem.Value
}

/*
exists
@Desc: check key is exists
@receiver: c
@param: key
@return: bool
*/
func (c *cache) exists(key string) bool {
	c.mu.RLock()
	elem, exists := c.elems[key]
	if !exists {
		c.mu.RUnlock()
		return false
	}

	// expired
	if elem.expired() {
		c.mu.RUnlock()
		return false
	}
	c.mu.RUnlock()
	return true
}

/*
ttl
@Desc: get key ttl
@receiver: c
@param: key
@return: int64
*/
func (c *cache) ttl(key string) int64 {
	c.mu.RLock()

	//expired
	elem, exists := c.elems[key]
	if !exists || elem.expired() {
		c.mu.RUnlock()
		return EXPIRATION_IS_EXPIRED
	}
	//forever
	if elem.Expiration == EXPIRATION_NOT_SET {
		c.mu.RUnlock()
		return EXPIRATION_NOT_SET
	}
	c.mu.RUnlock()
	return elem.Expiration - time.Now().Unix()
}

/*
Expire
@Desc: set key expire
@receiver: c
@param: key
@param: expire
@return: bool
*/
func (c *cache) Expire(key string, expire time.Duration) bool {
	c.mu.RLock()

	elem, exists := c.elems[key]
	if !exists || elem.expired() {
		c.mu.RUnlock()
		return false
	}
	c.mu.RUnlock()
	c.mu.Lock()
	elem.Expiration = paramsCacheExpiration(expire)
	c.elems[key] = elem
	c.mu.Unlock()
	return true
}

/*
delete
@Desc: 删除key
@receiver: c
@param: k
*/
func (c *cache) del(key string) {
	delete(c.elems, key)
}

/*
expired
@Desc: key过期处理
@receiver: c
*/
func (c *cache) expiredRoutine() {
	now := time.Now().Unix()
	c.mu.Lock()
	for k, v := range c.elems {
		if v.Expiration > 0 && now > v.Expiration {
			c.del(k)
		}
	}
	c.mu.Unlock()
}
