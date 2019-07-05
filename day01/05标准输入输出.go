package main

import "fmt"

/*

*/

func getArea1(radius float32) float32{
	return 3.14 * radius * radius
}

func getArea2(name string,radius float32){
	area := 3.14 * radius * radius
	fmt.Printf("亲爱的%s：您有%.2f分薄面",name,area)
}

func main() {
	demo()
	demo2()
}

func demo() {

	fmt.Println("请输入你的半径：")

	var r float32
	fmt.Scanf("%f", &r)
	//接收用户的输入，写入到r的地址中，&r取r的地址

	area := getArea1(r)
	fmt.Printf("你的半径有%.2f,你的面子有%.2f\n", r, area)
}

func demo2() {
	var name string
	var r float32

	fmt.Println("请输入面部半径+姓名：")
	fmt.Scanf("%f+%s",&r,&name)
	//接收用户的两次输入，以"+"为分隔符，分别写入到r的地址、name的地址

	getArea2(name,r)
	//将用户输入的name和r作为参数，传给函数进行处理
}
