/*
自定义类型和类型别名
*/

package main

import "fmt"

type myInt int     //  自定义类型
type yourInt = int // 类型别名

// go 提供了一种自定义的数据类型, 可以封装多个基本数据类型, 这种数据类型叫结构体, 英文名称 struct, 即我们可以通过struct来定义自己的类型

// 定义结构体

/*
type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
	字段名 字段类型
}
*/

// 匿名结构体

// 创建指针类型结构体

type Person struct {
	name string
}

// 创建指针类型结构体(结构体是值类型)

func main() {
	var p Person
	fmt.Println(p) // {}

	p.name = "ha"

	// 匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "h"
	fmt.Printf("type:%T value:%v\n", s, s)
}
