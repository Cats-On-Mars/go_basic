package main

import (
	"image/color"
	"os"
	"fmt"
	"encoding/json"
)

type Bear struct{
	Name string
	Color color.RGBA
	Sex bool
}


var (
	dM,cM Bear
)

func main() {
	enCoder()
	//deCoder()
}

func enCoder() {
	dM = Bear{"短毛", color.RGBA{1, 4, 6, 34}, true}
	cM = Bear{"长毛", color.RGBA{234, 150, 34, 23}, true}


	//创建文件并打开    Create()函数返回的是一个打开的文件的指针return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	file, err := os.Create("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json")
	defer file.Close()
	if err != nil {
		fmt.Println("文件创建失败")
	}


	encoder := json.NewEncoder(file)
	bears := []Bear{dM, cM}
	err = encoder.Encode(bears)
	if err !=nil {
		fmt.Println("编码失败")
	}else{
		fmt.Println("编码成功")
	}
}


func deCoder() {
	file,_ := os.Open("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json")
	defer file.Close()


	decoder := json.NewDecoder(file)
	bears := []Bear{}
	err := decoder.Decode(&bears)
	if err !=nil {
		fmt.Println("解码失败")
	}else{
		fmt.Println(bears)
	}
}
