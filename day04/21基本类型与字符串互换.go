package main

import (
	"strconv"
	"fmt"
	"unsafe"
)

func main() {

	//将其他基本类型转换为字符串
	mstr1 := strconv.FormatBool(true)
	mstr2 := strconv.FormatInt(-123,2)   //base可以是 2，8，10，16
	mstr3 := strconv.FormatUint(123,2)   //base 代表将第一个参数（10进制的int）以什么进制输出
	mstr4 := strconv.FormatFloat(12987653.456,'g',5,32)
	mstr5 := strconv.FormatFloat(123987654.4563456,'G',5,32)
	//fmt有6种取值 'b' 'e', 'E', 'f', 'g',  'G'
	//当fmt为'f','e','E'时，prec代表浮点后的位数
	//当fmt为'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）时， prec代表总的长度
	//如果prec为1，表示适用最小数量、但又必须的数字来表示f
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入

	fmt.Println(mstr1)
	fmt.Println(mstr2)
	fmt.Println(mstr3)
	fmt.Println(mstr4)
	fmt.Println(mstr5)
	fmt.Println(unsafe.Sizeof(mstr4))


	//将字符串转化为基本类型
	var result,err interface{}
	result,err = strconv.ParseBool("1")  //"1", "t", "T", "true", "TRUE", "True"// "0", "f", "F", "false", "FALSE", "False"
	result,err = strconv.ParseInt("-12345",10,32)
	result,err = strconv.ParseUint("12345",10,16)
	result,err = strconv.ParseFloat("123.456",32)

	fmt.Println(result,err)


	//内建函数(builtin) 不用写包名  查中文文档 studygolang.com/pkgdoc
	//new()
}
