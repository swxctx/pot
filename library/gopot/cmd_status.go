package gopot

/*
StatusCmd
@Description: action status
*/
type StatusCmd struct {
	baseCmd
	// 执行操作成功
	success bool
	// 执行结果
	result int64
}

// newStatusCmd
func newStatusCmd(key string) *StatusCmd {
	return &StatusCmd{
		baseCmd: baseCmd{
			key: key,
		},
	}
}

func (cmd *StatusCmd) setSuccess(success bool) {
	cmd.success = success
}

func (cmd *StatusCmd) setResult(result int64) {
	cmd.result = result
}

// Success set is success?
func (cmd *StatusCmd) Success() bool {
	return cmd.success
}

// Result set result 1 OR 0
func (cmd *StatusCmd) Result() int64 {
	return cmd.result
}
