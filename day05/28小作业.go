package main

import (
	"game"
	"fmt"
)

//???问题：如果输入非"退散"的字符串，
//switch 变量的类型 {}   这样的表达式可以吗
//使用类型断言！！！！
//使用reflect.Typeof{}

func main() {
	//猜数字
	game.GuessNumber()

	//是否质数
	var num int
	for {
		fmt.Print("请输入数字，查看是否为质数：")
		fmt.Scan(&num)
		fmt.Printf("%d %t\n\n",num, game.IsPrimeNumber(num))
	}

	//福布斯排行榜
	game.Forbes()


}



