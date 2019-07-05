package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8848")
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("爱恋不过是一场高烧"))
	checkError(err)
	fmt.Println("发送消息至服务端")

	buf := make([]byte,30)
	n, err := conn.Read(buf)
	checkError(err)
	fmt.Println("收到服务端反馈：",string(buf[:n]))

}



