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
	//set
	if err := client.Set("pot:test", "1"); err != nil {
		xlog.Errorf("pot set err-> %v", err)
	}
	xlog.Infof("pot set key-> pot:test, value-> 1")

	getCmd := client.Get("pot:test")
	if getCmd.GetErr() != nil {
		xlog.Errorf("pot get err-> %v", getCmd.GetErr())
	}
	xlog.Infof("pot get key-> pot:test, value-> %v", getCmd.String())
}
