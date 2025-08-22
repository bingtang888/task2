package main

import "fmt"

func main() {
	// 创建整形通道
	ch := make(chan int)

	// 生产者协程：发送数字
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			fmt.Println("接收：", num)
		}
	}()

	fmt.Scanln()
}
