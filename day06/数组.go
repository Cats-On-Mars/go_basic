package main

import "fmt"

//改
//查

func main() {
	//demo71()
	//demo72()
	//demo73()
	var a [5]int = [5]int{0,1}
	fmt.Printf("%p,%p\n",&a,&a[0])

	a[2]=233
	a[3]=38
	fmt.Printf("%p,%p\n",&a,&a[0])
}

//访问数元素+数组遍历
func demo73() {
	a1 := [...]int{0, 1, 2, 3, 4, 5, 666, 7, 8, 9}
	//按照下标访问元素
	fmt.Println(a1[6])
	fmt.Printf("type=%T,len=%d\n", a1, len(a1))
	for i := 0; i < len(a1); i++ {
		fmt.Println("下标=", i, "值=", a1[i])
	}
	for index, value := range a1 {
		fmt.Println(index, value)
	}
}

//修改数组中的元素
func demo72() {
	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1[6] = 666
	fmt.Println(a1)
}

//创建数组
func demo71() {
	var a1 [10]int
	var a2 [10]int = [10]int{}
	var a3 = [10]int{0, 1, 2, 3, 4, 5}
	var a4 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a1, a2, a3, a4, a5)
}