package main

import (
	"unsafe"
	"fmt"
)

//size测的是字节
func main() {
	//int占8个字节
	var abc int = 123
	var def = 123
	ghi := 123
	fmt.Println(unsafe.Sizeof(abc))
	fmt.Println(unsafe.Sizeof(def))
	fmt.Println(unsafe.Sizeof(ghi))

	//int32占4个字节
	var mint int32 = 123
	fmt.Println(unsafe.Sizeof(mint))

	//占用1个字节
	var mint8 uint8 = 123
	var mbyte byte =123
	var mbool bool = true
	fmt.Println(unsafe.Sizeof(mint8))
	fmt.Println(unsafe.Sizeof(mbyte))
	fmt.Println(unsafe.Sizeof(mbool))

	//ASCII字符占1字节
	mbyte = 'c'
	//字符串默认分配16个字节
	var mstring string = "c"
	fmt.Println(unsafe.Sizeof(mbyte))
	fmt.Println(unsafe.Sizeof(mstring))

	//占用3个string的开销
	marray := [3]string{"bill","steve","mark"}
	fmt.Println(unsafe.Sizeof(marray))

}
