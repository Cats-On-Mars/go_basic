package _1管道读写阻塞

import (
	"time"
	"fmt"
)

func main() {
	newTimer()

	after()

	stopReset()
}

func newTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())

	t := <-timer.C                  //相当于睡两秒  阻塞两秒
	fmt.Println("t=", t)
}

func after(){
	<- time.After(2*time.Second)    //return NewTimer(d).C       相当于睡两秒
	fmt.Println("时间到")
}


func stopReset(){
	timer := time.NewTimer(time.Second)

	go func(){
		<- timer.C
		fmt.Println("时间到，子协程可以打印了")
	}()

	timer.Stop()                //暂停计时 挂在那儿 被timer.C阻塞的协程永远读不出数据
	timer.Reset(5*time.Second)  //重置定时的时间，重新开始计时
								//如果timer还在等待中返回真，如果已停止返回假

	time.Sleep(6*time.Second)
}