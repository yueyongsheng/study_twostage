package main

import (
	"fmt"
	"sync"
)

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制

var wg sync.WaitGroup

func main() {
	bufChan := make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			bufChan <- i
			// 打印通道当前状态（方便观察缓冲机制）：len(bufChan)是当前缓存的元素个数
			// fmt.Printf("生产者：发送整数 %d → 通道缓存数：%d/%d\n", i, len(bufChan), cap(bufChan))
		}
		close(bufChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range bufChan {
			fmt.Printf("消费者：收到整数 → %d\n", num)
		}
	}()

	wg.Wait()
	fmt.Println("\n程序结束")
}
