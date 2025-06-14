package retry

import (
	"fmt"
	"time"
)

// 周期性的检查条件是否满足
func Eventually(condition func() bool, waitFor time.Duration, tick time.Duration) bool {
	ch := make(chan bool, 1)

	timer := time.NewTicker(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			fmt.Println("condition not met in time")
			return false
		case <-tick:
			tick = nil
			go func() { ch <- condition() }()
		case v := <-ch:
			if v {
				return true
			}
			// 这里不能去掉(因为上面的for的tick.C是只执行一次)
			tick = ticker.C
		}
	}
}
