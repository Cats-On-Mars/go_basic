package main

import (
	"fmt"
	"time"
	"os"
)

/*
·命令行键入一行诗句启动应用：idiom.exe -cmd start -poem 大王派我来巡山
·将诗句中的每个字丢入【模糊查询管道】
·另外再建立【精确查询管道】和【结束管道】，分别存储【成语】（大鹏展翅、占山为王、龟派气功...）和【结束指令】（fuckoff）
·时钟每秒随机读入一条管道数据：
	如果是【模糊查询管道】：起协程进行模糊查询，并汇总数据在内存
	如果是【精确查询管道】：起协程进行精确查询，并汇总数据在内存
	如果是【结束指令】：停止查询，将内存中的数据持久化为json并退出；
*/

const DB_PATH = "d:/temp/idioms-v2.0.json"

var (
	//数据管道
	chanAmbiguous = make(chan string, 20)
	chanAccurate  = make(chan string, 20)
	chanQuit      = make(chan string, 0)

	//全局内存数据
	dbDataMap = make(map[string]Idiom)
)

func main0() {

	//读入命令行参数
	//idiom.exe -cmd start -poem 大王派我来巡山
	cmdInfo := [3]interface{}{"cmd","未知命令","你打算干什么"}
	poemInfo := [3]interface{}{"poem","绞尽果汁想不出","用于启动的一行诗句"}
	retValuesMap := GetCmdlineArgs(cmdInfo,poemInfo)
	cmd := retValuesMap["cmd"].(string)
	poem := retValuesMap["poem"].(string)
	fmt.Println(cmd, poem)

	//将读入的诗句打碎丢入模糊管道
	for _, v := range poem {
		keyword := fmt.Sprintf("%c", v)
		chanAmbiguous <- keyword
	}

	//三选一读入管道数据，周期性执行
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			select {
			case keyword := <-chanAmbiguous:
				go DoAmbiguousQuery(keyword,"1",chanAccurate)
			case keyword := <-chanAccurate:
				go DoAccurateQuery(keyword)

			case <-chanQuit:
				WriteIdioms2File(dbDataMap, DB_PATH)
				os.Exit(0)
			}
		}
	}()

	//定时20秒结束主程序
	timer := time.NewTimer(20 * time.Second)
	<-timer.C
	chanQuit <- "OVER"

}
