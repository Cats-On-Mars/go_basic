package main

import "fmt"

func main() {
	demo1()
	demo2()
	demo3()
}


//嵌套循环
func demo1() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("你妹%d%d\t", i, j)
		}
		fmt.Println()
	}
}

//逆向遍历   逆序打印/跳着打印
func demo2(){
	for i:=10;i>0;i--{
		fmt.Println(i)
	}
	for j:=10;j>0;j-=2{
		fmt.Println(j)
	}
	for l:=10;l>0;l-=5{
		fmt.Println(l)
	}
}

//跳出指定循环
func demo3(){
	MARK:for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("你妹%d%d\t", i, j)
			if i==5 && j==5{
				//break //直接跳出当下循环，会继续下一个循环
				break MARK //跳出标记的那层循环
			}
		}
		fmt.Println()
	}
}