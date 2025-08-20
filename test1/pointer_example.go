package main

import "fmt"

func addTen(num *int) {
	*num += 10

}

func main() {
	value := 20
	fmt.Printf("修改前的值：%d\n", value)

	addTen(&value)
	fmt.Printf("修改后的值：%d\n", value)
}
