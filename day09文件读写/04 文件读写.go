package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
	"image/color"
)

func main() {

	//打开与关闭文件
	//openClose()

	//用bufio读取文件
	//readByBufio()

	//用bufio写入文件  创写/覆写/追加
	//writeByBufio()

	//先写后读文件   读写-文件句柄
	//writeRead()

	//用ioutil快捷读取文件
	//readByIoutil()

	//用ioutil快捷写入文件
	//writeByIoutil()

	//判断文件是否存在
	isExist("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json")
	isExist("/Users/shenhaihui/go/src/learn/goBiliblili/day09/仓鼠们.json")
}


func openClose(){
	file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json", os.O_RDONLY,0666)
	//os.Open(name) === os.OpenFile(name,O_RDONLY, 0)
	//flag: O_RDONLY   O_WRONLY   O_RDWR    O_APPEND   O_CREATE   O_TRUNC  等等常量，用|隔开
	//perm:  不同用户对文件的操作权限  （第二位：文件的主人，第三位：跟主人一个用户组，第四位，其他人）
	//有三种操作 读 = 4  写 = 2   执行 = 1    例：06(4+2)66  07(4+2+1)5(4+1)4
	if err !=nil{
		fmt.Println("文件打开失败")
		return
	}else{
		fmt.Println("文件打开成功")
	}
	defer file.Close()

}

func readByBufio() {
	file, err := os.Open("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json")
	if err !=nil{
		fmt.Println("文件打开失败")
		return
	}else{
		fmt.Println("文件打开成功")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n') //s中包括了换行符
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			fmt.Println("读取失败")
		} else {
			fmt.Print(s) //此处不用Println()，只要Print()
		}
	}
}

func writeByBufio() {
	//创写模式
	//file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/创写模式,.txt", os.O_RDWR|os.O_CREATE, 0644)
	//覆写模式
	//file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/覆写模式,.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	//追加模式
	file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/覆写模式,.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err !=nil{
		fmt.Println("文件创建失败")
		return
	}else{
		fmt.Println("文件创建成功")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("努力努力努力\n")
	writer.WriteString("专注专注专注\n")
	writer.WriteString("加油加油加油\n")
	writer.WriteString("嘿嘿嘿嘿嘿嘿\n")

	err = writer.Flush()
	if err !=nil{
		fmt.Println("文件写入失败")
	}else{
		fmt.Println("文件写入成功")
	}
}

func writeRead(){
	file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/覆写模式,.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err !=nil{
		fmt.Println("文件创建失败")
		return
	}else{
		fmt.Println("文件创建成功")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("第四轮写入\n")
	writer.WriteString("专注专注专注\n")
	writer.WriteString("加油加油加油\n")
	writer.WriteString("嘿嘿嘿嘿嘿嘿\n")

	err = writer.Flush()
	if err !=nil{
		fmt.Println("文件写入失败")
	}else{
		fmt.Println("文件写入成功")
	}


	//将文件“光标”（句柄/文件指针）拨回开头位置
	file.Seek(0,io.SeekStart)

	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n') //s中包括了换行符
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			fmt.Println("读取失败")
		} else {
			fmt.Print(s) //此处不用Println()，只要Print()
		}
	}
}

func writeByIoutil() {        //默认 os.O_WRONLY|os.O_CREATE|os.O_TRUNC 会直接覆盖掉
	type Bear struct{
		Name string
		Color color.RGBA
		Sex bool
	}

	mX := Bear{"一只未知的母熊熊",color.RGBA{23,98,0,56},false}
	err := ioutil.WriteFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json", []byte(fmt.Sprint(mX)), 0644)
	if err != nil{
		fmt.Println("写入文件失败")
	}else{
		fmt.Print("写入文件成功")
	}
}

func readByIoutil(){
	bytes, err := ioutil.ReadFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/熊熊们.json")
	//函数内封装了open和close 不需要手动关闭
	if err != nil{
		fmt.Println("读取失败")
	}else{
		fmt.Print(string(bytes))
	}
}

func isExist(filename string) (exists bool, info string) {
	fileInfo, err := os.Stat(filename)
	if fileInfo != nil && err == nil {
		fmt.Printf("%s文件存在\n", filename)
		exists = true
	} else if os.IsNotExist(err) {
		fmt.Printf("%s文件不存在\n", filename)
		exists = false
	} else {
		fmt.Println("搞不清存不存在!")
		exists = false
		info = "发生了一些奇奇怪怪的事情..."
	}
	return
}