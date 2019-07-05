package main

func main0() {
	ch := make(chan int, 0)

	//想法：自写自读——没人并发读取，永远写不进去，死锁
	ch <- 1

	<-ch
}
