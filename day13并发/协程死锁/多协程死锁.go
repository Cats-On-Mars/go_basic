package main

import (
	"sync"
	"fmt"
	"time"
)

func main1() {

	ch := make(chan int, 0)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		//没有人从另一端读取，永远写不进去，永远阻塞——死锁deadlock
		ch <- 1
		wg.Done()
	}()

	/*
	解决方案：有人写，有人读，死锁开
	wg.Add(1)
		go func() {
			<-ch
			wg.Done()
		}()*/

	wg.Wait()
}

func main() {
	ch := make(chan int, 3)   //TODO：是否跟管道的大小有关？ 无关！
	go func() {
		for i := 0; i < 9; i++ {
			ch <- i
		}

		/*		//解决方案
				//使用完毕，给一个明确的关闭信号，让循环读取的人结束读取
				close(ch)*/

	}()

	//循环读取，没有明确的关闭信号时，会永远阻塞读取——死锁
	for v := range ch {
		fmt.Println(v)
	}

	time.Sleep(time.Second)
}