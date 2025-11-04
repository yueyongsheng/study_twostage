package main

import (
	"fmt"
	"sync"
	"time"
)

//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//考察点 ：协程原理、并发任务调度

// 1. 定义任务类型：无参数、无返回值的函数（可根据需求扩展参数/返回值）
type Task func()

// 2. 定义任务执行结果：存储任务名称和执行时间
type TaskResult struct {
	TaskName string        // 任务名称
	CostTime time.Duration // 执行耗时
}

// 3. 任务调度器：接收任务列表，并发执行并统计时间
func ScheduleTasks(tasks map[string]Task) []TaskResult {
	var (
		wg      sync.WaitGroup                      // 等待所有任务完成
		results = make([]TaskResult, 0, len(tasks)) // 存储所有任务结果
		mu      sync.Mutex                          // 互斥锁：防止多协程同时修改 results 切片（避免数据竞争）
	)

	// 3.1 注册任务数量：有多少个任务就等待多少个协程
	wg.Add(len(tasks))

	// 3.2 遍历任务列表，启动协程执行每个任务
	for taskName, taskFunc := range tasks {
		// 注意：循环中启动协程，必须用局部变量接收 taskName 和 taskFunc（避免闭包引用循环变量的问题）
		name := taskName
		task := taskFunc

		// 启动协程执行任务
		go func() {
			// 任务完成后，标记 WaitGroup 计数减 1（defer 确保无论任务是否出错都会执行）
			defer wg.Done()

			fmt.Printf("任务 [%s] 开始执行...\n", name)

			// 记录任务开始时间
			startTime := time.Now()

			// 执行任务（调用任务函数）
			task()

			// 计算任务执行时间（当前时间 - 开始时间）
			costTime := time.Since(startTime)

			// 3.3 存储任务结果：由于多个协程会同时修改 results 切片，必须用互斥锁保护
			mu.Lock() // 加锁：禁止其他协程同时操作
			results = append(results, TaskResult{
				TaskName: name,
				CostTime: costTime,
			})
			mu.Unlock() // 解锁：允许其他协程操作

			fmt.Printf("任务 [%s] 执行完成！\n", name)
		}()
	}

	// 3.4 主协程阻塞，等待所有任务执行完成
	wg.Wait()

	// 3.5 返回所有任务的执行结果
	return results
}

// 任务1：模拟耗时 1 秒的任务（比如文件读取）
func task1() {
	time.Sleep(1 * time.Second)
}

// 任务2：模拟耗时 500 毫秒的任务（比如数据库查询）
func task2() {
	time.Sleep(500 * time.Millisecond)
}

// 任务3：模拟耗时 1.5 秒的任务（比如网络请求）
func task3() {
	time.Sleep(1500 * time.Millisecond)
}

func main() {
	fmt.Println("=== 任务调度器启动 ===")

	// 4. 构造任务列表：key=任务名称，value=任务函数
	tasks := map[string]Task{
		"文件读取任务":  task1,
		"数据库查询任务": task2,
		"网络请求任务":  task3,
	}

	// 5. 调用调度器，并发执行任务并获取结果
	start := time.Now()
	results := ScheduleTasks(tasks)
	totalCost := time.Since(start)

	// 6. 输出统计结果
	fmt.Println("\n=== 所有任务执行完成，统计结果如下 ===")
	fmt.Printf("调度器总执行时间：%v\n", totalCost)
	fmt.Println("各任务执行详情：")
	for _, result := range results {
		fmt.Printf("  - %s：%v\n", result.TaskName, result.CostTime)
	}
}
