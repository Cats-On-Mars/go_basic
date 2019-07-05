package main

import (
	"fmt"
	"time"
)

func main() {
	demo1()
	demo2()
	fallthroughDemo()
	continueDemo()
	breakDemo()
}

func demo1(){
	//for 	起始条件；循环条件；增长条件
	for i:=1;i<11;i++{		//Go语言里只有i++，i--，没有++i，--i
		fmt.Println("爱天台，爱日耳曼战车")
		time.Sleep(500*time.Millisecond)
	}
}

func demo2(){
	for i:=1;i<11;i++{
		if i % 5 == 0{
			fmt.Println("最爱是天台")
		}else if i%2 == 1{
			fmt.Println("爱要强")
		}else{
			fmt.Println("更爱战车")
		}
		time.Sleep(500*time.Millisecond)
	}
}

func fallthroughDemo() {
	for i:=1;i<11;i++{
		switch i%5{

		//1，3都走爱要强
		case 1:
			//继续乡向下执行
			fallthrough
		case 3:
			fmt.Println("爱要强")

		//2，4都走爱战车
		case 2:
			fallthrough
		case 4:
			fmt.Println("更爱战车")

		//case 0：
		default:
			fmt.Println("最爱是天台")
		}
		time.Sleep(500*time.Millisecond)
	}
}

func continueDemo() {
	for i:=0;i<100;i++{
		if i%5==0{
			continue     //跳过本次循环的剩余部分，进入下次循环
		}
		fmt.Println(i)
	}
}

func breakDemo() {
	var cmd int
	for {
		fmt.Scanf("%d",&cmd)
		if cmd == -1{
			break		//结束循环
		}
		fmt.Println(cmd)
	}
}