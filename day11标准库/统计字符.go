package main

import (
	"io/ioutil"
	"fmt"
)

//unicode码
func main() {
	bytes, err := ioutil.ReadFile("/Users/shenhaihui/go/src/learn/goBiliblili/day10/成语释义.json")
	if err != nil {
		fmt.Println("读取文件出错", err)
		return
	}

	data := string(bytes)
	var figurecount,lettercount,blankcount,othercount int
	for _,v := range data{
		switch {
		case v>'0'&&v<'9':
			figurecount++
		case v>'a'&&v<'z'||v>'A'&&v<'Z':
			lettercount++
		case v==' '||v=='\t':
			blankcount++
		default:
			othercount++
		}
	}
	fmt.Println(figurecount,lettercount,blankcount,othercount)
}
