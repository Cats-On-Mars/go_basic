package main

import "fmt"

//声明布尔型变量
//var isStupid bool = true
//var isStupid bool = false
var isStupid bool   //默认值为false

/*声明数值型变量*/
var age int
var rmb float32
var vector complex64  //复数 var vector complex64 = 3+4i

//声明字符串变量
var name string



/*声明复合型变量*/
var ages [10]int //数组 [10]int{1,2,3,4,5,6,7,8,9,2}
var heights []int //可变长的整型切片  []int{1,2,3,4,5}
var info map[string]float32 //映射（键1：值1，键2：值2...）map[string]float32{"资产":123.45,"负债":678.90}
/*var worker struct{ //结构体
	name string
	age int
	rmb float32
}*/
var worker struct{}
//var weight int = 300
//var weightPoint *int = &weight //整型指针，存放的不是300，而逝放置300的内存地址
var weightPoint *int

//声明函数变量
var getArea func()

//声明一个接口
var Animal interface{}

func main() {
	a := 3   //冒等方法只能声明在函数内
	fmt.Print(a)

	fmt.Printf("isStupid的类型是%T，默认值是%v\n",isStupid,isStupid)
	fmt.Printf("age的类型是%T，默认值是%v\n",age,age)
	fmt.Printf("rmb的类型是%T，默认值是%v\n",rmb,rmb)
	fmt.Printf("vector的类型是%T，默认值是%v\n",vector,vector)
	fmt.Printf("name的类型是%T，默认值是%v\n",name,name)
	fmt.Printf("ages的类型是%T，默认值是%v\n",ages,ages)
	fmt.Printf("heights的类型是%T，默认值是%v\n",heights,heights)
	fmt.Printf("info的类型是%T，默认值是%v\n",info,info)
	fmt.Printf("worker的类型是%T，默认值是%v\n",worker,worker)
	fmt.Printf("weightPoint的类型是%T，默认值是%v\n",weightPoint,weightPoint)
	fmt.Printf("getArea的类型是%T，默认值是%v\n",getArea,getArea)
}

