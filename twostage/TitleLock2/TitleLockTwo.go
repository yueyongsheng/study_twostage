package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全

var wg sync.WaitGroup //协程

func main() {
	var counter int32 = 0

	const (
		goroutineNum = 10   // 10个协程
		incrementNum = 1000 // 每个协程递增1000次
	)

	wg.Add(goroutineNum)
	for i := 1; i <= goroutineNum; i++ {
		go func(coroutineID int) {
			// 传入协程ID，方便日志区分
			defer wg.Done()
			for j := 0; j < incrementNum; j++ {
				atomic.AddInt32(&counter, 1)
			}
			fmt.Printf("协程%d：完成1000次递增\n", coroutineID)
		}(i)
	}
	wg.Wait()
	finalCount := atomic.LoadInt32(&counter)
	fmt.Printf("最终计数器值：%d\n", finalCount)
}
