package main

import "fmt"

func main() {
	gen := func() <-chan int {
		dst := make(chan int)

		n := 1
		go func() {
			dst <- n

			if n == 5 {
				close(dst)
			}

			n++

		}()

		return dst
	}

	// 这里只发送了一个数据， 并且没有关闭 chan， 所以会死锁  fatal error: all goroutines are asleep - deadlock!
	for n := range gen() {
		fmt.Println(n)

		if n == 5 {
			break
		}
	}
}
