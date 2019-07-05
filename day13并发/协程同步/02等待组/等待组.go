package main

import (
	"time"
	"fmt"
	"sync"
)

//主协程等待子协程全部结束：通过管道阻塞
func main0() {
	chanRets := make(chan int, 3)
	fmt.Println(len(chanRets),cap(chanRets))
	for i := 0; i < 3; i++ {
		go func(index int) {
			ret := getFibonacci(index)
			chanRets <- ret
			fmt.Println(index,ret)
		}(i)
	}

	for{
		if len(chanRets)==3{
			time.Sleep(time.Nanosecond)
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		//等待组中协程数+1(主协程中)
		wg.Add(1)

		go func(index int) {
			ret := getFibonacci(index)
			fmt.Println(index,ret)
			//等待组中协程数-1(子协程中)
			wg.Done()
		}(i)
	}

	//阻塞至等待组中的协程数为0
	wg.Wait()

}

func getFibonacci(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	<-time.After(3 * time.Second)
	return x
}
