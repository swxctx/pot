## Pot
该项目为完整的可部署远程调用的 k-v 缓存工具，目前支持单机部署，多业务远程调用，架构与Redis、Memcache类似，会逐渐进行功能完善，包括缓存落地存储、数据恢复、热加载等。

目前实现的命令如下

- set
- get
- ttl
- expire
- del
- exists

[Pot Server通信协议](https://github.com/swxctx/pot/blob/main/server/protocol.md)

## 安装
`Pot`提供源码，所以比较灵活，可以直接编译启动即可，同时也提供了打包好的安装包，直接下载启动即可。

***支持的安装方式如下***

- 直接运行代码(不推荐)
- 源码编译安装
- 直接下载安装包启动

### 直接运行源码启动

- 下载代码

```
git clone https://github.com/swxctx/pot.git
```

- 进入`server`目录

```
cd pot/server
```

- 运行`Pot`服务

```
go run *.go
```

- 运行效果

	运行后出现如下效果，即表示`Pot`服务已经启动了，监听端口为`9577`

	![](./doc/code_run.jpg)

### 源码编译安装

- 下载代码

```
git clone https://github.com/swxctx/pot.git
```

- 进入`server`目录

```
cd pot/server
```

- 编译

```
go build
```

- 运行服务

```
./server
```

- 运行效果

	运行后出现如下效果，即表示`Pot`服务已经启动了，监听端口为`9577`

	![](./doc/build_run.jpg)

### 直接下载安装包启动

目前简单提供了`Mac`版本、`Linux`版本，下载到本地后直接运行`./pot-server`即可启动，下载链接如下：

- [Mac](https://github.com/swxctx/pot/releases/tag/1.0.0-beta)
- [Linux amd64](https://github.com/swxctx/pot/releases/tag/1.0.0-beta)

### 配置修改

打开`config/config.json`文件，会有如下内容：

```
{
    "port": "9577", // 监听端口号
    "cleaner_interval": 1000, // 过期key清除间隔
    "show_trace": true // 是否展示详细的调用信息
}
```