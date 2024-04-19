package pot

import (
	"github.com/swxctx/pot/plog"
)

/*
Pot
@Description: pot
*/
type Pot struct {
	// config
	Config *Config
	// cache
	cache *cache
	// cleaner
	cleaner *cleaner
}

/*
NewClient
@Desc:
@param: cfg
*/
func Init(cfg *Config) *Pot {
	plog.Infof("show trace info-> %v", cfg.ShowTrace)
	if cfg.ShowTrace {
		plog.SetLevel("trace")
	}

	client := &Pot{
		Config:  cfg,
		cache:   newCache(),
		cleaner: newCleaner(cfg.CleanerInterval),
	}

	// start cleaner
	client.startCleaner()

	plog.Infof("init finish...")
	return client
}

/*
getCache
@Desc: 获取cache
@receiver: c
@return: *cache
*/
func (p *Pot) getCache() *cache {
	if p.cache != nil {
		return p.cache
	}
	p.cache = newCache()
	return p.cache
}
