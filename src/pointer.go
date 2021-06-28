/*
指针:
取地址操作符 & 和 取值操作符 * 是一对互补操作符
go不支持指针的运算
*/

// new 函数, 开辟一块内存

// make也是内存分配, 区别于new, 只能用于slice, map, chan的内存创建, make返回的内心就是这三个类型他们本身, 而不是他们的指针类型, 因为这三种类型本身就是引用类型, 没必要返回他们的指针了.

// new很少用, 一般用来给基本数据类型申请内存, 返回的是对应类型的指针
package main

import "fmt"

func main() {
	n := 18
	p := &n
	fmt.Printf("%T\n", p)

	m := *p
	fmt.Printf("%T\n", m)

	// 典型错误, 不能对nil使用*操作符
	// var a *int
	// *a = 100
	// fmt.Println(*a)

	a := new(int)
	*a = 100
	fmt.Println(*a)

}
