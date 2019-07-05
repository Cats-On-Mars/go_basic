package main

import (
	"fmt"
)

/*
·结合命令行参数、复合类型、包管理、排序等知识，写一个商品管理系统
·商品属性包括：整型商品ID（gid）、字符串型名称(name)、整型类别(cid)、浮点型价格(price)
·自己造一些假数据，实现如下功能：
·shop.exe -cmd single -gid xxx 查看商品ID为gid的单品详细信息
·shop.exe -cmd category -cid xxx -order 0 查看指定类别下的所有商品，cid=类别ID，order为排序规则
·排序规则参数：0=按gid升序，1=按gid降序，2=按price升序，3=按price降序
*/

//所有商品
var allGoods []*Goods
//按gid查询商品
var allGoodsMap map[int]*Goods
//按类别查询商品
var cateGoodsMap map[int][]*Goods

/*
包的初始化函数
会在入口函数之前被执行
写在包内任何一个go中都是可以的
*/
func init() {
	//自己造一些假数据
	//初始化全部商品
	allGoods = make([]*Goods, 0)

	//初始化全部商品Map
	allGoodsMap = make(map[int]*Goods)

	//初始化分类商品map，为每类商品创建空切片（将来方便动态地添加新的同类商品）
	cateGoodsMap = make(map[int][]*Goods)
	cateGoodsMap[Toy] = []*Goods{}
	cateGoodsMap[Weapon] = []*Goods{}

	/*	//麻烦死了，垃圾
		gPtr := new(Goods)
		gPtr.name = "滋水枪"
		gPtr.gid = 123456
		gPtr.cid = Toy
		gPtr.price = 50
		allGoods = append(allGoods,gPtr)

		//gid要硬编码，如果万一修改要修改很多处，复用率极低，垃圾
		allGoods = append(allGoods,&Goods{12346,"橡皮泥",Toy,5})*/

	//使用工厂模式自动生产Goods
	// 全自动生成gid，
	// 如果要修改产品生成逻辑，只需要修改一处，代码复用率高
	// 真厉害
	allGoods = append(allGoods, CreateGoods("滋水枪", Toy, 50))
	allGoods = append(allGoods, CreateGoods("阴沟里的泥巴", Toy, 5))
	allGoods = append(allGoods, CreateGoods("弹球", Toy, 10))
	allGoods = append(allGoods, CreateGoods("皮球", Toy, 100))
	allGoods = append(allGoods, CreateGoods("双色球", Toy, 10e6))

	//武器
	allGoods = append(allGoods, CreateGoods("绷弓子", Weapon, 3))
	allGoods = append(allGoods, CreateGoods("板砖", Weapon, 1))
	allGoods = append(allGoods, CreateGoods("平底锅", Weapon, 45))
	allGoods = append(allGoods, CreateGoods("板凳", Weapon, 23))
	allGoods = append(allGoods, CreateGoods("菜刀", Weapon, 99))

	//打印所有商品
	for i, v := range allGoods {
		fmt.Println(i, *v)
	}
}

func main() {

	//为每个命令创建一个三长度的数组，分别放置其[名称 默认值 用法说明]
	cmdInfoArray := [3]interface{}{"cmd", "未知命令", "你想干嘛"}
	gidInfoArray := [3]interface{}{"gid", 0, "要查询的商品ID"}
	cidInfoArray := [3]interface{}{"cid", 0, "要查询的类型ID"}
	orderInfoArray := [3]interface{}{"order", 0, "排序规则参数：0=按gid升序，1=按gid降序，2=按price升序，3=按price降序"}
	cmdlineArgValuesMap := getCmdlineArgs(cmdInfoArray, gidInfoArray, cidInfoArray, orderInfoArray)
	fmt.Println(cmdlineArgValuesMap)

	//分类实现不同需求
	switch cmdlineArgValuesMap["cmd"] {

	//·shop.exe -cmd single -gid xxx 查看商品ID为gid的单品详细信息
	case "single":
		//查看单品
		//校验用户是否输入了gid
		if value, ok := cmdlineArgValuesMap["gid"]; ok {
			if intValue, ok := value.(int); ok {
				gPtr := allGoodsMap[intValue]
				fmt.Printf("%-10d%-10s\t\t%-10d%10.2f\n", gPtr.gid, gPtr.name, gPtr.cid, gPtr.price)
			} else {
				fmt.Println("gid必须为整数")
			}
		}

		//·shop.exe -cmd category -cid xxx -order 0 查看指定类别下的所有商品，cid=类别ID，order为排序规则
	case "category":
		//查看指定分类下的全部单品
		cidValue := cmdlineArgValuesMap["cid"]
		if intValue, ok := cidValue.(int); ok {
			cateGoods := cateGoodsMap[intValue]

			//获取order参数
			if orderValue, ok := cmdlineArgValuesMap["order"]; ok {
				if orderIntValue, ok := orderValue.(int); ok {
					//按照orderIntValue定义的顺序对cateGoods进行排序
					sortGoods(cateGoods, orderIntValue)
				} else {
					fmt.Println("order必须为整数")
				}
			}

			for i, gPtr := range cateGoods {
				fmt.Printf("%-10d%-10d%-10s\t\t%-10d%10.2f\n", i, gPtr.gid, gPtr.name, gPtr.cid, gPtr.price)
			}
		}

	default:
		//用户输入错误
		fmt.Println("xjbinput,fuck off!")
	}

}
