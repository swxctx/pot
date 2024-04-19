## Pot CLI

`Pot CLI` 主要用于进行终端命令行的操作，下面介绍命令行工具的使用。

### 安装
可以直接运行 `Pot` 下的 `cli` 目录代码启动命令行终端，也可以直接下载打包好的二进制包在本地运行。

- [直接运行代码](https://github.com/swxctx/pot/tree/main/cli)
- [Linux](https://github.com/swxctx/pot/releases/tag/1.0.0-beta)
- [Mac](https://github.com/swxctx/pot/releases/tag/1.0.0-beta)

### 命令使用

- 查看帮助 `pot-cli -h`

```
(base) ➜  mac git:(main) ✗ pot-cli -h

flag needs an argument: -h
Usage of Pot CLI:
  -h string
    	host address of the server (default "127.0.0.1")
  -p string
    	port of the server (default "9577")
  -version
    	Print version information and exit

Commands:
  set <key> <value> [expiration] - Set the value of a key
  get <key> - Get the value of a key
  del <key> - Delete a key
  exists <key> - Check if a key exists
  ttl <key> - Get the TTL of a key
  expire <key> <seconds> - Set the expiration of a key
```

- 连接 `Pot Server`

```shell
1. 默认连接本地 Pot Server

mac git:(main) ✗ pot-cli

2. 指定 Pot Server IP 及端口

mac git:(main) ✗ pot-cli -h 127.0.0.1 -p 9577

执行命令后输出如下，即表示连接成功：

Welcome to Pot CLI
Connecting to 127.0.0.1:9577
Type 'exit' to quit.
Connecting pot server success.
```

- set

```
格式1：set <key> <value> <expiration>(秒)
格式2：set <key> <value>(永久有效)
127.0.0.1:9577> set test 1 10
true
```

- get

```
格式：get <key>
127.0.0.1:9577> get test
1
```

- del

```
格式：del <key>
127.0.0.1:9577> del test
true
```

- exists

```
格式：exists <key>
127.0.0.1:9577> exists test
-2
```

- ttl

```
格式：ttl <key>
127.0.0.1:9577> ttl test
988
```

- expire

```
格式：expire <key> <expiration>(秒)
127.0.0.1:9577> expire test 100
true
```