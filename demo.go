package goBiliblili

import (
	"time"
	"sync"
	"fmt"
)



func main() {
	sema = make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			getdouble(index)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

var sema chan int
func getdouble(i int){
	sema<-1
	i *= 2
	time.Sleep(time.Second)
	fmt.Println(i)
	<- sema
}