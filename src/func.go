/*
函数可以作为函数的参数
函数可以作为函数的返回值
匿名函数
*/

package main

import "fmt"

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数但有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但是有返回值
func f3() int {
	return 3
}

// 返回值可以命名, 也可以不命名, 命名的返回值相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 需要加上 return
}

// 多个返回值

// 参数的类型简写
func f5(x, y int) int {
	return x + y
}

// 变长的参数

func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// 没有默认参数这个概念(所见即所得)

func main() {
	f7("1", 1, 2, 3)
}
