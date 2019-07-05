package main

import "fmt"


/*
标识符：用于命名常量、变量、表达式的字符
标识符可以是：大小写字母，数字，_ (下划线)
数字不能开头
*/

const PI float32 = 3.14   //定义浮点型常量PI，一般为全大写
var r int             //定义整型变量r，用于存储半径
var area float32      //定义浮点型变量，用于存储面积

func Area() {

	r = 10            //给r赋值为10（将10写入r的内存地址）
	area = PI * float32(r) * float32(r)   //将r强制转换成浮点型

	//area = 10000  //重新赋值，变量可以反复擦写

	fmt.Println("你的脸有",area,"辣么大")
	fmt.Printf("你的脸有%f辣么大\n",area)
	fmt.Printf("你的脸有%.2f辣么大\n",area)
	fmt.Printf("你的脸有%10.2f辣么大\n",area)  //总位数10，小数点后2位 自动右对齐
	fmt.Printf("你的脸有%-10.2f辣么大\n\n",area)  //总位数10，小数点后2位 -号代表左对齐

}

//驼峰命名风格
func getFaceArea(r float32){
	result := PI * r *r  // := 声明、赋值合二为一，外加动态检测类型
	fmt.Printf("result的类型为%T\n",result)
	fmt.Printf("半径为%.2f的脸有%.2f大\n\n",r,result)
}

//小写加下划线风格
func get_face_area(r float32) float32{
	result := PI * r *r
	return result
}


func get_face_area_ii(r float32)(area float32){
	//area = PI * r *r   //area已经在函数声明中定义过了
	//return             //预定义了返回值名字，无需说明return谁
	return PI * r * r
}

func main() {

	Area()

	getFaceArea(5)

	ret := get_face_area(10)
	fmt.Printf("ret=%.2f\n\n",ret)

	fmt.Println(get_face_area_ii(20))
	//a := get_face_area_ii(20)
	//fmt.Println(a)

}
