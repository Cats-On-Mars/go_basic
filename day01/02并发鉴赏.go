package main

import (
	"fmt"    //格式化包
	"time"   //系统时间包
)


//单协程内部，自上而下【串行】执行
//多协程，【并行】执行（同时执行）
func doSomething(i int){
	//time.Sleep(1*time.Second) 睡太久还来不及执行子协程，主协程就结束了
	//这里的代码是跑在子协程中
	fmt.Println("我是战狼小分队",i)
}

func main() {
	//main中的代码都是跑在主协程中
	//这里是主协程
	fmt.Println("总部要行动了")

	//其他语言并发：进程-线程（并发的资源开销大）
	//并发100条协程（微线程）
	for i:=1;i<101;i++{

		//go开并发是多么的容易啊！！！
		//开辟独立的子协程中并发执行doSomething(i)  //"开辟"是由主协程做的，doSomething()里面的事是子协程做的
		go doSomething(i)
	}

	//这里是主协程
	fmt.Println("总司令晕过去了")
	time.Sleep(1*time.Second)   //如果主协程结束，子协程也结束
	fmt.Println("GAME OVER。。。")


}
