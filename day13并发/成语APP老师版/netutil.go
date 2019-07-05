package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

//从url拿到json数据
func GetJson(url string) (jsonStr string, err error) {

	//获得网络数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http请求失败,err=", err)
		return
	}
	//延时关闭网络IO资源
	defer resp.Body.Close()

	//resp.Body实现了Reader接口，对其进行数据读入
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络数据失败,err=", err)
		return
	}

	//将网络数据转化为字符串输出
	jsonStr = string(bytes)
	//fmt.Println(jsonStr)

	return
}


//模糊查询
func DoAmbiguousQuery(keyword string,page string)(jsonStr string){
	/*TODO*/
	return ""
}

//精确查询
func DoAccurateQuery(keyword string,page string)(jsonStr string){
	/*TODO*/
	return ""
}





