package main

import "fmt"

func main() {
	var smlint uint8 = 123
	fmt.Println(uint64(smlint))

	var bigint uint64 = 123456789
	fmt.Println(uint8(bigint)) //转换时只保留到最低位
								//大整型强转为小整型可能会导致数据丢失
}
