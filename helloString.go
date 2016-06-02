//示例代码
//go run helloString.go 
package main
import "fmt"
var frenchHello string  // 声明变量为字符串的一般方法
var emptyString string = ""  // 声明了一个字符串变量，初始化为空字符串
func main() {
	fmt.Println("**************Result*****************")
	fmt.Println("【Normal Declare of String】")
    //no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
    japaneseHello := "Konichiwa"  // 同上
    frenchHello = "Bonjour"  // 常规赋值
    fmt.Printf("hello %s\n",frenchHello)
    fmt.Printf("hello JP: %s\n",japaneseHello)
    frenchHello :="Bonjour2"
    //c := []byte(s) 
    fmt.Printf("hello %s\n",frenchHello)

    frenchHello ="Bonjour3"
    //c := []byte(s) 
    fmt.Printf("hello %s\n",frenchHello)


 	fmt.Println("【Compose of String】")
    //字串結合列印
	s := "hello,"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)

	fmt.Println("【Like Substring】")
	//字串substring功能
	s1 := "hello"
	s1 = "c" + s[2:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s1)



	fmt.Println("【Array】")
	//Array 陣列
	var arr [10]int  // 声明了一个int类型的数组
	arr[0] = 42      // 数组下标是从0开始的
	arr[1] = 13      // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0	


	fmt.Println("【MAP】")
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int
	// 另一种map的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3




}