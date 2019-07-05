package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"io"
)

var (
	chanQuit = make(chan bool, 0)
	conn     net.Conn
)

func CHandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)

		os.Exit(1)
	}
}

func main() {

	//拨号连接，获得connection
	var e error
	conn, e = net.Dial("tcp", "127.0.0.1:8888")
	CHandleError(e, "net.Dial")
	defer func() {
		conn.Close()
	}()

	//在一条独立的协程中输入，并发送消息
	go handleSend(conn)

	//在一条独立的协程中接收服务端消息
	go handleReceive(conn)

	//设置优雅退出逻辑
	<-chanQuit

}

func handleReceive(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != io.EOF {
			CHandleError(err, "conn.Read")
		}

		if n > 0 {
			msg := string(buffer[:n])
			fmt.Println(msg)
		}
	}

}

func handleSend(conn net.Conn) {

	reader := bufio.NewReader(os.Stdin)
	for {
		//读取标准输入
		lineBytes, _, _ := reader.ReadLine()

		//发送到服务端
		_, err := conn.Write(lineBytes)
		CHandleError(err, "conn.Write")

		//正常退出
		if string(lineBytes) == "exit" {
			os.Exit(0)
		}

	}
}

