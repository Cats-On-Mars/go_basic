package 并发作业

import (
	"testing"
	"time"
)

func TestGenerateNum(t *testing.T){

	ints := make(map[int]int)     		//利用map来计独特的数
	for i := 0; i < 1000; i++ {
		num := generateNum(100,1000)
		ints[num] = 1
		if num<100||num>1000{
			t.Error("出错")
			return
		}
		time.Sleep(time.Nanosecond)
	}


	if len(ints)<400{
		t.Error("ERROR：随机数重复率太高")
		return
	}else{
		t.Log("成功：随机数重复率低")
	}
}


/*func TestOddEvenCount(t *testing.T){
	var data []int
	for i := 1; i < 101; i++ {
		data = append(data,i)
	}

	count := oddEvenCount(data)
	if count[true] != 50||count[false]!=50{
		t.Error("出错")
	}
}*/


func TestClassifiedCount(t *testing.T) {
	var data []interface{}
	for i := 1; i < 101; i++ {
		data = append(data,i)
	}

	count := ClassifiedCount(data, func(n interface{})bool{
		return (n.(int)%2 == 1)

	})

	if count[true] != 50||count[false]!=50{
		t.Error("出错")
	}
}