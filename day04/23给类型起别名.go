package main

import "fmt"

/*
减少代码量
对于特定类型，赋予【指示意义】 例：type byte = uint8
*/

type SS string

type BB bool

type Scores [10]int

type person struct{
	name string
	age int
}

type P person

func main() {
	var mstr SS = "hello"
	fmt.Printf("类型是：%T,值是：%v\n",mstr,mstr)

	var isStupid BB = true
	fmt.Printf("类型是：%T,值是：%v\n",isStupid,isStupid)

	var examScores Scores = [10]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Printf("类型是：%T,值是：%v\n",examScores,examScores)

	var zhangsan P = P{"zhangsan",20}
	fmt.Printf("类型是：%T,值是：%v\n",zhangsan,zhangsan)
}
