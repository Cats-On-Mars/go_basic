package main

import (
	"fmt"
	"time"
)

func main() {
	//顺序执行
	fmt.Println("hello word")
	fmt.Println("golang")

	//选择结构
	if time.Now().Hour() % 2 == 0{
		fmt.Println("情绪稳定")
	}else{
		fmt.Println("情绪激动")
	}

	//循环
	i := 0
	for {
		if i > 10{
			//去到HAHA标记的地方
			goto HAHA
		}
		fmt.Println(i)
		time.Sleep(500*time.Millisecond)
		i++
	}

	//此处执行不到
	fmt.Println("此处领取靠海别野一套")
	fmt.Println("此处领取满汉全席一桌")
	fmt.Println("此处领取我厂生产的女朋友一个")

	HAHA:
		fmt.Println("GAME OVER")
}
