package main

import "fmt"

func main() {
	getPanic()
	fmt.Println("异常解除，执行到我啦")    //因为发生异常函数正常返回了，所以之后的行能正常执行到33
}

func getPanic() {
	defer func() {							//无论是系统抛的还是自己定义的，recover都能接
		if err := recover(); err != nil {   //监听到异常，立马执行defer里的recover()，接住异常并放到定义的变量里，然后函数正常返回
			fmt.Println(err)				//recover返回的类型为interface{}
		}
	}()
	getArea(-5)  							//发生异常的行
	fmt.Println("啦啦啦这里执行不到哦")     //一旦发生异常，立马执行defer的函数，异常行之后的行都执行不到
}

func getArea(r float32) float32{
	if r<0{
		panic("异常:半径不能为负数")        //自己抛异常   抛出的类型为interface{}
	}
	return 3.14*r*r
}