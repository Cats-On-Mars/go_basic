package main

import "fmt"

//递归：自己调用自己
//注意事项：要有终止条件
func main() {
	//fmt.Println(getSum(10))

	//for i:=1;i<10 ;i++  {
	//	fmt.Println(i,getFb(i))
	//}

	fmt.Println(getFbSum(20))
}


func getSum(n int) int{
	if n==1{
		return 1
	}else{
		return n + getSum(n-1)
	}

}

func getFb(n int) int{
	if n==1 {
		return 1
	}else if n==2{
		return 1
	}else{
		return getFb(n-1) + getFb(n-2)
	}

}

func getFbSum(n int) int{
	if n==1{
		return 1
	}else{
		return getFb(n)
	}
}