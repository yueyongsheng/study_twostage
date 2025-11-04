package main

import (
	"fmt"
	"math"
)

//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格

// 定义 Shape 接口（要求实现 Area 和 Perimeter 两个方法）
type Shape interface {
	Area() float64      // 计算面积，返回浮点型
	Perimeter() float64 // 计算周长，返回浮点型
}

// 定义 Rectangle 结构体（长方形）
type Rectangle struct {
	Width  float64 // 宽度
	Height float64 // 高度
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // 长方形面积 = 长 × 宽
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // 长方形周长 = 2×(长+宽)
}

// 定义 Circle 结构体（圆形）
type Circle struct {
	Radius float64 // 半径
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	// 圆形面积 = π × 半径²，math.Pi 是 Go 内置的 π 常量（精度足够）
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法（圆形周长也叫圆周长）
func (c Circle) Perimeter() float64 {
	// 圆形周长 = 2 × π × 半径
	return 2 * math.Pi * c.Radius
}
func main() {
	// 创建 Rectangle 实例（长方形：宽3，高4）
	rect := Rectangle{Width: 3, Height: 4}
	//创建 Circle 实例（圆形：半径2）
	circle := Circle{Radius: 2}

	// 直接调用结构体的方法（直观易懂）
	fmt.Println("===== 长方形 =====")
	fmt.Printf("宽度：%.1f，高度：%.1f\n", rect.Width, rect.Height)
	fmt.Printf("面积：%.2f\n", rect.Area())      // 保留2位小数，输出 12.00
	fmt.Printf("周长：%.2f\n", rect.Perimeter()) // 输出 14.00

	fmt.Println("\n===== 圆形 =====")
	fmt.Printf("半径：%.1f\n", circle.Radius)
	fmt.Printf("面积：%.2f\n", circle.Area())      // 输出 12.57（π×2²≈12.566）
	fmt.Printf("周长：%.2f\n", circle.Perimeter()) // 输出 12.57（2×π×2≈12.566）

	// 使用 Shape 接口变量接收不同结构体（体现接口的多态特性）
	var shape1 Shape = rect   // 接口变量存储长方形实例
	var shape2 Shape = circle // 接口变量存储圆形实例

	fmt.Println("\n===== 接口多态演示 =====")
	fmt.Printf("接口存储长方形的面积：%.2f\n", shape1.Area())     // 同样输出 12.00
	fmt.Printf("接口存储圆形的周长：%.2f\n", shape2.Perimeter()) // 同样输出 12.57
}
