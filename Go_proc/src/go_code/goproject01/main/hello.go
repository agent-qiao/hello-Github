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

  // int8 int16 int32 int64
  // uint8 uint16 uint32 uint64
  // 测试一下int8的范围

  // 变量不能超过数据类型的范围
  var testNum  int8= -127  // 不能超过int8的范围
  fmt.Println("signed num = ",testNum)
  var testNum2  uint8= 255   // 不能超过uint8的范围
  fmt.Println("unsigned num = ",testNum2)

  // uint、int 默认为有符号 且和计算机实现有关
  var mychar byte= 'a' 
  fmt.Println("mychar : ",mychar)

}


