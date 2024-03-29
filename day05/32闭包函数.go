package main

import "fmt"

var index =0
var heros = [...]string{"卢员外","林冲","武松","鲁达","史文恭","关胜"}

func main() {
	for i:=0;i<10;i++{
		fmt.Println(getHero())
	}

}
func getHero() string {
	name := heros[index]
	if index>5{
		index = 0
	}
	index++
	return name
}

/*闭包函数：返回函数的函数
闭包的好处：【内层函数的状态】被保存在闭包中
不使用闭包，就要开辟多个全局变量来保存函数以外的数据
如果说这个函数被多方调用，大家都需要各保存各的数据，那么此时就需要开辟多个全局变量
具体使用哪个全局变量，还要在函数内做判断——增大了重复的代码量，令代码看起来比较垃圾*/

//使用多个全局变量保存多套副本的索引

//全局变量
var heros = [...]string{"关胜", "林冲", "秦明","呼延灼", "武松", "鲁达"}
//宋江的索引
var index1 = 0
//吴用的索引
var index2 = 0
//脑补卢员外的索引、柴进的索引...

func useNormal() {
	for i := 0; i < 10; i++ {
		fmt.Println(giveMeOne("宋江"))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(giveMeOne("吴用"))
	}
}

func giveMeOne(who string) string {
	var theOne = ""

	//差不多的东西写两遍，很垃圾
	//万一卢俊义也来带队，那就需要三个全局变量，三个if分支...
	if who == "宋江" {
		theOne = heros[index1]
		index1++
		if index1 > len(heros)-1 {
			index1 = 0
		}
	} else {
		theOne = heros[index2]
		index2++
		if index2 > len(heros)-1 {
			index2 = 0
		}
	}

	return theOne
}






//使用闭包函数搞

//全局变量
var heros = [...]string{"关胜", "林冲", "秦明","呼延灼", "武松", "鲁达"}

//使用函数闭包的案例
func useClosure() {
	//得到返回的闭包内函数
	fSongjiang := giveHimOne(0)
	fWuyong := giveHimOne(4)
	for i := 0; i < 10; i++ {
		fmt.Println("宋江线：", fSongjiang("黑子"))
		//time.Sleep(1000*time.Millisecond)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("吴用线：", fWuyong("大坏比"))
		//time.Sleep(500*time.Millisecond)
	}
	fmt.Println("吴用线：", "全军休息")
	fmt.Println(fWuyong("大坏比"))
	fmt.Println("宋江线：", fSongjiang("黑子"))
}

//闭包函数：返回函数的函数
func giveHimOne(start int) func(name string) string {
	//保存闭包系统内的状态
	var index int = start

	//内层函数
	return func(name string) string {
		theOne := heros[index]

		//状态被保存在外层的闭包中
		index++
		if index > len(heros)-1 {
			index = 0
		}
		return name + ":" + theOne
	}

}