package main

import "fmt"

//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
//考察点 ：指针的使用、值传递与引用传递的区别。

func addTen(numPtr *int) {
	*numPtr += 10
	fmt.Println("增加10，当前值为", *numPtr)
}

func main() {
	//初始值
	originalNum := 20
	fmt.Println("原始变量值 =", originalNum)
	fmt.Println("内存地址 =", &originalNum) // 获取变量地址

	//调用函数：传递变量的内存地址（指针）
	addTen(&originalNum)

	//修改后的结果
	fmt.Println("主函数最终值", originalNum)
}
