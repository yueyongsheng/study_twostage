package main

import (
	"fmt"
	"sync"
)

//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ： sync.Mutex 的使用、并发数据安全。

var wg sync.WaitGroup //协程
var mutex sync.Mutex  //创建互斥锁（sync.Mutex）：保护共享资源，避免并发竞争

func main() {
	var counter int = 0

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
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			fmt.Printf("协程%d：完成1000次递增\n", coroutineID)
		}(i)
	}
	wg.Wait()
	fmt.Printf("\n所有协程执行完毕！最终计数器值：%d\n", counter)
}
