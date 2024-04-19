package pot

/*
StringCmd
@Description: result string cmd
*/
type StringCmd struct {
	baseCmd
}

// newStringCmd
func newStringCmd(key string) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{
			key: key,
		},
	}
}

// Val return original value
func (cmd *StringCmd) Val() interface{} {
	return cmd.val
}

// Result return string value and error
func (cmd *StringCmd) Result() (string, error) {
	return cmd.String(), cmd.err
}

// Int64 return int64 value and error
func (cmd *StringCmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertInt64(cmd.val), nil
}

// Uint64 return Uint64 value and error
func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertUint64(cmd.val), nil
}

// Float64 return Float64 value and error
func (cmd *StringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return valConvertFloat64(cmd.val), nil
}

// String only return String
func (cmd *StringCmd) String() string {
	return valConvertString(cmd.val)
}
