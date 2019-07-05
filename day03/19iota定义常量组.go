package main

import "fmt"

const(
	sunday = iota	//常量组中的iota会从0开始自动增长
	monday
	tuesday
)

const (
	E = 1<<iota     //常量组中的iota会从0开始自动增长
	D               //沿用第一个计算方式
	C
	B
	A
	S

)

const (
	Eq = 2*iota
	Dq
	Cq
	Bq
	Aq
	Sq

)

func main() {
	fmt.Println(sunday,monday,tuesday)
	fmt.Println(S,A,B,C,D,E)
	fmt.Println(Sq,Aq,Bq,Cq,Dq,Eq)
}
