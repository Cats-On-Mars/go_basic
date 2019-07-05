package main

import "fmt"

func main() {
	type cat struct {
		name string
		lives int
	}
	c := cat{"Coding喵",1}

	/*
	·几乎所有占位符都可以前置宽度和左右对齐：10%s代表10宽度右对齐，-10%s代表10宽度左对齐
	·输出中带有小数位的，都可以指定其精度：.2%e
	*/

	//通用占位符
	fmt.Printf("%v\n",c)
	fmt.Printf("%+v\n",c)
	fmt.Printf("%#v\n",c)
	fmt.Printf("%T\n",c)
	fmt.Printf("%d%%\n",c.lives)

	//基本类型占位符
	fmt.Println("-----基本类型占位符-----")
	fmt.Printf("%d\n",123)//123，整型(decimal)
	fmt.Printf("%f\n",123.45)//123.450000，浮点型(float)
	fmt.Printf("%t\n",false)//false，布尔类型(true)
	fmt.Printf("%s\n","hello")//hello，字符串类型(string)
	fmt.Printf("%q\n","hello")//"hello"，带引号字符串类型(quotation)
	fmt.Printf("%p\n",&c)//0xc04205c3e0，指针类型(pointer)

	//进制占位符
	fmt.Println("-----进制占位符-----")
	fmt.Printf("%b\n",123)//1111011，二进制(binary)
	fmt.Printf("%o\n",123)//173，八进制进制(octal)
	fmt.Printf("%d\n",123)//123，十进制(decimal)
	fmt.Printf("%x\n",123)//7b，十六进制小写(hex)
	fmt.Printf("%X\n",123)//7B，十进制大写(hex)

	//科学计数法
	fmt.Println("-----科学计数法-----")
	fmt.Printf("%e\n",123456789012345.0)//1.234568e+14
	fmt.Printf("%e\n",0.00000000012345)//1.234500e-10
	fmt.Printf("%E\n",123456789012345.0)//1.234568E+14
	fmt.Printf("%E\n",0.00000000012345)//1.234500E-10
	fmt.Printf("%g\n",12345.0)//12345 智能选择【常规浮点数】或【科学计数法】显示
	fmt.Printf("%g\n",123456789012345.0)//1.23456789012345e+14
	fmt.Printf("%G\n",12345.0)//12345
	fmt.Printf("%G\n",0.0000000012345)//1.2345E-09

	//字符集占位符
	fmt.Println("-----字符集占位符-----")
	fmt.Printf("%c\n",123)//{ Unicode字符的【字符形式】
	fmt.Printf("%U\n",123)//U+007B Unicode字符的【十六进制序号形式】
	fmt.Printf("%c\n",'{')//{ Unicode字符的【字符形式】
	fmt.Printf("%U\n",'{')//U+007B Unicode字符的【字符形式】

	//设置输出的宽度和对齐方式
	fmt.Println("-----宽度和对齐-----")
	fmt.Printf("%20v\n",c)
	fmt.Printf("%20d\n",123)
	fmt.Printf("%20b\n",123)
	fmt.Printf("%20e\n",123456789012345.0)
	fmt.Printf("%20c\n",123)
	fmt.Printf("%-20v\n",c)
	fmt.Printf("%-20d\n",123)
	fmt.Printf("%-20b\n",123)
	fmt.Printf("%-20e\n",123456789012345.0)
	fmt.Printf("%-20c\n",123)

	//设置输出的小数精度
	fmt.Println("-----小数精度-----")
	fmt.Printf("%.2f%%\n",12.34)//12.340000%
	fmt.Printf("%.2f\n",123.45)
	fmt.Printf("%.2e\n",123456789012345.0)//1.234568e+14
	fmt.Printf("%.2e\n",0.00000000012345)//1.234500e-10
	fmt.Printf("%.2E\n",123456789012345.0)//1.234568E+14
	fmt.Printf("%.2E\n",0.00000000012345)//1.234500E-10

	//g/G的精度代表数字的总位数
	fmt.Printf("%.3g\n",12345.0)//12345 智能选择【常规浮点数】或【科学计数法】显示
	fmt.Printf("%.3g\n",123456789012345.0)//1.23456789012345e+14
	fmt.Printf("%.3G\n",12345.0)//12345
	fmt.Printf("%.3G\n",0.0000000012345)//1.2345E-09
}
