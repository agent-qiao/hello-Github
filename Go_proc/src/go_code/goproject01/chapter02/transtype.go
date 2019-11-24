package main

import (
    "fmt"
    "strconv"
)

// 该文件用来演示golang中基本数据类型转换

func main(){
    fmt.Println("transtype\n")

    var number01 int64 = 100
    var number02 float64 = float64(number01)        // 将int64 转化为 float64
                                                    // 将int32 转化为 int32 的步骤是必须的
                                                    // golang中的数据类型转换必须是显示的，当存在类型不匹配是存在编译失败的情况
    fmt.Printf("number01 = %v\nnumber02 = %v\n",number01,number02)

    // 类型转换只是当时改变了当前的数值，原来变量的类型不会发生变化
    fmt.Printf("转换之前 number01的数据类型为 %T\n",number01)
    fmt.Printf("转换之后 number01的数据类型为 %T\n",number01)


    // sting 字符串转基本数据类型
    // 方式1 fmt.Sprintf(format string, a ...interface{}) 根据format生成格式化字符串并返回该字符串
    fmt.Println("string <-- basetype method1 :\n")

    var num1 int32 = 10
    var num2 byte = 'h'
    var num3 bool = true
    var num4 float64 = 12.30
    var str string // 默认空串

    str = fmt.Sprintf("%d", num1)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    str = fmt.Sprintf("%c", num2)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    str = fmt.Sprintf("%t", num3)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    str = fmt.Sprintf("%f", num4)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    // 方式2 利用strconv包函数
    fmt.Println("string <-- basetype method2 :\n")

    var num7 int32 = 10
    var num9 bool = true
    var num10 float64 = 12.30

    str = strconv.FormatInt(int64(num7), 8)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    // FormatFloat(f float64, fmt byte, prec, bitsize int)
    str = strconv.FormatFloat(num10, 'f', -1, 64)
    fmt.Printf("str'type : %T str = %q\n", str, str)

    str = strconv.FormatBool(num9)
    fmt.Printf("str'type : %T str = %q\n", str, str)


    // string转为其他基本类型
    // 利用strconv包函数来实现这一目标
    fmt.Println("string --> basetype method :\n")
    var strForconvert string = "true"
    var boolean bool
    boolean, _ = strconv.ParseBool(strForconvert)
    fmt.Printf("boolean'type : %T boolean = %t\n",boolean, boolean)

    var strForconvert1 string = "19971010"
    var myNum int64
    // 参数说明：1 str 2 basesize 3 返回值可以不改变类型的赋值位数，建议为64，如果溢出会报错
    myNum, _ =strconv.ParseInt(strForconvert1, 10, 64)
    fmt.Printf("myNum'type : %T myNum = %d\n", myNum, myNum)

    var strForconvert2 string = "19.92"
    var myNum1 float64
    myNum1, _ = strconv.ParseFloat(strForconvert2, 64)
    fmt.Printf("myNum1'type : %T myNum2 = %f\n", myNum1, myNum1)

    // 即使字符串的中以数字开始，转换无效时，只会返回默认值，如下
    var strForconvert3 string = "99hello"
    var myNum3 int64
    myNum3, _ =strconv.ParseInt(strForconvert3, 10, 64)
    fmt.Printf("myNum'type : %T myNum = %v\n", myNum3, myNum3)





}