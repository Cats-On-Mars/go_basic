package main

import "fmt"

//go build hello.go 生成可执行程序    ./hello 执行编译好的程序
// go build -o xxx  hello.go   (将hello.go生成名为xxx的可执行程序、指定输出位置）
//go run hello.go 编译执行合二为一


//快捷键是command+/

/*
快捷键是option+command+/
*/

//从调用生成函数 alt+enter
//shift+alt 同时选中 shift+>向右选择
//封装函数 alt+command+m

func main() {
	fmt.Println("hello golang,去浪吧")
	fmt.Println()

	getUint()
	getInt()

	fmt.Println("byte=uint8")
	fmt.Println("rune=int32")
	fmt.Println("uint=uint32或uint64")
	fmt.Println("int=与uint一样大小？？？")
	fmt.Println("uintptr=无符号整型，用于存放一个指针")
}



func getUint(){
	var a uint64 = 1
	for i:=1;i<65;i++{
		a = a * 2
		switch i{
		case 8:
			fmt.Printf("uint8的范围是%d--%d\n",0,a-1)
		case 16:
			fmt.Printf("uint16的范围是%d--%d\n",0,a-1)
		case 32:
			fmt.Printf("uint32的范围是%d--%d\n",0,a-1)
		case 64:
			fmt.Printf("uint64的范围是%d--%d\n\n",0,a-1)
		}
	}
}


func getInt() {
	var a int64 = 1
	for i := 1; i < 64; i++ {
		a = a * 2
		switch i+1 {        //因为最高位用来表示整数还是负数 所以取值范围减半
		case 8:
			fmt.Printf("int8的范围是%d--%d\n", -a , a-1)
		case 16:
			fmt.Printf("int16的范围是%d--%d\n", -a , a-1)
		case 32:
			fmt.Printf("int32的范围是%d--%d\n", -a , a-1)
		case 64:
			fmt.Printf("int64的范围是%d--%d\n\n", -a , a-1)
		}
	}
}


