package main

import (
	"fmt"
)
func main() {
	var str string = "hello, world"
	
	// 传统的遍历字符串
	
	for i := 0; i < len(str); i++ {
		fmt.Printf("i=%d, character=%c\n", i, str[i])
	}
	// 传统的字符串按照字节的方式来遍历，面对中文字符串串时，需要
	// 转换为切片来完成

	// 利用go语言的range关键字完成字符串的遍历
	for index, value := range str{
		fmt.Printf("index=%d, value=%c\n", index, value)
	}
	// for - range使用字符来的方式来便利，但是index是按字节来取
	// 读的

}