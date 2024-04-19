## Pot Server 协议说明

`Pot Server`采用`TCP`进行通信，下面将主要描述发送包与响应包格式。

### 发送包

- 结构

```
{
	"c": 	"set",
	"key": 	"test",
	"value": "123",
	"expire": 10
}

c: 操作的命令
key: 缓存key
value: 缓存的内容
expire: 缓存有效期，单位为秒
```

- `CMD命令`

```
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
```

### 接收包

- 结构

```
{
	"c": 0,
	"v": "123",
	"tl": 10,
	"r": 1
}

c: 操作响应code
v: get cache的值
tl: TTL命令时返回的值
r: exists命令时返回的值[1: 存在]
```

- `Code 值`

```
// code_cmd_action_normal 正常执行
code_cmd_action_normal = 0

// code_err_for_parse_internal 内部错误
code_err_for_parse_internal = 100

// code_err_for_parse_cmd 命令格式不正确
code_err_for_parse_cmd = 101

// code_err_for_cmd_valid 命令非法
code_err_for_cmd_valid = 102

// code_err_for_cmd_action_failed 操作不成功
code_err_for_cmd_action_failed = 103
```