package util

//输入今年的年份，和今年过去的天数，计算今天的月份和日期
func GetDate(year,days int) (month,date int) {
	normalMonths := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapMonths := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var months [12]int
	if IsLeapYear(year) {
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

func IsLeapYear(year int) bool{
	if year%4==0&&year%100!=0||year%400==0{
		return true
	}else{
		return false
	}
}
