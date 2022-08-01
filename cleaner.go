package pot

import "time"

// cleaner 清理器
type cleaner struct {
	//执行间隔
	Interval time.Duration
	stop     chan int
}

/*
  Run
  @Desc: run value clean rountine
  @receiver: cl
  @param: c
*/
func (cl *cleaner) Run(c *cache) {
	ticker := time.NewTicker(cl.Interval)
	for {
		select {
		case <-ticker.C:
			cl.DeleteExpired()
		case <-cl.stop:
			ticker.Stop()
			return
		}
	}
}

func (c *Client) stopJanitor() {
	c.Cache.cleaner.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.Run(c)
}
