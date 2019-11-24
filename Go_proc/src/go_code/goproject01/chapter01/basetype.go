// 输出hello world 
// 变量声明的三种方法
package main   // 表示该文件属于main包 go语言的每个文件必须归属于一个包
// import "fmt"
// import "unsafe"
import(
  "fmt"
  "unsafe"
)
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

  // 用fmt.printf("%T",a)输出变量的类型
  fmt.Println("\n########\n")
  var c = 100 // c的类型默认是int
  fmt.Printf("c的类型为%T\n",c)

  // 如何查看变量占用的字节大小
  var x int64 = 19971010
  fmt.Printf("x变量类型为%T,x变量占用的字节大小%d\n",x,unsafe.Sizeof(x))

  // 浮点数的使用
  var myfloat1 float32 = -19.2
  var myfloat2 float64 = -19.30
  fmt.Println("myfloat1",myfloat1,"myfloat2",myfloat2)
  var myFloat3 = 19.4               // float 默认类型为float64
  fmt.Printf("myFloat3变量类型为 %T\n",myFloat3)

  // 字符的使用
  var mychar1 byte = 'a'
  var mychar2 byte = '0'
  var mychar3 int = '妍'  //汉字占用三个字节
  fmt.Printf("mychar1 = %c 对应的码值为%d\n",mychar1,mychar1)
  fmt.Printf("mychar2 = %c 对应的码值为%d\n",mychar2,mychar2)
  fmt.Printf("mychar = %c 对应的码值为%d\n",mychar3,mychar3)

  // 字符串的使用
  var mystr1 string = "zhongguo"
  mystr2 := "nihao\n"
  fmt.Println(mystr1,mystr2)
  mystr3 := "hello" + "hello"+  // 多个字符串的拼接需要以“+”作为一行的结束
  "world"
  fmt.Println(mystr3)

}


