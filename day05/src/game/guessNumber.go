package game

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func GuessNumber(){
	fmt.Println("游戏开始，输入【退散】结束游戏")
	compRand := rand.New(rand.NewSource(time.Now().Unix()))
	compNum := compRand.Intn(1000)

	var userNum string

	for {
		fmt.Print("请输入数字：")
		fmt.Scan(&userNum)
		if userNum == "退散"{
			fmt.Printf("随机数为%d\nGAME OVER",compNum)
			return
		}
		ret,_ := strconv.Atoi(userNum)
		if ret > compNum{
			fmt.Println("猜大了\n")
		}else if ret < compNum{
			fmt.Println("猜小了\n")
		}else{
			fmt.Println("猜对啦！")
			return
		}
	}
}

