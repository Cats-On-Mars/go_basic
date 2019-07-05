package game

import "flag"

func GetCmd()(cmd string, gid,cid,order int){

	flag.StringVar(&cmd,"cmd","","命令")
	flag.IntVar(&gid,"gid",0,"商品ID")
	flag.IntVar(&cid,"cid",0,"类别ID")
	flag.IntVar(&order,"order",0,"排序方式")

	flag.Parse()

	return
}