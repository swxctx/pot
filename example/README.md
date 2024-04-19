### Install

```go
go get github.com/swxctx/pot
```

### Example
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