package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int            // 共享的计数器
		mutex   sync.Mutex     // 用于保护计数器的互斥锁
		wg      sync.WaitGroup // 用于等待所有协程完成
	)

	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment()
	}

	wg.Wait()

	fmt.Printf("最终计数器值：%d\n", counter)

}
