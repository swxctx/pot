package pot

// baseCmd
type baseCmd struct {
	// cache key
	key string
	// cache error
	err error
	// cache value
	val interface{}
}

func (c *baseCmd) setVal(val interface{}) {
	c.val = val
}

func (c *baseCmd) setErr(err error) {
	c.err = err
}

func (c *baseCmd) Err() error {
	return c.err
}

func (c *baseCmd) Val() interface{} {
	return c.val
}
