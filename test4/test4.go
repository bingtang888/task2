package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型，接收一个字符串参数，返回执行结果
type Task func(string) string

// 任务调度器：接收任务列表和任务参数，并发执行并统计时间
func scheduleTasks(tasks []Task, params []string) {
	var wg sync.WaitGroup
	taskCount := len(tasks)

	// 确保任务数量和参数数量一致
	if taskCount != len(params) {
		fmt.Println("任务数量与参数数量不匹配")
		return
	}

	wg.Add(taskCount)

	for i := 0; i < taskCount; i++ {
		// 使用索引副本，避免闭包引用问题
		idx := i
		go func() {
			defer wg.Done()

			// 记录开始时间
			startTime := time.Now()

			// 执行任务
			result := tasks[idx](params[idx])

			// 计算执行时间
			duration := time.Since(startTime)

			// 输出结果和执行时间
			fmt.Printf("任务 %d 执行完成\n", idx+1)
			fmt.Printf("  输入参数: %s\n", params[idx])
			fmt.Printf("  执行结果: %s\n", result)
			fmt.Printf("  执行时间: %v\n\n", duration)
		}()
	}

	// 等待所有任务完成
	wg.Wait()
	fmt.Println("所有任务执行完毕")
}

func main() {
	// 定义示例任务1：模拟数据处理
	task1 := func(input string) string {
		// 模拟任务执行耗时
		time.Sleep(500 * time.Millisecond)
		return "数据处理完成: " + input
	}

	// 定义示例任务2：模拟计算任务
	task2 := func(input string) string {
		// 模拟任务执行耗时
		time.Sleep(800 * time.Millisecond)
		return "计算结果: " + input
	}

	// 定义示例任务3：模拟网络请求
	task3 := func(input string) string {
		// 模拟任务执行耗时
		time.Sleep(300 * time.Millisecond)
		return "网络响应: " + input
	}

	// 任务列表
	tasks := []Task{task1, task2, task3}
	// 对应每个任务的参数
	params := []string{"用户数据", "100*200/300", "https://example.com"}

	// 执行任务调度
	scheduleTasks(tasks, params)
}
