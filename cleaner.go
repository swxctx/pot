package pot

import (
	"time"

	"github.com/swxctx/pot/plog"
)

// cleaner job for clean cache
type cleaner struct {
	// run interval
	Interval time.Duration
	// stop flag
	stop chan bool
}

/*
Run
@Desc: run value clean routine
@receiver: cl
@param: c
*/
func (cl *cleaner) run(c *cache) {
	plog.Infof("run cleaner job...")

	ticker := time.NewTicker(cl.Interval)
	for {
		select {
		case <-ticker.C:
			c.expiredRoutine()
		case <-cl.stop:
			ticker.Stop()
			return
		}
	}
}

/*
startCleaner
@Desc: 启动处理器
@receiver: c
*/
func (p *Pot) startCleaner() {
	go p.cleaner.run(p.getCache())
}

/*
stopJanitor
@Desc: 停止处理
@receiver: p
*/
func (p *Pot) stopCleaner() {
	p.cleaner.stop <- true
}

/*
newCleaner
@Desc: new cleaner
@param: interval
@return: *cleaner
*/
func newCleaner(intervalArg time.Duration) *cleaner {
	interval := intervalArg
	if interval.Seconds() == 0 {
		interval = time.Duration(1) * time.Second
	}
	cleanerVal := &cleaner{
		Interval: interval,
		stop:     make(chan bool),
	}

	plog.Infof("cleaner init finish, interval second-> %vs", interval.Seconds())
	return cleanerVal
}
