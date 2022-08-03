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
	"github.com/swxctx/pot"
	"github.com/swxctx/xlog"
	"time"
)

func main() {
	client := pot.NewClient(&pot.Config{
		CleanerInterval: time.Duration(1) * time.Second,
	})
	xlog.Infof("pot init finish")
	//set
	if err := client.Set("pot:test", 1); err != nil {
		xlog.Errorf("pot set err-> %v", err)
	}
	xlog.Infof("pot set key-> pot:test, value-> 1")

	value, err := client.Get("pot:test")
	if err != nil {
		xlog.Errorf("pot get err-> %v", err)
	}
	xlog.Infof("pot get key-> pot:test, value-> %v", value)
}
```