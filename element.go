package pot

import (
	"time"
)

/*
Element
@Description: value
*/
type Element struct {
	// cache value
	Value interface{}
	// cache expire in
	Expiration int64
}

/*
expired
@Desc: check element is expired
@param: elem
@return: bool
*/
func (e *Element) expired() bool {
	if e.Expiration > 0 && time.Now().Unix() < e.Expiration {
		return false
	}
	if e.Expiration == EXPIRATION_NOT_SET {
		return false
	}
	return true
}
