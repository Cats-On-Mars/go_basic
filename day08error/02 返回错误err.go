package main

import (
	"errors"
	"fmt"
)

func main() {

	ret, err := getAreaII(-5)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(ret)
	}

}


func getAreaII(r float32)(area float32,err error){
	if r<0{
		err = errors.New("半径不能为负")
		return
	}

	area = 3.14*r*r
	return

}