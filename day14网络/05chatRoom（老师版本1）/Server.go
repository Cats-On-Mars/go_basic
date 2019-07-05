package main

import (
	"net"
	"fmt"
	"os"
	"strings"
	"io"
)

/*
单撩
群撩
上线通知
*/

var (
	clientsMap = make(map[string]net.Conn)
)

func SHandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
		os.Exit(1)
	}
}

func main() {

	//建立服务端监听
	listener, e := net.Listen("tcp", "127.0.0.1:8888")
	SHandleError(e, "net.Listen")
	defer func() {
		for _, conn := range clientsMap {
			conn.Write([]byte("all:服务器进入维护状态，大家都洗洗睡吧！"))
		}
		listener.Close()
	}()

	for {
		//循环接入所有女朋友
		conn, e := listener.Accept()
		SHandleError(e, "listener.Accept")
		clientAddr := conn.RemoteAddr()
		fmt.Println(clientAddr.String()+"上线了")

		//给已经在线的用户发送上线通知
		for _,c := range clientsMap{
			c.Write([]byte(clientAddr.String()+"上线了"))
		}

		//将每一个女朋友丢入map
		clientsMap[clientAddr.String()] = conn

		//在单独的协程中与每一个具体的女朋友聊天
		go ioWithConn(conn)
	}

	//设置优雅退出逻辑

}

func ioWithConn(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	buffer := make([]byte, 1024)

	for{
		n, err := conn.Read(buffer)
		if err != io.EOF {
			SHandleError(err, "conn.Read")
		}

		if n > 0 {
			msg := string(buffer[:n])
			fmt.Printf("%s:%s\n", clientAddr, msg)

			strs := strings.Split(msg, "#")
			if len(strs) > 1 {
				targetAddr := strs[0]
				targetMsg := strs[1]

				if targetAddr == "all" {
					//群发消息
					for _, conn := range clientsMap {
						conn.Write([]byte(clientAddr + ":" + targetMsg))
					}
				} else {
					//点对点消息
					for addr, conn := range clientsMap {
						if addr == targetAddr {
							conn.Write([]byte(clientAddr + ":" + targetMsg))
							break
						}
					}
				}
			}else{

				//客户端主动下线
				if msg=="exit"{
					//将当前客户端从在线用户中除名
					//向其他用户发送下线通知
					for a,c := range clientsMap{
						if c == conn{
							delete(clientsMap,a)
						}else{
							c.Write([]byte(a+"下线了"))
						}
					}
				}else{
					conn.Write([]byte("已阅："+msg))
				}

			}

		}
	}


}
