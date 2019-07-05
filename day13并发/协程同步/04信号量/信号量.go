package main

import (
	"fmt"
	"time"
	"sync"
)

/*信号量：通过控制管道的“带宽”（缓存能力）控制并发数*/

func main() {

	//定义信号量为5“带宽”的管道
	sema = make(chan int, 5)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			ret := getPingfangshu(index)
			fmt.Println(index, ret)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//该函数只允许5并发执行
var sema chan int
func getPingfangshu(i int) int {
	sema <- 1
	<-time.After(2 * time.Second)
	<- sema
	return i
}
