package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	//demo51()
	//demo52()
	//demo53()
	demo54()
}

func demo54(){
	//type ErrorHandling int
	//const (
	//	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	//	ExitOnError                          // Call os.Exit(2).
	//	PanicOnError                         // Call panic with a descriptive error.
	//)
	flagset := flag.NewFlagSet("新命令行参数设置方法", flag.ExitOnError)

	i := flagset.String("new", "", "新命令行参数设置方法测试")
	fmt.Println(*i)

}

func demo53() {
	//05xxx.exe -name bill -age 60 -money 1234567890.12 -isstupid true
	//开辟4块内存，用于存储cmd-line参数
	var name string
	var age int
	var money float64
	var isstupid bool
	fmt.Println(&name, &age, &money, &isstupid)

	//分别从命令行读取name，age,money,isstupid参数，存入上边开辟的内存地址中
	flag.StringVar(&name, "name", "无名氏", "用户名")
	flag.IntVar(&age, "age", 0, "年龄")
	flag.Float64Var(&money, "money", 0, "财富")
	flag.BoolVar(&isstupid, "isstupid", true, "是否愚蠢")

	//解析命令行参数
	flag.Parse()

	//得到结果
	fmt.Println(name, age, money, isstupid)
}

//方式2
func demo52() {
	//05xxx.exe -name bill -age 60 -money 1234567890.12 -isstupid true

	//name=参数名字，value=参数的默认值，usage=参数的说明信息，返回值=存储了【用户输入的参数】的内存地址
	namePtr := flag.String("name", "无名氏", "用户名")
	agePtr := flag.Int("age", 0, "年龄")
	moneyPtr := flag.Float64("money", 0, "财富")
	isstupidPtr := flag.Bool("isstupid", true, "是否愚蠢")

	//解析参数
	flag.Parse()

	//从上边返回的地址中读取【用户输入的参数】
	fmt.Println(*namePtr, *agePtr, *moneyPtr, *isstupidPtr)
}

func demo51() {
	fmt.Println(os.Args)
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
