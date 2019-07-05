package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8898")
	CheckError(err)
	defer conn.Close()

	buf := make([]byte,1024)
	reader := bufio.NewReader(os.Stdin)
	for{
		line, _, err := reader.ReadLine()
		CheckError(err)

		_, err = conn.Write(line)
		CheckError(err)
		fmt.Println("发送消息至服务端")

		n, err := conn.Read(buf)
		CheckError(err)
		if n != 0{
			msg := string(buf[:n])
			fmt.Println("收到服务端反馈:",msg)
			if msg == "fuckoff"{
				break
			}
		}
	}

}


func CheckError(err error){
	if err != nil{
		fmt.Println("错误：",err.Error())
		os.Exit(1)
	}
}