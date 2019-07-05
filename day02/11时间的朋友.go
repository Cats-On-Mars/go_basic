package main

import (
	"time"
	"fmt"
)

func main() {

	//demo1()

	//Now返回的是本地化的时间 local time
	//UTC universal time coordinated 国际协调时间
	//北京时间比UTC快8小时
	//now := time.Now()
	//now := time.Date(2019,time.January,31,7,56,30,0,time.Local)  //时区问题
	//elapsedSeconds := int(now.Unix())

	elapsedSeconds := int(time.Now().Unix())


	daySeconds := 3600*24
	elapsedDays := elapsedSeconds/daySeconds+1            //一共过去多少天

	//计算有多少个闰年
	yearSeconds := 3600*24*365
	elapsedYears := elapsedSeconds/yearSeconds           //大致过去多少年
	countLeapYear := 0
	for i:=1970;i<=(1970+elapsedYears);i++{
		if isLeapYear(i){
			countLeapYear = countLeapYear+1              //获得1970至今的闰年数
		}
	}

	//计算年份
	elapsedYears = (elapsedDays-(366*countLeapYear))/365+countLeapYear    //精确计算过去了多少年
	year := 1970+elapsedYears
	fmt.Printf("今年是%d年",year)

	//计算月份和日期
	days := elapsedDays-(366*countLeapYear)-(365*(elapsedYears-countLeapYear))   //精确计算今年过去了多少天
	fmt.Printf("今年过去了%d天\n",days)
	month,date := getDate(year,days)   //获得今天的月份和日期
	fmt.Printf("今天是%d月%d日",month,date)

	/*//计算月份
		var fDays int
		if isLeapYear(year) {        //今年是闰年吗
			fDays = 29
		}else{
			fDays = 28
		}*/
    /*
	month := 0
	for i:=1;days>=0;i++{
		if i<8 && i%2!=0{
			days = days-31
		}else if i==2{
			days = days-fDays
		}else if i<8&&i%2==0{
			days = days-30
		}else if i>=8 && i%2 ==0{
			days = days-31
		}else if i>8 && i%2 !=0{
			days = days-30
		}else{
			panic("错误")
		}
		month = month+1
		if days<0{
			if i<8 && i%2!=0{
				days = days+31
			}else if i==2{
				days = days+fDays
			}else if i<8&&i%2==0{
				days = days+30
			}else if i>=8 && i%2 ==0{
				days = days+31
			}else if i>8 && i%2 !=0{
				days = days+30
			}else{
				panic("错误")
			}
			break
		}
	}
	fmt.Println(month,days)
    */
}

//输入今年的年份，和今年过去的天数，计算今天的月份和日期
func getDate(year,days int) (month,date int) {
	normalMonths := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapMonths := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var months [12]int
	if isLeapYear(year) {
		months = leapMonths
	} else {
		months = normalMonths
	}

	for i := 0; i < 12; i++ {
		days = days - months[i]
		if days <= 0 {
			month = i + 1
			date = days + months[i]
			//break
			//结束循环继续执行循环后的内容

			return
			//返回结果，函数结束
			//return month,date
		}
	}
	return
}


/*
const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)
*/
func isLeapYear(year int) bool{
	if year%4==0&&year%100!=0||year%400==0{
		return true
	}else{
		return false
	}
}

//误差太大
func demo1() {
	elapsedSeconds := int(time.Now().Unix())
	daySeconds := 3600 * 24
	yearSeconds := 3600 * 24 * 365
	monthSeconds := 3600 * 24 * 30
	elapsedYears := elapsedSeconds / yearSeconds
	fmt.Println(elapsedYears)
	elapsedMonths := (elapsedSeconds%yearSeconds)/monthSeconds + 1
	fmt.Println(elapsedMonths)
	elapsedDays := (elapsedSeconds%monthSeconds)/daySeconds + 1
	fmt.Println(elapsedDays)
}


