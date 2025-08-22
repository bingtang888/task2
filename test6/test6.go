package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d, 工号：%s\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	// 创建 Employee 实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E1001",
	}
	emp.PrintInfo()
}
