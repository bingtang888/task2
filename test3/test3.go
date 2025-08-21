package main

import (
	"fmt"
	"sync"
)

// 打印奇数（1-10）
func printOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数：%d\n", i)
	}
}

// 打印偶数（1-10）
func printEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数：%d\n", i)
	}
}

func main() {
	var wg sync.WaitGroup

	//注册两个待等待的协程
	wg.Add(2)

	// 启动协程打印奇数
	go printOdds(&wg)

	// 启动协程打印偶数
	go printEvens(&wg)

	wg.Wait()
	fmt.Println("所有数字打印完成")
}
