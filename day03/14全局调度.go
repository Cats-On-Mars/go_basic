package main

import (
	"fmt"
	"utils/mathutil"
	"utils/aviutil"
	"utils/fileutil"
)

func main() {
	numsArray := mathutil.GetNums(123)
	fmt.Println(numsArray)

	aviutil.Ecrypt()
	fileutil.Hellofile()
}
