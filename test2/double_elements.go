package main

import "fmt"

// DoubleElements 接收整数切片指针，将每个元素乘以2
func DoubleElements(nums *[]int) {
	// 检查指针是否为nil
	if nums == nil {
		return
	}

	// 通过指针访问切片，使用索引遍历并修改每个元素
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	// 测试示例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", numbers)

	// 传递切片指针给函数
	DoubleElements(&numbers)
	fmt.Println("翻倍后切片:", numbers)
}
