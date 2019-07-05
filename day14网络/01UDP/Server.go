package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8848")
	checkError(err)

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)
	defer conn.Close()

	recvUdpMsg(conn)
}


func checkError(err error){
	if err != nil{
		fmt.Println("错误：",err.Error())
		os.Exit(1)
	}
}

func recvUdpMsg(conn *net.UDPConn){
	buf := make([]byte,30)
	n, addr, err := conn.ReadFromUDP(buf)
	checkError(err)
	fmt.Println("收到客户端信息：",string(buf[:n]))

	_, err = conn.WriteToUDP([]byte("思念是紧跟着的好不了的咳"), addr)
	checkError(err)
	fmt.Println("发送反馈至客户端")
}