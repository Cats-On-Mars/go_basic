package main

import "fmt"

func main() {
	//demoa1()
	demoa2()
}

//存放指针的指针
func demoa2() {
	var age int = 20
	//agePtr中存放的是age的内存地址
	var agePtr *int = &age
	fmt.Printf("type=%T,value=%p\n", agePtr, agePtr)
	//apPtr是一个指向指针（agePtr）的指针
	var apPtr **int = &agePtr
	//type=**int,value=0xc042082018
	fmt.Printf("type=%T,value=%p\n", apPtr, apPtr)
	//*apPtr得到agePtr（存放age的地址），*（*apPtr）得到age的值20
	fmt.Println(**apPtr)
}

//指针就是地址
//所有普通类型都有对应的指针类型：*int,*float64,*bool,*book,*[10]int,*map[string]interface{}
// 对普通变量取地址 &age
// 取地址中的值 *agePtr
func demoa1() {
	var age int = 20
	//对age取地址，赋值给整型指针agePointer
	var agePointer *int = &age
	//类型：整型指针，值：age的内存地址
	//type=*int,value=0xc042062080
	fmt.Printf("type=%T,value=%p\n", agePointer, agePointer)
	//*agePointer取地址中的值，20
	fmt.Println("agePointer所指向的地址中的值", *agePointer)
}