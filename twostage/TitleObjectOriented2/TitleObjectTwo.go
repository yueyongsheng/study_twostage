package main

import (
	"fmt"
)

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int64
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("工号：%s\n", e.EmployeeID)
	fmt.Printf("姓名：%s\n", e.Name) // 直接访问组合的 Name 字段
	fmt.Printf("年龄：%d\n", e.Age)  // 直接访问组合的 Age 字段
}

func main() {
	cs := Employee{
		Person:     Person{Name: "乐某某", Age: 18},
		EmployeeID: "001",
	}
	cs.PrintInfo()
}
