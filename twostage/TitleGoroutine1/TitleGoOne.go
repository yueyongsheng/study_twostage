package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行

var wg sync.WaitGroup

func OddNumber() {
	defer wg.Done()
	// 奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数：%d\n", i)
	}
	fmt.Println("奇数协程结束")
}

func EvenNumber() {
	defer wg.Done() // 协程结束时通知 WaitGroup
	// 偶数
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数：%d\n", i)
	}
	fmt.Println("偶数协程结束")
}

func main() {

	wg.Add(2) // 设置等待的协程数量

	go OddNumber()  // 启动奇数协程
	go EvenNumber() // 启动偶数协程

	wg.Wait() // 等待所有协程完成
	fmt.Println("所有协程执行完毕，程序结束")
}
