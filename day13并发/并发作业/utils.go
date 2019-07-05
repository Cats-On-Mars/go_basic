package 并发作业

import (
	"math/rand"
	"time"
)

func GenerateNum(start,end int) (n int){
		numrand := rand.New(rand.NewSource(time.Now().UnixNano()))
		n = start + numrand.Intn(end-start)
		return
}

func IsOdd(n int) bool{
	if n%2 == 1{
		return true
	}else{
		return false
	}
}



//根据规则来统计数量
func ClassifiedCount(data []interface{},rule func(interface{})bool) map[bool]int{
	trueFalseMap := make(map[bool]int)
	trueFalseMap[true] = 0
	trueFalseMap[false] = 0

	for _,v := range data{
		trueFalseMap[rule(v)]++
	}
	return trueFalseMap



}

func oddEvenCount(data []int) map[bool]int{
	oddEvenMap := make(map[bool]int)
	oddEvenMap[true] = 0
	oddEvenMap[false] = 0

	for _,v := range data{
		oddEvenMap[IsOdd(v)]++
	}
	return oddEvenMap
}
