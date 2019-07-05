package main

import (
	"strconv"
	"fmt"
)

func main() {
	var ret interface{}

	//int to array (string - byteArray)
	ret = strconv.Itoa(123)
	fmt.Printf("类型是%T,值是%v\n",ret,ret)

	ret,err := strconv.Atoi("123")
	fmt.Printf("类型是%T,值是%v,错误:%v",ret,ret,err)
}
