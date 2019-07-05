package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func main() {
	//createFile()
	copyFile()
}

func createFile() {
	file, _ := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/文件操作案例.txt", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("第二次\n")
	writer.WriteString("窗外的雨刚刚的停\n")
	writer.WriteString("午后气息浓浓的才散去\n")
	writer.WriteString("迷迷糊糊睁开眼\n")
	writer.WriteString("刚刚的梦我似乎在一瞬间看见你\n")
	writer.WriteString("oh my god 已经不知多久没想起\n")
	writer.WriteString("我淡淡地想着你\n")
	writer.WriteString("那年夏天最后的那一天\n")
	writer.WriteString("你轻轻地唱着歌\n")
	writer.WriteString("未曾感受的温柔模糊我的双眼\n")
	writer.WriteString("终于也可以开始一个人看明天\n")
	writer.Flush()
}

//文件拷贝
func copyFile(){

	//1使用ioutil
	//bytes, _ := ioutil.ReadFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/文件操作案例.txt")
	//_ = ioutil.WriteFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/文件操作案例copy.txt", bytes, 0754)

	//2使用io.copy
	/*file, _ := os.Open("/Users/shenhaihui/Downloads/extra/文字/屏幕快照 2019-01-12 上午12.55.58.png")
	defer file.Close()

	file2,err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/屏幕快照2 2019-01-12 上午12.55.58.png",os.O_CREATE|os.O_RDWR,0754)
	if err != nil{
		fmt.Println(err)
	}
	defer file2.Close()

	written, err := io.Copy(file2, file)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(written)
	}*/

	//3使用bufio
	srcfile, err := os.OpenFile("/Users/shenhaihui/Downloads/extra/文字/屏幕快照 2019-01-12 上午12.55.58.png", os.O_RDWR, 0644)
	if err != nil{
		fmt.Println("源文件打开失败：",err)
		return
	}
	defer srcfile.Close()

	dstfile,err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day09/人类一定要珍惜自己的牵绊.png",os.O_CREATE|os.O_RDWR,0644)
	if err != nil{
		fmt.Println("目标文件创建失败：",err)
		return
	}
	defer dstfile.Close()

	reader := bufio.NewReader(srcfile)
	writer := bufio.NewWriter(dstfile)
/*	for {
		b, err := reader.ReadByte()
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Println("源文件读取失败")
			return
		}
		err = writer.WriteByte(b)
		if err != nil{
			fmt.Println("目标文件写入失败")
			return
		}
	}
	writer.Flush()*/

	buf := make([]byte,8)
	for{
		n, err := reader.Read(buf)
		if n==0 && err==io.EOF{
			break
		}
		writer.Write(buf)
		writer.Flush()
	}

	//png.Decode(dstfile)  ???
	fmt.Println("COPY成功！！")
}
