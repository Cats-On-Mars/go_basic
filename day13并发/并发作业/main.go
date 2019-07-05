package 并发作业

import (
	"time"
	"fmt"
)


var chAB = make(chan int)
var chBC = make(chan bool)
var chAmain = make(chan bool)
var oddEvenMap = map[bool]int{true:0,false:0}

func main() {
	go func(){
		for{
			n:= GenerateNum(100,999)
			if n == (n/100)*(n/100)*(n/100)+(n%100/10)*(n%100/10)*(n%100/10)+(n%10)*(n%10)*(n%10){
				fmt.Println("水仙花数",n)
				chAmain <- true
				break
			}
			chAB <- n
			time.Sleep(500*time.Millisecond)
		}

	}()

	go func() {
		for{
			n := <- chAB
			chBC <- IsOdd(n)
		}

	}()

	go func() {
		for{
			oddEvenMap[<- chBC]++
		}

	}()

	<- chAmain

}

