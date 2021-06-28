package main

import "fmt"

/*
多维数据外层可以加...
*/

var a = [...][2]int{
	{1, 2},
	{1, 2},
	{1, 2},
}

func main() {
	fmt.Println(a)
}
