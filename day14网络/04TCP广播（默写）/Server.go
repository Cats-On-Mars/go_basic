package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)




//TODO:缺少优雅退出逻辑




var adrConMap = make(map[string]net.Conn)
var ch = make(chan []byte,1024)

func main() {

	//建立网络监听
	listener, err := net.Listen("tcp", "127.0.0.1:8898")
	if err != nil{
		fmt.Println("listenError：",err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("服务端等待连接......")



	//并发执行 发送消息至客户端
	go writeToConn()



	//循环接入服务端 并接收消息
	for{
		conn, err := listener.Accept()         //接入客户端
		if err != nil{
			fmt.Println("acceptError：",err)
			os.Exit(1)
		}

		adr := conn.RemoteAddr().String()      //获取客户端地址
		adrConMap[adr]=conn                    //加入地址-连接map

		fmt.Println("客户端列表")
		fmt.Println("-------------")
		for key := range adrConMap{            //打印当前所有连接的客户端
			fmt.Println(key)
		}


		//并发执行 从客户端接收消息
		go readConn(conn)
	}
}


func readConn(conn net.Conn){

	//延时关闭网络io
	defer func() {
		conn.Close()
		fmt.Println("连接已关闭:",conn.RemoteAddr())
	}()

	buf := make([]byte,1024)


	//循环从conn中读取消息
	for{
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("ReadFromClientError：",err)
			os.Exit(1)
		}

		if n != 0 {                             //如果消息不为空才加入消息队列 为空则不处理
			ch <- buf[:n]
			fmt.Println("该条消息已加入队列...")
		}else{
			fmt.Println("读取客户端消息长度为0")
		}
	}
}


func writeToConn() {

	//循环写入消息至conn
	for{

		//从消息队列中取出内容
		message := <-ch
		infoSlice := strings.Split(string(message), "#")

		if len(infoSlice)>1{
			addr := infoSlice[0]
			addr = strings.TrimSpace(addr)
			info := infoSlice[1]

			//群发
			for adr,con := range adrConMap{
				_, err := con.Write([]byte(info))
				fmt.Println("服务器发送消息", info, "给", adr)
				if err != nil{
					fmt.Println("WriteToClientError：",err)

					os.Exit(1)
				}
			}

			//转发
			/*if conn,ok := adrConMap[addr];ok{
				_, err := conn.Write([]byte(info))
				fmt.Println("服务器转发消息", info,"给",addr)
				if err != nil{
					fmt.Println("WriteToClientError：",err)
					os.Exit(1)
				}
			}*/
		}
	}

}