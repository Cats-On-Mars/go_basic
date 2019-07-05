package main

import "fmt"

func soManyParams(a... interface{}){
	//fmt.Println(a)
	//for index,value := range a{
	//	fmt.Println(index,value)
	//}
	for i,x := range a{
		fmt.Println(i,x)
	}
}

func selfInfo(name string,age int,fool bool,hobbys ...string){
	fmt.Printf("我是%s，我今年%d了，我是否愚蠢呢：%t\n",name,age,fool)
	for i,x := range hobbys{
		fmt.Printf("我的第%d个爱好：%s\n",i+1,x)
	}
}

func selfInfo2(name string,age int,fool bool,hobbys ...interface{}){
	fmt.Printf("我是%s，我今年%d了，我是否愚蠢呢：%t\n",name,age,fool)
	for i,x := range hobbys{
		fmt.Printf("我的第%d个爱好：%v\n",i+1,x)
	}
}

func demo(){
	type Person struct{
		name string
		age int
		skill string
	}
	selfInfo2("张全蛋",30,true,"搬砖",1024,Person{"桑老师",30,"教学"})
}

func main() {
	soManyParams(213,"helloword",[3]int{2,4,66})
	selfInfo("张全蛋",30,true,"搬砖","敲代码","流水线","拍视频")
	demo()
}
