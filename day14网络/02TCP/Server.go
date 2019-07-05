package main

import (
	"fmt"
	"os"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8898")
	defer listener.Close()
	CheckError(err)

	for{
		conn, err := listener.Accept()
		CheckError(err)

		go processInfo(conn)
	}


}


func CheckError(err error){
	if err != nil{
		fmt.Println("错误：",err.Error())
		os.Exit(1)
	}
}


func processInfo(conn net.Conn){
	buf := make([]byte,1024)
	defer conn.Close()

	for{
		n, err := conn.Read(buf)
		CheckError(err)

		if n != 0{
			remoteAddr := conn.RemoteAddr()
			fmt.Println("客户端地址：",remoteAddr.String())
			fmt.Println("收到客户端消息：",string(buf[:n]))
			_, err = conn.Write([]byte("才发现无人等候"))
			CheckError(err)
			fmt.Println("发送反馈至客户端")

			if string(buf[:n])=="分手吧"{
				break
			}
		}
	}
}