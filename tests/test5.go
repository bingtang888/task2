package main

import (
	"fmt"
	"math"
)

// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 辅助函数：打印形状信息
func printShapeInfo(s Shape) {
	fmt.Printf("面积：%.2f, 周长：%.2f\n", s.Area(), s.Perimeter())
}

func main() {

	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形：")
	printShapeInfo(rect)
	fmt.Printf("详细: 宽=%.1f, 高=%.1f\n\n", rect.Width, rect.Height)

	// 创建 Circle实例
	circle := Circle{Radius: 4}
	fmt.Println("圆形：")
	printShapeInfo(circle)
	fmt.Printf("详细: 半径=%.1f\n\n", circle.Radius)

}
