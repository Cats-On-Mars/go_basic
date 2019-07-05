package main

import "fmt"

type Dog struct {
	name string
	sex int
}

func (d *Dog)SetName(name string){
	d.name = name
}

func (d *Dog)GetName(){                   //属性名是私有的，获取属性的方法是共有的 封装起来 防止更改
	//在这里可以加一些权限校验
	fmt.Printf("本汪名叫%s\n",d.name)
}

func (d *Dog)Bite(){
	//在这里可以加一些权限校验
	fmt.Printf("%s汪汪汪\n",d.name)
}

//创建实例的工厂方法
func createDog(name string,sex int) *Dog{
	dog := new(Dog)
	dog.name = name
	dog.sex = sex
	return dog
}

func main() {
	//先创建空白对象
	dog1 := Dog{}
	dog1.name = "臭皮"
	dog1.sex =1
	dog1.GetName()
	dog1.Bite()

	//直接赋值创建
	dog2 := Dog{"旺财",1}
	dog2.GetName()
	dog2.Bite()

	dog3 := Dog{sex:0} //有选择地赋值必须声明属性名
	dog3.SetName("来福")
	dog3.GetName()
	dog3.Bite()


	//通过new函数创建指针对象
	//通过【指针】访问结构体的【成员】（包括属性和方法）和通过【值】访问一模一样
	dog4Ptr := new(Dog)
	dog4Ptr = &Dog{name:"小白"}
	dog4Ptr.sex = 0
	dog4Ptr.GetName()
	dog4Ptr.Bite()

	//通过工厂方法创建狗对象
	dog5Ptr := createDog("小黑",1)
	dog5Ptr.GetName()
	dog5Ptr.Bite()
}
