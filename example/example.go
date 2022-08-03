package main

import (
	"time"

	"github.com/swxctx/pot"
	"github.com/swxctx/xlog"
)

func main() {
	client := pot.NewClient(&pot.Config{
		CleanerInterval: time.Duration(1) * time.Second,
	})
	xlog.Infof("pot init finish")

	key := "pot:test"
	//set
	setCmd := client.Set(key, "1", time.Duration(10)*time.Second)
	if setCmd.GetErr() != nil {
		xlog.Errorf("pot set err-> %v", setCmd.GetErr())
	}
	xlog.Infof("pot set key-> %s, value-> 1, success-> %v", key, setCmd.Success())

	existsCmd := client.Exists(key)
	if existsCmd.GetErr() != nil {
		xlog.Errorf("pot exists err-> %v", existsCmd.GetErr())
	}
	xlog.Infof("pot exists key-> %s, value-> 1, success-> %v, exists-> %v", key, existsCmd.Success(), existsCmd.Result() > 0)

	ttlCmd := client.TTL(key)
	if ttlCmd.GetErr() != nil {
		xlog.Errorf("pot ttl err-> %v", ttlCmd.GetErr())
	}
	xlog.Infof("pot ttl key-> %s, value-> 1, success-> %v, ttl-> %v", key, ttlCmd.Success(), ttlCmd.Result())

	expireCmd := client.Expire(key, time.Duration(5)*time.Second)
	if expireCmd.GetErr() != nil {
		xlog.Errorf("pot expire err-> %v", expireCmd.GetErr())
	}
	xlog.Infof("pot expire key-> %s, value-> 1, success-> %v, current expire-> %v", key, expireCmd.Success(), client.TTL(key).Result())

	time.Sleep(time.Duration(10) * time.Second)

	getCmd := client.Get(key)
	if getCmd.GetErr() != nil {
		xlog.Errorf("pot get err-> %v", getCmd.GetErr())
	}
	xlog.Infof("pot get key-> %s, value-> %v", key, getCmd.String())
}
