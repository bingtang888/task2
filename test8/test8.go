package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	//生产者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	//消费者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收：", num)
		}
	}()

	wg.Wait()
	fmt.Println("所以数字处理完成")
}
