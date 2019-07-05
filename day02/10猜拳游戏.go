package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guessFist()
}

func guessFist() {
	fmt.Println("游戏开始，结束请输入退散")
	fists := [3]string{"石头", "剪子", "布"}
	var userFist, computerFist string

	//OUTER:
	for{
		fmt.Print("请出石头、剪子、布：")
		fmt.Scanf("%s", &userFist)
		if userFist == "退散"{
			fmt.Println("GAME OVER")
			break
		}

		compRand := rand.New(rand.NewSource(time.Now().Unix()))
		cf := compRand.Intn(3)
		computerFist = fists[cf]
		fmt.Println("电脑出拳:",computerFist)


		switch userFist {
		//case "退散":
		//	fmt.Println("GAME OVER")
		//	break OUTER
		case "石头":
			if computerFist == "石头" {
				fmt.Println("平局")
			} else if computerFist == "剪子" {
				fmt.Println("大官人赢了")
			} else {
				fmt.Println("电脑赢了")
			}
		case "剪子":
			if computerFist == "剪子" {
				fmt.Println("平局")
			} else if computerFist == "布" {
				fmt.Println("大官人赢了")
			} else {
				fmt.Println("电脑赢了")
			}
		case "布":
			if computerFist == "布" {
				fmt.Println("平局")
			} else if computerFist == "石头" {
				fmt.Println("大官人赢了")
			} else {
				fmt.Println("电脑赢了")
			}
		default:
			fmt.Println("您的智商余额不足，请充值")
		}

		fmt.Println()
		time.Sleep(500*time.Millisecond)
	}

}
