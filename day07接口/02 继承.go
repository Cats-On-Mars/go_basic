package main

import "fmt"

/*
继承的意义：
①一行代码拥有父类的全部成员
②发展出新的属性和方法
③覆写父类的方法
*/

type Doggy struct{                              //父类
	name string
	sex bool
}

func (d *Doggy)bite(){                          //父类方法
	fmt.Printf("%s要咬你了啊\n",d.name)
}

type PoliceDog struct{							//子类
	Doggy                 //包含一个父类类型
	skill string								//子类新属性
}

func (pd *PoliceDog)doPoliceJob(){				//子类新方法
	fmt.Println("警犬新方法：缉毒")
}

func (pd *PoliceDog)bite(){						//子类覆写父类方法
	fmt.Printf("本警犬%s还没下嘴你就狗带了\n",pd.name)
}


func main() {
	//继承了父类所有的属性和方法
	pdog := new(PoliceDog)
	pdog.name = "战狼"
	pdog.sex = true
	pdog.bite()    //该方法已经被覆写。。。

	//发展出自身新的属性和方法
	pdog.skill = "徒手接RPG"
	fmt.Println(pdog.skill)
	pdog.doPoliceJob()


	//覆写父类的方法            //注意赋值方式
	pdog2 := PoliceDog{Doggy{"战狼2",true},"豪饮"}
	pdog2.bite()
}
