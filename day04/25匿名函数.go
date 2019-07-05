package main

import (
	"fmt"
	"time"
)

func main() {
	//三种典型用途
	demo()
	demo1()
	demo2()
}

func demo() {

	go func(a,b int) {
		fmt.Println(a+b)
	}(43,867)

	time.Sleep(500*time.Microsecond)
	fmt.Println("GAME OVER")
}

func demo1(){
	var hanshu = func(a,b int) int {
		return a+b
	}
	fmt.Printf("类型是：%T,值是：%v\n",hanshu,hanshu)  //值是地址
	ret := hanshu(34,45)
	fmt.Println(ret)
}

func demo2(){
	defer func() {
		fmt.Println("GAME OVER")
	}()

	for i:=1;i<3;i++{
		fmt.Println(i)
	}
}