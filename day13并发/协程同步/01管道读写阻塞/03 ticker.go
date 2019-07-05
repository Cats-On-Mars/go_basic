package main

import (
	"time"
	"fmt"
)


/*
timer:一次性的定时任务
ticker:周期性的定时任务
*/

func main() {
	ticker := time.NewTicker(time.Second)

	i := 0

	for {
		t := <-ticker.C
		i++
		fmt.Println("i=", i, t)

		if i == 5 {
			ticker.Stop()
			break
		}
	}




}
