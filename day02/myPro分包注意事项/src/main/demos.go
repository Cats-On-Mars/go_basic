package main

import (
	"util"
	"fmt"
	"time"
)

func demo1() {
	elapsedSeconds := int(time.Now().Unix())

	daySeconds := 3600 * 24
	elapsedDays := elapsedSeconds/daySeconds + 1      //一共过去多少天

	//计算有多少个闰年
	yearSeconds := 3600 * 24 * 365
	elapsedYears := elapsedSeconds / yearSeconds     //大致过去多少年

	countLeapYear := 0
	for i := 1970; i <= (1970 + elapsedYears); i++ {
		if util.IsLeapYear(i) {
			countLeapYear = countLeapYear + 1 //获得1970至今的闰年数
		}
	}
	//计算年份
	elapsedYears = (elapsedDays-(366*countLeapYear))/365 + countLeapYear    //精确计算过去了多少年
	year := 1970 + elapsedYears
	fmt.Printf("今年是%d年\n", year)

	//计算月份和日期
	days := elapsedDays - (366 * countLeapYear) - (365 * (elapsedYears - countLeapYear))   //精确计算今年过去了多少天
	fmt.Printf("今年过去了%d天\n", days)
	month, date := util.GetDate(year, days)     //获得今天的月份和日期
	fmt.Printf("今天是%d月%d日\n", month, date)
}

func demo2(){
	fmt.Println("人家是呆猫2")
}

func demo3(){
	fmt.Println("人家是呆萌3")
}

