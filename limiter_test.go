package limiter

import (
	"fmt"
	"testing"
	"time"
)

// 模拟 2秒钟有 10次的数据可以放行。
var (
	number   = 10
	waitTime = time.Second * 2
)

func TestLimiter(t *testing.T) {
	l := NewLimiter(number)
	// 这些是嗷嗷待哺的用户。
	for i := 0; i < 1000; i++ {
		go func() {
			l.Limit()
			fmt.Println("hi")
		}()
	}
	// 这个goroutine是控制的加入令牌这个功能。
	go func() {
		// 使用一个ticker然后来控制间隔。
		t := time.NewTicker(waitTime)
		for {
			select {
			case <-t.C:
				fmt.Println("启动加入令牌的措施。")
				l.Add()
			default:
			}
		}
	}()
	time.Sleep(time.Second * 100)
}
