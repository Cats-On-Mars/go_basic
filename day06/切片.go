package main

import (
	"fmt"
)

func main() {
	//demob1()
	//demob2()
	//demob3()
	//demob4()
	//demob5()
	//demob6()
	//demob7()
}

//遍历访问切片元素
func demob7() {
	s := make([]int, 0)
	s = append(s, 10, 20, 30, 40, 50)
	for index, value := range s {
		fmt.Println(index, value)
	}
}

//拷贝src切片，覆盖dst切片
func demob6() {
	//var dst = []int{1,2,3,4}
	//var src = []int{10,20,30,40}
	//var dst = []int{1,2,3,4}
	//var src = []int{10,20,30,40,50}
	var dst = []int{1, 2, 3, 4, 5, 6}
	var src = []int{10, 20, 30, 40}
	//src中的元素从头覆盖dst中的元素，返回受影响的元素个数
	n := copy(dst, src)
	fmt.Println(dst, n)
}

//切片和底层数组的关系：
//一开始切片引用底层数组的元素地址——切片和数组内容的修改会相互影响（地址引用）
//切片扩容到突破原有容量时，就拷贝内容到新的地址——此时就已经完全脱离了底层数组
func demob5() {
	var a [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//切片不是值拷贝，而是直接底层数组地址
	as1 := a[:5]
	as2 := as1[:]
	fmt.Println(as1, as2)
	fmt.Printf("%p,%p,%p\n", &a[0], &as1[0], &as2[0])
	//切片不是值拷贝，而是直接底层数组地址
	as1[0] = 123
	as2[1] = 38
	fmt.Println(a, as1, as2)
	//扩容突破了容量(cap)的上限，切片拷贝得到新的元素地址，脱离了原先的底层数组
	fmt.Printf("before:cap=%d,addr=%p,elem addr=%p\n", cap(as2), &as2, &as2[0])
	as2 = append(as2, 666, 777, 888,999,1000,1100,1100)
	fmt.Printf("after:cap=%d,addr=%p,elem addr=%p\n", cap(as2), &as2, &as2[0])
	as2[0] = 90000
	fmt.Println(a, as1, as2)
	a[0] = 0
	fmt.Println(a, as1, as2)
}

//截取数组获得切片
func demob4() {
	var a [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//在数组的基础上截取，得到切片
	as1 := a[2:8]
	fmt.Printf("type=%T,value=%v\n", as1, as1)
	//截取的方式[头下标:尾下标]，含头不含尾,不写头默认头=0，不写尾默截取到最后一位
	as2 := a[:8]
	as3 := a[2:]
	as4 := a[:]
	fmt.Println(as2, as3, as4)
}

//兼并新的切片
func demob3() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	s2 := []int{4, 5, 6}
	//追加s2中所有元素到s
	s = append(s, s2...)
	fmt.Println(s, s2)
}

//向切片中追加新的元素
func demob2() {
	s := make([]int, 0)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 2)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 3)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 4)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 5)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 6, 7, 8)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
	s = append(s, 9)
	fmt.Printf("value=%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

//创建切片
func demob1() {
	//var s1 []int = []int{0,1,2}
	//var s2 = []int{0,1,2}
	//s3 := []int{0,1,2}
	//fmt.Printf("type=%T,value=%v\n",s3,s3)
	s4 := make([]int, 3)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s4, s4, len(s4), cap(s4))
	s5 := make([]int, 4, 10)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s5, s5, len(s5), cap(s5))
}