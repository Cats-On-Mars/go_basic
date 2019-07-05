package main

import "fmt"

func main() {
	//demo81()
	//demo82()
	//demo83()
	demo84()

}

//map的查询和遍历
func demo84() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5

	//没有校验的查询
	value1 := mmp["sex"]
	fmt.Println(value1)

	//带有校验的查询
	value2, ok := mmp["最爱的书"]
	if ok == true {
		fmt.Println(value2)
	} else {
		fmt.Println("东哥：洒家没有最爱的书，查你妹")
	}
	//遍历
	for key, value := range mmp {
		fmt.Println(key, value)
	}
}

//修改map
func demo83() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5
	//修改特定的键值对
	mmp["hobby"] = [...]string{"抽烟", "喝酒", "去浪"}
	mmp["rmb"] = -29999.5
	mmp["最爱的书籍"] = "平凡的世界"
	fmt.Println(mmp)
}

//删除键值对
func demo82() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5
	//删除hobby对应的键值对
	delete(mmp, "hobby")
	fmt.Println(mmp)
}

//创建映射
func demo81() {
	//var mmp = map[string]int{}
	//var mmp map[string]int
	//var mmp = map[string]int{"代码量":10000,"酒量":20}
	//通过内建函数make创建
	//mmp := make(map[string]int)

	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})

	//丢入键值对到map中
	mmp["name"]="西门东"
	mmp["sex"]="male"
	mmp["hobby"]="female"
	mmp["rmb"]=0.5

	fmt.Printf("type=%T,value=%v,len=%d\n", mmp, mmp, len(mmp))
}