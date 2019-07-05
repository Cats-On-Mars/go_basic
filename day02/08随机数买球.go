package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

		var team string
		var money int

	for {
		fmt.Println("请输入您要购买的金额和球队,以空格隔开，若退出请输入over")
		fmt.Scanf("%s %d",&team,&money)
		if team =="over"{
			break
		}

		fmt.Printf("\n您购买了%s队%d元，人生巅峰即将开始\n", team, money)


		time.Sleep(500 * time.Millisecond)
		fmt.Print("\n等待中")
		for i := 0; i < 6; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Print(". ")
			/*witch i%3{
		case 0:
			fmt.Print("/")
		case 1:
			fmt.Print("-")
		case 2:
			fmt.Print("\\\n")
		}*/
		}

		//luckyNumber := rand.Intn(100)
		// Intn()以int的形式返回[0,n)中的非负伪随机数  伪随机数就是概率确实为设定好的数字
		//time.Now().Unix()获取当前时间距离1970年零时逝去的秒数
		//rand.NewSource(time.Now().Unix())每秒更新一个随机数的种子
		//Unix()函数返回的是int64类型，要注意
		myRand := rand.New(rand.NewSource(time.Now().Unix()))
		luckyNumber := myRand.Intn(100)

		fmt.Printf("\n\nluckyNumber是:%d\n\n",luckyNumber)

		if luckyNumber >= 90 {
			fmt.Println("靠海别野欢迎你！\n")
		} else {
			fmt.Println("天台欢迎你!\n")
		}

		time.Sleep(2*time.Second)
	}
}


/*
type Rand struct {        //一个 Rand 是随机数的来源。
	src Source
	s64 Source64
	readVal int64
	readPos int8
}

func (r *Rand) Intn(n int) int {}

func New(src Source) *Rand {}

func NewSource(seed int64) Source {}


函数返回t作为Unix时间，即经过的秒数
从1970年1月1日起
func (t Time) Unix() int64 {return t.unixSec()}

func (t Time) UnixNano() int64 {return (t.unixSec())*1e9 + int64(t.nsec())}
*/