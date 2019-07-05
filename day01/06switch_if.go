package main

import "fmt"

func main() {
	demo1()
	diff()
	demo2()
}

func demo1() {
	var ret float32
	ret = getResult(5, 3, "+")
	fmt.Println("ret=", ret)
	ret = getResult(5, 3, "-")
	fmt.Println("ret=", ret)
	ret = getResult(5, 3, "*")
	fmt.Println("ret=", ret)
	ret = getResult2(5, 3, "/")
	fmt.Println("ret=", ret)
	ret = getResult2(5, 3, "%")
	fmt.Println("ret=", ret)
}

func demo2(){
	var a,b,ret float32
	var operator string
	fmt.Println("请输入式子,字符间用空格隔开")
	fmt.Scanf("%f%s%f",&a,&operator,&b)
	ret = getResult2(a,b,operator)
	fmt.Println(ret)
}

func getResult(a float32, b float32, operator string) (ret float32){
	/*
	//单分支
	if operator == "+"{
		ret = a+b
	}
	*/

	/*
	//双分支
	if operator == "+"{
		ret = a+b
	}else{
		fmt.Println("不支持的操作符：%s\n",operator)
	}
	*/

	//多分支
	if operator == "+"{
		ret = a+b
	}else if operator == "-"{
		ret = a-b
	}else if operator == "*"{
		ret = a*b
	}else if operator == "/"{
		ret = a/b
	}else if operator == "%"{
		ret = float32(int(a) % int(b))
	}else{
		fmt.Println("不支持的操作符：%s\n",operator)
	}

	return
}

func getResult2(a float32, b float32, operator string) (ret float32){
	//专做多分支
	//判断operator的取值
	//一定是一些孤立的值
	switch operator{
	case"+":
		ret = a+b
	case"-":
		ret = a-b
	case"*":
		ret = a*b
	case"/":
		ret = a/b
	case"%":
		ret = float32(int(a)%int(b))
		//如果不符合上述任何一种情形
	default:
		fmt.Printf("不支持的操作符%s\n",operator)
	}

	return
}

func diff(){
	/*
	if条件分支 VS switch条件分支
	if条件可以判断【连续的范围】，还支持条件与或非
	switch只适用于多个孤立的【点】，例如1，2，3或"a"，"b"，"c"或true，false
	*/
	fmt.Println("请输入您的考试成绩：")
	var score float32
	fmt.Scan(&score)

	if score >100 || score < 0{     //条件表达式的值是一个连续的范围，所以switch不适用
		fmt.Println("非法成绩")
	}else if score >= 90 && score <= 100{
		fmt.Println("天秀")
	}else if score > 59 && score < 90 {
		fmt.Println("陈独秀同学请你坐下")
	}else {
		fmt.Println("不及格")
	}
}