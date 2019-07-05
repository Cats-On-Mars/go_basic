package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8898")
	CheckError(err)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte,1024)

	for{
		line, _, err := reader.ReadLine()
		CheckError(err)
		_, err = conn.Write(line)
		CheckError(err)
		fmt.Println("发送信息至服务端")


		n, err := conn.Read(buf)
		CheckError(err)
		fmt.Println("收到服务端反馈：",string(buf[:n]))
	}

}

func CheckError(err error){
	if err != nil{
		fmt.Println("错误：",err.Error())
		os.Exit(1)
	}
}
