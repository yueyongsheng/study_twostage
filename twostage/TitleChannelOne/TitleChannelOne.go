package main

import (
	"fmt"
	"sync"
)

//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信

var wg sync.WaitGroup

func main() {
	numchan := make(chan int)
	// 生成协程：wg加1（标记有一个协程需要等待）
	wg.Add(1)
	go func() {
		defer wg.Done() // 协程结束时，wg减1（标记协程完成）
		for i := 1; i <= 10; i++ {
			numchan <- i
			fmt.Printf("生成协程：已发送整数 %d\n", i)
		}
		close(numchan)
	}()
	// 接收协程：wg加1
	wg.Add(1)
	go func() {
		defer wg.Done() // 协程结束时，wg减1
		for num := range numchan {
			fmt.Printf("接收协程：收到整数 %d\n", num)
		}
	}()
	// 主协程等待所有标记的协程完成（wg计数器减到0）
	wg.Wait()
	fmt.Println("执行完毕")
}
