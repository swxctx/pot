package pot

import (
	"time"
)

// cleaner 清理器
type cleaner struct {
	//执行间隔
	Interval time.Duration
	stop     chan bool
}

/*
Run
@Desc: run value clean rountine
@receiver: cl
@param: c
*/
func (cl *cleaner) run(c *cache) {
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
func (c *Client) startCleaner() {
	go c.cleaner.run(c.getCache())
}

/*
stopJanitor
@Desc: 停止处理
@receiver: c
*/
func (c *Client) stopCleaner() {
	c.cleaner.stop <- true
}

/*
newCleaner
@Desc: new cleaner
@param: interval
@return: *cleaner
*/
func newCleaner(interval time.Duration) *cleaner {
	cleaner := &cleaner{
		Interval: interval,
		stop:     make(chan bool),
	}
	return cleaner
}
