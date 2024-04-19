# pot
- 基于Golang实现的 k-v存储，目前规划为轻量、便捷、小容量，用于在简单业务中进行一些临时数据、轻量逻辑的实现。
- 接下来会逐步进行完善，实现完善的k-v存储DB

# 功能规划
- [x] 基础架构
- [x] 基础存储功能
- [x] key有效期处理
- [x] 返回数据类型处理
- [ ] 支持基础数据、有序集合、无序集合
- [ ] 支持备份到磁盘
- [ ] 支持重启数据恢复
- [ ] 支持HTTP接口操作
- [ ] 支持TCP协议支持
- [ ] 支持多机分布式部署
- [ ] 待定

# Install

```go
go get github.com/swxctx/pot
```

# Example
```go
package main

import (
	"time"

	"github.com/swxctx/pot/plog"

	"github.com/swxctx/pot"
)

func main() {
	client := pot.Init(&pot.Config{
		CleanerInterval: time.Duration(1) * time.Second,
		ShowTrace:       true,
	})
	plog.Infof("pot init finish")

	key := "pot:test"
	//set
	setCmd := client.Set(key, "1", time.Duration(10)*time.Second)
	if setCmd.Err() != nil {
		plog.Errorf("pot set err-> %v", setCmd.Err())
	}
	plog.Infof("pot set key-> %s, value-> 1, success-> %v", key, setCmd.Success())

	existsCmd := client.Exists(key)
	if existsCmd.Err() != nil {
		plog.Errorf("pot exists err-> %v", existsCmd.Err())
	}
	plog.Infof("pot exists key-> %s, value-> 1, success-> %v, exists-> %v", key, existsCmd.Success(), existsCmd.Result() > 0)

	ttlCmd := client.TTL(key)
	if ttlCmd.Err() != nil {
		plog.Errorf("pot ttl err-> %v", ttlCmd.Err())
	}
	plog.Infof("pot ttl key-> %s, value-> 1, success-> %v, ttl-> %v", key, ttlCmd.Success(), ttlCmd.Result())

	expireCmd := client.Expire(key, time.Duration(5)*time.Second)
	if expireCmd.Err() != nil {
		plog.Errorf("pot expire err-> %v", expireCmd.Err())
	}
	plog.Infof("pot expire key-> %s, value-> 1, success-> %v, current expire-> %v", key, expireCmd.Success(), client.TTL(key).Result())

	time.Sleep(time.Duration(10) * time.Second)

	getCmd := client.Get(key)
	if getCmd.Err() != nil {
		plog.Errorf("pot get err-> %v", getCmd.Err())
	}
	plog.Infof("pot get key-> %s, value-> %v", key, getCmd.String())

	select {}
}
```