package main

import (
	"net"
	"fmt"
	"strings"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8898")
	CheckError(err)
	defer listener.Close()                //listener要断开连接


	for{
		conn, err := listener.Accept()
		CheckError(err)

		go processMsg(conn)
	}

}


func processMsg(conn net.Conn){
	buf := make([]byte,1024)
	defer conn.Close()                    //conn也要断开连接

	for {
		n, err := conn.Read(buf)
		CheckError(err)
		msg := string(buf[:n])
		if n!=0{
			fmt.Println("收到客户端消息",msg)

			if strings.Contains(msg,"钱"){
				conn.Write([]byte("fuck off"))
				break
			}
			_, err = conn.Write([]byte("已阅:" + msg))
			CheckError(err)
		}
	}

}


func CheckError(err error){
	if err != nil{
		fmt.Println("错误：",err.Error())
		os.Exit(1)
	}
}