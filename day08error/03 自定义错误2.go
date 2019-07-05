package main

import (
	"time"
	"fmt"
	"math"
)

type myError struct{
	info string
	when time.Time
}

func (err *myError)Error()string{
	return fmt.Sprintln(err.info,err.when)
}

func newMyError(info string) error{
	ne := new(myError)
	ne.info = info
	ne.when = time.Now()
	return ne
}

func main() {
	var num float64
	for{
		fmt.Print("请输入数字:")
		fmt.Scan(&num)
		sqRoot, err := getRoot(num)
		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Printf("%.2f\n",sqRoot)
		}
	}

}


func getRoot(a float64) (sqRoot float64,err error){
	if a < 0{
		err = newMyError("错误：负数没有平方根")
		return
	}
	sqRoot = math.Sqrt(a)
	return
}