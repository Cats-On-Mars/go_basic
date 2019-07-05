package main

import "fmt"

var wage =1000

func main() {
	fmt.Printf("位置：%p,值：%d\n",&wage,wage)
	doubleWage(wage)
	fmt.Println("值传递结果 wage=",wage)
	doubleWage2(&wage)
	fmt.Println("引用传递结果 wage=",wage)
}
//值传递 是【拷贝式】的
func doubleWage(wage int){
	fmt.Printf("位置：%p,值：%d\n",&wage,wage)
	wage *=2
}

//引用传递 传【地址】
func doubleWage2(wage *int){
	fmt.Printf("位置：%p,值：%d\n",wage,*wage)
	(*wage) *=2
}