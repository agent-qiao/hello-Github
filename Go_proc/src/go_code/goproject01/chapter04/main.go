package main
import (
	"fmt"
)
func main(){
	var name 	string
	var age  	byte
	var sal  	float32
	var isPass 	bool
	// scanln()一次读取一行
	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)
	fmt.Printf("name:%v \n age:%v\n sal:%v \n isPass:%v\n",name, age, sal, isPass)
	

	// 第二种使用scanf输入
	fmt.Println("请输入name/age/sal/ispass with blank")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name:%v \n age:%v\n sal:%v \n isPass:%v\n",name, age, sal, isPass)
}