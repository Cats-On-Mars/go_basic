package main

import (
	"time"
	"fmt"
)

//判断year是否闰年
func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d年是闰年\n", year)
		return true
	} else {
		return false
	}
}

//求起止年份之间有多少个闰年
//startYear, endYear int 两个整型参数
func getLeapYears(startYear, endYear int) (leapYears int) {
	for i := startYear; i < endYear+1; i++ {
		if isLeapYear(i) {
			//leapYears=leapYears+1
			leapYears++
		}
	}
	return
}

//输入今天是今年的第几天，得到当前日期
func getDate(thisYear, thisYearDays int) (month, date int) {
	monthDaysNormal := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	monthDaysLeap := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var monthDays [12]int
	if isLeapYear(thisYear) {
		monthDays = monthDaysLeap
	} else {
		monthDays = monthDaysNormal
	}

	//将monthDays中的天数逐个相加，一旦大于thisYearDays，代表当前元素就是月份，大多少就是多少号，正好相等则说明是这个月最后一天
	var temp int
	for i := 0; i < 12; i++ {
		//temp = temp + monthDays[i]
		temp += monthDays[i]

		if temp >= thisYearDays{
			month = i+1
			date = thisYearDays-(temp-monthDays[i])

			//结束循环继续执行循环后的内容
			//break
			//返回结果，函数结束
			//return month,date
			return
		}

	}

	return 0, 0
}

func main() {

	daySecondes := 3600 * 24
	yearSeconds := int(3600 * 24 * 365.2425)

	//demo41()

	//Now()返回的是本地化的时间local time
	//UTC Universal Time Coordinated国际协调时间
	//北京时间比UTC快8小时
	//now := time.Now()
	now := time.Date(2018,time.July,26,7,0,0,0,time.UTC)
	elapsedSeconds := int(now.Unix())

	//算出总共逝去了多少天
	elapsedDays := elapsedSeconds/daySecondes + 1
	fmt.Println("elapsedDays=", elapsedDays)

	/*
	算出今年过去了多少天{
	·逝去了多少年
	·逝去的整年中一共包含多少天：365*n+闰年天数
	·求出今年过去的天数
	}
	*/
	//逝去了多少年
	elapsedYears := elapsedSeconds / yearSeconds
	fmt.Println("elapsedYears=", elapsedYears)

	//今年以前一共包含多少天
	thisYear := 1970 + elapsedYears
	beforeThisYearDays := 365*elapsedYears + getLeapYears(1970, thisYear-1)
	//求出今年过去的天数
	thisYearDays := elapsedDays - beforeThisYearDays
	fmt.Println("thisYearDays=", thisYearDays)

	//
	month, date := getDate(thisYear, thisYearDays)
	fmt.Printf("当前是%d年%d月%d日", thisYear, month, date)

}

//误差太大
func demo41() {
	elapsedSeconds := int(time.Now().Unix())
	daySeconds := 3600 * 24
	monthSeconds := 3600 * 24 * 30
	yearSeconds := 3600 * 24 * 365
	elapsedYears := elapsedSeconds / yearSeconds
	fmt.Println("elapsedYears=", elapsedYears)
	elapsedMonths := (elapsedSeconds%yearSeconds)/monthSeconds + 1
	fmt.Println("elapsedMonths=", elapsedMonths)
	elapsedDays := (elapsedSeconds%monthSeconds)/daySeconds + 1
	fmt.Println("elapsedDays=", elapsedDays)
}
