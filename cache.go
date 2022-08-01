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
  Element
  @Description: value 基本信息
*/
type Element struct {
	//数据
	Value interface{}
	//过期时间
	Expiration int64
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
	if elem.Expiration > 0 && time.Now().UnixNano() < elem.Expiration {
		c.mu.RUnlock()
		return nil
	}
	c.mu.RUnlock()
	return elem.Value
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
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.elems {
		if v.Expiration > 0 && now > v.Expiration {
			c.del(k)
		}
	}
	c.mu.Unlock()
}
