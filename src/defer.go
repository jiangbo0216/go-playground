package main

import "fmt"

// defer 语句会将后面跟随的语句进行延迟处理, 在defer函数即将返回时, 将延迟处理的语句按 defer 定义的逆序进行执行, 先进后出

// defer 多用于函数结束之前释放资源 (文件句柄, 数据库链接, socket)

// return 这个语句在底层不是原子操作, 它分为返回值赋值和RET指令两步, 而 defer 语句执行的时间就在返回值操作后, RET指令之前.

func f1() int {
	x := 5
	defer func() {
		x++
	}()

	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()

	return 5
}

func main() {

	fmt.Println(f1())
	fmt.Println(f2())

	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
