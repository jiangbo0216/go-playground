// 数组的长度是类型的一部分. 所以数组有很多的局限性

// a := [3]int {1,2,3}

// slice 是可变长度的序列, 基于数组类型做的一层分装, 自动扩容, 引用类型,
// slice 内部结构包含 地址, 长度, 容量, 切片一般用于快速操作一块数据集合
package main

import "fmt"

// 声明切片
// var name []T

// 长度和容量

// 可以由数组得到切片(长度和容量会受到限制)

// 切片的切片

// make函数

// 切片的本质就是一块连续的内存, 属于引用类型, 真正的数据保存在底层的数组中
// 切片之前不能比较, 唯一合法的比较操作是和nil比较, 一个nil值的切片没有底层数组. 一个为nil的切片的长度和容量都是0, 但是我们不能说一个长度和宽度都是0的切片一定是nil

// 切片的遍历, 支持索引遍历和for range 遍历

// api
/*
append
copy
*/

// 扩容策略

// ... 扩展运算符

// 使用append删除元素, 容量不会变

func main() {
	var s []int
	fmt.Println(s)
	fmt.Println(s == nil)

	s = []int{1, 2, 3}
	fmt.Println(s)

	// 长度和容量
	// 注意只有Printf可以使用输出占位符
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	a := [...]int{1, 2, 3, 4} // 生命一个数组
	a1 := a[1:2]              // 基于数组, 左闭右开, 可以省略
	// 切片的容量是指底层数组的容量(从切片第一个元素到数组最后一个元素的数量)
	fmt.Println(a1)

	//make(type, 0)
	m := make([]int, 5, 10) // 如果没有传10这个参数, 长度和容量一致
	fmt.Printf("s1=%v, len(m)=%d cap(m)=%d\n", m, len(m), cap(m))

	s11 := []string{"a", "b", "c"}
	// 错误的写法 panic: runtime error: index out of range
	// s11[3] = "d"
	s11 = append(s11, "d") // 返回值赋值给s11, append 追加元素, 原来底层的数组放不下的时候, go 底层就会吧底层的数组扩容 地址可能会变
	// 注意扩容的策略

	fmt.Println(s11)

	s22 := []string{"e", "f", "g"}
	s22 = append(s11, s22...)
	fmt.Println(s22)

	var s33 = /*不能声明用做copy的dist, 因为为nil*/ make([]string, 3) /*长度一定要大于source 切片*/

	copy(s33, s11)

	fmt.Println(s33)
}
