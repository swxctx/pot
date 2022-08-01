package pot

import "time"

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
func (c *cache) Set(key string, value interface{}, expiration ...time.Duration) {
	c.mu.Lock()
	c.elems[key] = Element{
		Value:      value,
		Expiration: paramsCacheExpiration(expiration...),
	}
	c.mu.Lock()
}

/*
  Get
  @Desc: get value by cache
  @receiver: c
  @param: key
  @return: interface{}
  @return: bool
*/
func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	elem, exists := c.elems[key]
	if !exists {
		c.mu.RUnlock()
		return nil, false
	}
	// expired
	if elem.Expiration > 0 && time.Now().UnixNano() < elem.Expiration {
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()
	return elem, false
}
