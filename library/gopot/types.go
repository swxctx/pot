package gopot

// command pot server command
type command struct {
	// cache cmd
	Cmd string `json:"c,omitempty"`
	// cache key
	Key string `json:"k,omitempty"`
	// cache value
	Value string `json:"v,omitempty"`
	// cache key expire
	Expire int64 `json:"ep,omitempty"`
}

// response
type response struct {
	// cache err
	Code int `json:"c"`
	// cache value
	Value string `json:"v"`
	// cache ttl
	TTL int64 `json:"tl"`
	// result exists
	R int64 `json:"r,omitempty"`
}
