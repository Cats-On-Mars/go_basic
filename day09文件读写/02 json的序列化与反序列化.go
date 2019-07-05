package main

import (
	"fmt"
	"encoding/json"
)

type Hamster struct{
	//Species string
	Variety int
	Name string
	Sex bool
}

const(
	BEAR = iota
	AlINE
	THREELINE
	Roborovsk
)

var (
	dm,cm,lgg Hamster
	hs []Hamster
	hMap map[string]Hamster
)

func init(){
	dm = Hamster{BEAR,"短毛",true}
	cm = Hamster{BEAR,"长毛",true}
	lgg = Hamster{Roborovsk,"老公公",false}

	hs = append(hs,dm,cm,lgg)

	hMap = make(map[string]Hamster)
	hMap["短毛"]=dm
	hMap["长毛"]=cm
	hMap["老公公"]=lgg
}

func main() {
	//bytes, err := json.Marshal(dm)          //序列化结构体
	//bytes, err := json.Marshal(hs)			//序列化切片
	bytes, err := json.Marshal(hMap)          //序列化映射/Map
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(string(bytes))            //要把[]byte转化为string方便看
	}


	jstr := []byte("{\"短毛\":{\"Variety\":0,\"Name\":\"短毛\",\"Sex\":true}," +
		"\"老公公\":{\"Variety\":3,\"Name\":\"老公公\",\"Sex\":false}," +
		"\"长毛\":{\"Variety\":0,\"Name\":\"长毛\",\"Sex\":true}}")
	a :=map[string]Hamster{}   				   //先建一个容器
	err = json.Unmarshal(jstr, &a)            //json的反序列化   注意要先把字符串转化为[]byte
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(a)
	}

}
