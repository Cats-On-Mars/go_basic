package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += "golang"
	fmt.Println(str1)

	var strArr = []string{"bill","steve","mark"}
	strArr = append(strArr, "haihui")
	retStr := strings.Join(strArr, " & ")
	fmt.Println(retStr)

}