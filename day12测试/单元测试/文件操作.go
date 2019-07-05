package 单元测试

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type color struct {
	Name string
	RGBAnum [4]int
}

func main() {

	var blue =color{"blue",[4]int{1,54,65,24}}
	truncfile(blue)

	var isBlue color
	loadfile(isBlue)
}

func truncfile(v interface{}) {
	file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day12测试/单元测试/文件操作.json",os.O_CREATE|os.O_RDWR|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(v)
	if err != nil{
		fmt.Println("写入json文件失败")
		return
	}
	fmt.Println("写入json文件成功")
}

func loadfile(v interface{}) {
	bytes, err := ioutil.ReadFile("/Users/shenhaihui/go/src/learn/goBiliblili/day12测试/单元测试/文件操作.json")
	if err != nil{
		fmt.Println("json文件读取失败",err)
		return
	}
	err = json.Unmarshal(bytes,&v)
	if err != nil{
		fmt.Println("json文件反序列化化失败",err)
	}
	fmt.Println(v)
}
