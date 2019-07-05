package _1管道读写阻塞

import (
	"fmt"
	"time"
)


func producter(out chan<- int){
	for i := 0; i<10;i++ {
		fmt.Println("生产：",i*i)
		out <- i*i
	}
	close(out)
}

func consumer(in <-chan int){
	for num := range in{
		fmt.Println("消费：", num)
		time.Sleep(time.Second)
	}
}

func main() {
	chPC := make(chan int)
	go producter(chPC)
	consumer(chPC)
}
