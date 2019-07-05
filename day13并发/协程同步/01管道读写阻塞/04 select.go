package _1管道读写阻塞

import "fmt"


func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func(){
		for i := 0; i < 8; i++ {
			num := <- ch
			fmt.Println(num)
		}
		quit <- true
	}()


	fib(ch,quit)
}

func fib(ch chan<- int,quit <-chan bool){
	x := 1
	y := 1
	for{
		select{								//select 在能走的路里面随机选一条
		case ch <- x:
			x,y = y,x+y

		case flag := <- quit:
			fmt.Println("flag=",flag)
			return
		}
	}

}
