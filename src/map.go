/*
map是一种基于key-value的数据结构, 引用类型, 必须初始化之后才能使用
*/

// 使用make初始化

// 取值的时候, 要判断是否成功 v, ok = m["d"]

// for range 遍历

// delete 删除 (删除不存在的key不会报错)

// 顺序遍历 key 排序之后取值

package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m == nil) // 还没有初始化

	m = make(map[string]int, 10) // 支持自动扩容

	m["a"] = 1
	m["b"] = 1

	fmt.Println(m)

	v, ok := m["d"]

	if !ok {
		fmt.Println("不存在")
	}
	// 不存在是默认值 0
	fmt.Println(v)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "g")

	// panic: runtime error: index out of range
	// panic: assignment to entry in nil map

	/*
		// 没有对切片和map初始化
		ss = make([]map[int]string, 0, 10)
		ss[0][100] = "A"
		fmt.Println(ss)
	*/

	ss := make([]map[int]string, 10)
	ss[0] = make(map[int]string, 1)
	ss[0][10] = "A"
	fmt.Println(ss)

}
