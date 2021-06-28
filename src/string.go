package main

import "fmt"

// 字符串
var s = "hello"

// 字符
var s1 = 'h'

// 字符转义符
/*
\r
\n
\'
\"
\\
*/

// 多行字符串 ``

// api
/*
len
+/fmt.Sprintf
strings.Split
strings.Contains
strings.HasPrefix(HasSufffix)
strings.Index(LastIndex)
strings.Join
*/

// 遍历
/*
for _, c = range s {
	fmt.Println(c)
	fmt.Println("%c", c)
}
*/

// go语言中为了处理非ascii码类型的字符, 定义了新的rune类型
// rune 类型实际上类型别名, 实际上是int32类型

// 修改字符串. 需要转换成[]rune 或者 []byte, 完成之后在转换成string类型, 都会重新分配内存, 并复制字节数组
// byte 是 int8

// 类型转换


func main() {
	s2 := fmt.Sprintf("%s%s", "hi", "good")
	fmt.Println(s2, "hi"+"good")
	// byte字节的数量
	fmt.Println(len("你好"))

	// 字符串修改
	sss := "你好"
	ssss := []rune(sss)
	ssss[0] = 'n'
	fmt.Println(string(ssss))

	// rune string
	c1 := "h"
	c2 := 'h'

	fmt.Printf("c1:%T, c2: %T\n", c1, c2)
	// 字符可以转数字, 字符串不可以
	fmt.Printf("%d\n", c2)

}
