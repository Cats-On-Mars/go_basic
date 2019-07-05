package main

import "fmt"


/*
#接口->父类实现->多种子类实现->多态
·定义接口IPerson，定义吃喝睡三个抽象方法；
·定义一个IPerson的实现类Worker即劳动者，拥有劳动方法Work()(output string)其中output是其工作产出，和休息方法Rest()；
·继承Worker实现三个不同职业的子类：程序员Coder、老师Teacher、农民Farmer，并创建一个Worker的集合；
·从命令行循环接收【今天星期几】的输入，如果是周一到周五全体工作，如果是周六日程序员工作其他人休息；
*/

//定义接口
type USBDevice interface{
	Connect() bool
	ReadData() []byte
	DisConnect() bool
}


func main() {
	var usbDevice USBDevice                //定义接口对象

	up := Up{8*1024*1024*1024}
	usbDevice = &up                        //用具体实现类去为接口赋值

	fmt.Println(usbDevice.Connect())
	fmt.Println(usbDevice.ReadData())
	fmt.Println(usbDevice.DisConnect())


}

//定义结构体
type Up struct{
	storage int
}

//实现接口里的方法
//实现了所有方法，就说这个结构体实现了这个接口
//（这个结构体类型的实例也是这个接口的实例，实例类型是这个接口）可以给接口类型的实例赋值
func (u *Up)Connect() bool{
	fmt.Println("U盘连接成功")
	return true
}

func (u *Up)ReadData() []byte{
	return []byte("U盘读取成功")
}

func (u *Up)DisConnect() bool{
	fmt.Println("U盘已断开成功")
	return true
}