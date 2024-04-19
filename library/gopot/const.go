package gopot

const (
	// cmd_set set cache cmd
	cmd_set = "s"

	// cmd_get get cache cmd
	cmd_get = "g"

	// cmd_del del cache cmd
	cmd_del = "d"

	// cmd_exists check cache cmd
	cmd_exists = "ex"

	// cmd_ttl get ttl cmd
	cmd_ttl = "tl"

	// cmd_expire set expire cmd
	cmd_expire = "ep"
)

const (
	// EXPIRATION_NOT_SET 未设置过期时间
	EXPIRATION_NOT_SET = -1

	// EXPIRATION_IS_EXPIRED key不存在||已过期
	EXPIRATION_IS_EXPIRED = -2
)

const (
	// POT_ACTION_RESULT_EXISTS 操作有数据/成功
	POT_ACTION_RESULT_EXISTS = 1
)
