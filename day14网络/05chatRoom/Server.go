package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

/*
需求：
I D： 通过昵称识别  -name
在线：上下线通知  客户端优雅的退出机制：exit
     获取当前聊天室所有在线的人/单独查询某某是否在线 online@all  online@xx
聊天：单聊 xx# /群聊 all#  若发送内容不包含#则回复"已阅xx"   客户端直接发送单聊时若对方不在线就回复通知对方不在线
调天记录：查询全部 通过时间/对象查询  log@all log@xx     xx:/xx#
*/

//客户端结构体
type Client struct{
	Nickname string
	Addr net.Addr
	Conn net.Conn
	file os.File
}

//定义并初始化map【昵称】客户端结构体
var ClientMap = map[string]Client{}

func main() {

	//建立监听
	listener, e := net.Listen("tcp", "127.0.0.1:8888")
	checkError(e,"net.Listen")
	defer listener.Close()
	fmt.Println("服务器等待中。。。")


	//循环接入
	buf := make([]byte,10)
	for{
		conn, e := listener.Accept()
		checkError(e,"listener.Accept")
		defer conn.Close()


		//读取内容（昵称）上线通知 写入map  打印当前在线
		n, e := conn.Read(buf)
		checkError(e,"conn.Read")

		nickname := string(buf[:n])
		for _,client := range ClientMap{
			client.Conn.Write([]byte(nickname+"上线了"))
		}

		client := Client{Nickname:nickname,Addr:conn.RemoteAddr(),Conn:conn}
		ClientMap[nickname]=client

		fmt.Println("当前在线\n----------")
		for key := range ClientMap{
			fmt.Println(key)
		}


		//并发读取写入  每接入一个客户端就开一条协程 所以该协程开在for循环中
		go ioWithClient(client)
	}
}


//与客户端进行交互
func ioWithClient(client Client){
	buf := make([]byte,1024)
	//循环IO
	for{
		//从客户端读取消息
		n, err := client.Conn.Read(buf)
		checkError(err,"client.Conn.Read")


		//将消息存入聊天记录中
		msg := string(buf[:n])
		if n>0{
			write2log(msg)
		}


		//分析消息并回应
		if msg == "exit"{
			if client,ok := ClientMap[client.Nickname];ok{
				delete(ClientMap,client.Nickname)
			}
			client.Conn.Close()
		}else if strings.Contains(msg,"@"){
			msgSlice := strings.Split(msg, "@")
			switch msgSlice[0]{
			case "online":
				if msgSlice[1] == "all"{
					client.Conn.Write()
				}else{

				}

			case "log":
				if msgSlice[1] == "all"{
					client.Conn.Write()
				}else{

				}

			}
		}else if strings.Contains(msg,"#") {
			msgSlice := strings.Split(msg, "#")
			if msgSlice[0] == "all" {

			}else{
				
			}

		}

	}
}


//检查错误函数
func checkError(err error,why string){
	if err != nil{
		fmt.Printf("%sError:%v",why,err)
	}
}


//生成log文件
func write2log(msg string){
	file, e := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day14网络/05chatRoom/log文件",
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	checkError(e,"os.OpenFile")
	defer file.Close()

	file.WriteString(msg)
}

//分析消息并回应
func write2conn(){

}

//读出log文件写入conn


//是否在线函数

//check错误函数
