package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"flag"
	"encoding/json"
	"os"
)

/*
·从开源接口showapi.com中获取成语数据；
·将其重构为Go的结构体数据；
·实现在命令行进行成语释义查询功能；
·要求实现数据持久化
*/

func main() {

	APINum,keyWord := getKeyword()

	url1 := "http://route.showapi.com/1196-1?showapi_appid=86850&showapi_sign=0014d67cea944b63ad8f4c5dc6f70917&keyword="+keyWord
	url2 := "http://route.showapi.com/1196-2?showapi_appid=86850&showapi_sign=0014d67cea944b63ad8f4c5dc6f70917&keyword="+keyWord

	var jsonStr string

	switch APINum{
	case 1:
		jsonStr, _ = GetJsonStr(url1)
	case 2:
		jsonStr, _ = GetJsonStr(url2)
	}
	//fmt.Println(jsonStr)     //为了调试方便

	goData := ParseJson2Godata(jsonStr)
	//fmt.Println(goData)

	Encode2File(goData)

}

func getKeyword() (APINum int,keyword string) {
	flag.StringVar(&keyword,"keyword","","要搜索的成语")
	flag.IntVar(&APINum,"APINum",0,"要调用的接口")
	flag.Parse()
	return
}


func GetJsonStr(url string)(jsonStr string,err error){
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("HTTP请求失败",err)
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("读取响应失败",err)
		return
	}

	jsonStr = string(bytes)
	return
}


func ParseJson2Godata(jsonStr string) interface{}{

	var tempData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &tempData)
	val,_ := tempData["showapi_res_body"].(map[string]interface {})
	return val["data"]

}


func Encode2File(data interface{}){

	file, err := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/day10/成语释义.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil{
		fmt.Println("写出JSON文件失败")
		return
	}else{
		fmt.Println("写出JSON文件成功")
	}
}
