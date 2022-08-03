package pot

type baseCmd struct {
	key string
	err error
	val interface{}
}

func (c *baseCmd) GetErr() error {
	return c.err
}

func (c *baseCmd) SetErr(err error) {
	c.err = err
}

func (c *baseCmd) GetVal() interface{} {
	return c.val
}

func (c *baseCmd) SetVal(val interface{}) {
	c.val = val
}

/*
  StringCmd
  @Description: result string cmd
*/
type StringCmd struct {
	baseCmd
}

/*
  NewStringCmd
  @Desc:
  @param: args
  @return: *StringCmd
*/
func NewStringCmd(key string) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{
			key: key,
		},
	}
}

func (cmd *StringCmd) Val() interface{} {
	return cmd.val
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.String(), cmd.err
}

func (cmd *StringCmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertInt64(cmd.val), nil
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertUint64(cmd.val), nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertFloat64(cmd.val), nil
}

func (cmd *StringCmd) String() string {
	return valConvertString(cmd.val)
}
