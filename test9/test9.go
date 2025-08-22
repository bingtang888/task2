package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建带缓冲的通道，缓冲区大小为10
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	// 生产者协程：发送100个整数
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 消费者协程：接收并打印整数
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收: %d\n", num)
		}
	}()

	wg.Wait()
	fmt.Println("所有数字处理完成")
}
