package main

import "fmt"

func f1() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("defer recover")
	}()
	panic("error")
	fmt.Println('b')
}

func f2() {
	fmt.Println("c")
}

func main() {
	f1()
	f2()
}
