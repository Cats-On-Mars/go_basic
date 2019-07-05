package main

import "fmt"

type A interface{
	funcA1()
	funcA2()
}

type B interface{
	funcB1()
	funcB2()
}

type C interface{
	//显式地继承A类接口
	A

	//隐式地继承B类接口：没有明确地说继承B类，但事实上定义了其全部抽象方法
	funcB1()
	funcB2()

	//自有方法
	funcC()
}

func main() {
	var a A
	var b B
	var c C

	cobj := &Cobj{}
	a = cobj
	b = cobj
	c = cobj

	a.funcA1()
	a.funcA2()

	b.funcB1()
	b.funcB2()

	c.funcC()

}

type Cobj struct{

}

func (o *Cobj)funcA1(){
	fmt.Println("funcA1")
}
func (o *Cobj)funcA2(){
	fmt.Println("funcA2")
}
func (o *Cobj)funcB1(){
	fmt.Println("funcB1")
}
func (o *Cobj)funcB2(){
	fmt.Println("funcB2")
}
func (o *Cobj)funcC(){
	fmt.Println("funcC")
}
