package main

import (
	"time"

	"github.com/swxctx/pot"
	"github.com/swxctx/pot/plog"
)

var (
	// pot cache client
	potClient *pot.Pot
)

func main() {
	plog.Infof("pot server start...")
	plog.Infof("pot server reload config...")

	// init config
	if err := reloadConfig(); err != nil {
		plog.Errorf("%v", err)
		panic(err)
	}

	plog.Infof("pot server config loaded...")
	plog.Infof("pot cache client init start...")

	// start pot cache
	potClient = pot.Init(&pot.Config{
		CleanerInterval: time.Duration(cfg.CleanerInterval) * time.Millisecond,
		ShowTrace:       cfg.ShowTrace,
	})
	if potClient == nil {
		plog.Errorf("pot cache client init err, client is nil...")
		return
	}
	plog.Infof("pot cache client init finished...")

	plog.Infof("listen port-> %s", cfg.Port)
	plog.Infof("start tcp listen...")

	// new pot server
	if err := newServer(); err != nil {
		plog.Errorf("%v", err)
		panic(err)
	}
}
