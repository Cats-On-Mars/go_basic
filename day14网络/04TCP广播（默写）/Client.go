package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)




//缺少优雅退出逻辑




func main() {

	//拨号连接
	conn, err := net.Dial("tcp", "127.0.0.1:8898")
	if err != nil{
		fmt.Println("DailError:",err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("已连接服务器")


	//并发执行向服务器发送消息
	go sendMessage(conn)


	//循环从服务器接收消息
	buf := make([]byte,1024)
	for{
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read From Server Error:",err)
		}
		fmt.Println("收到服务端消息：",string(buf[:n]))
	}


	//fmt.Println("客户端关闭")

}

func sendMessage(conn net.Conn) {

	//创建读取器 从命令行读取内容
	reader := bufio.NewReader(os.Stdin)

	//循环读取内容
	for{
		data, _, _ := reader.ReadLine()

		input := string(data)                //键盘输入转化为字符串

		if input == "exit" {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

		_, err := conn.Write([]byte(input))   //输入写入字符串
		fmt.Println("发送消息", input)
		if err != nil {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

	}
}
