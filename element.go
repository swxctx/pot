package pot

import (
	"time"
)

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
