package game

import (
	"fmt"
	"math"
)

func IsPrimeNumber(num int) bool{

	if num <2{
		fmt.Println("请输入2以上的数字")
	}

	if num ==2 {
		return true
	}

	sqRoot := math.Sqrt(float64(num))      //从2到(n-1)     或者先求平方根，再2到(平方根+1)
	for i:=2;i<int(sqRoot)+1;i++{
		if num%i == 0{
			fmt.Println("发现因子",i)
			return false
		}
	}
	return true
}
