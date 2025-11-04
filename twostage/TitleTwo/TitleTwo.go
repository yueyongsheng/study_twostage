package main

import "fmt"

//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
//考察点 ：指针运算、切片操作。

func multiplyByTwo(qiep *[]int) {
	qieplist := *qiep
	for v := range qieplist {
		qieplist[v] = qieplist[v] * 2
	}

}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	multiplyByTwo(&numbers)
	fmt.Println(numbers)
}
