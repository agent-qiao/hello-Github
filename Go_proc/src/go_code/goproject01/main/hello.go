// 输出hello world 
// 变量声明的三种方法
package main   // 表示该文件属于main包 go语言的每个文件必须归属于一个包
import "fmt"
func main(){
  var name string   // 采用这样的形式声明类型为string的“name“变量，默认值为空字符串
  fmt.Println("name = ", name)
  fmt.Println("hello,world!\n")

  var num int   // 先声明变量num
  num = 1       // 在使用变量
  fmt.Println("num = ", num)

  var myFloatNum = 1.1 // 自动推到变量
  fmt.Println("myFloatNum = ",myFloatNum)

  myVar, name := 19971010, "Rebecca"
  fmt.Println(name, ":", myVar)
}


