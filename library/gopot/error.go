package gopot

import "fmt"

const (
	// code_cmd_action_normal 正常执行
	code_cmd_action_normal = 0

	// code_err_for_parse_internal 内部错误
	code_err_for_parse_internal = 100

	// code_err_for_parse_cmd 命令格式解析错误
	code_err_for_parse_cmd = 101

	// code_err_for_cmd_valid 命令非法
	code_err_for_cmd_valid = 102

	// code_err_for_cmd_action_failed 操作不成功
	code_err_for_cmd_action_failed = 103
)

// getErrorForAction
func getErrorForAction(code int, key string) error {
	var (
		err error
	)
	switch code {
	case code_err_for_parse_internal:
		err = fmt.Errorf("pot server internal, key-> %s", key)
		break
	case code_err_for_parse_cmd:
		err = fmt.Errorf("pot server cmd err, interval, key-> %s", key)
		break
	case code_err_for_cmd_valid:
		err = fmt.Errorf("pot server cmd not exists, key-> %s", key)
		break
	case code_err_for_cmd_action_failed:
		err = fmt.Errorf("pot server action failed, key-> %s", key)
		break
	}
	return err
}
