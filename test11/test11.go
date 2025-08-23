package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int64          // 共享计数器（必须使用int64类型，原子操作要求）
		wg      sync.WaitGroup // 等待所有协程完成
	)

	// 每个协程对计数器进行1000次原子递增
	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			// 原子操作：将counter的值加1，返回新值
			atomic.AddInt64(&counter, 1)
		}
	}

	// 启动10个协程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment()
	}

	// 等待所有协程完成
	wg.Wait()

	// 原子操作：读取counter的当前值
	fmt.Printf("最终计数器值: %d\n", atomic.LoadInt64(&counter))
}
