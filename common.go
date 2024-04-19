package pot

import (
	"time"
)

// paramsCacheExpiration
func paramsCacheExpiration(expiation ...time.Duration) int64 {
	// 未设置，默认永久有效
	if len(expiation) <= 0 || expiation[0].Seconds() == 0 {
		return -1
	}
	return time.Now().Add(expiation[0]).Unix()
}
