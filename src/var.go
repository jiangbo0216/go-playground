package main

import "fmt"

// 变量: 单个声明, 批量声明, 声明同时赋值, 类型推导, 简短声明, 匿名变量
// 常量: 单个声明, 批量声明(iota, 省略),
const pi = 3.14
const (
	n0 = 100
	n2
	n3
	n4
)
const (
	n5 = iota // 0
	n6        // 1
	n7
	n8
)

// var name string
// var age int
// var sex int

// 批量声明
var (
	name string
	age  int
	sex  int
)

// 声明变量同时赋值
var s1 string = "hh"

// 类型推导
// var s2 = "ddd"

// 必须声明再使用, 声明的变量必须使用

func main() {
	// 简短变量声明, 只能在函数中使用
	// s3 := "ggg"
	// 匿名变量
	_, s0 := "good", "good"
	fmt.Println(s0)

	name = "hello"
	age = 16
	sex = 0
	fmt.Println("========")
	fmt.Printf("name: %s\n", name)
	fmt.Printf("sex: %v\n", sex)
	fmt.Printf("age: %v\n", age)
	fmt.Print("========")
}
