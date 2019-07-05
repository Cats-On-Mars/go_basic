package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"time"
	"strings"
)

/*
·从开源接口showapi.com中获取成语数据；
·将其重构为Go的结构体数据；
·实现在命令行进行成语释义查询功能；
·要求实现数据持久化
*/

type Idiom struct {
	Title      string
	Spell      string
	Content    string
	Sample     string
	Derivation string
}

var idiomsMap map[string]Idiom

var DBFilePath = "d:/成语大全.json"

func init0() {
	/*	fileInfo, err := os.Stat(DBFilePath)
		if fileInfo != nil && err == nil {
			fmt.Println("数据已就绪")
			return
		}*/

	idiomsMap = make(map[string]Idiom)

	//获得数据
	jsonStr, _ := GetJson("http://route.showapi.com/1196-1?showapi_appid=19988&showapi_sign=968ad4fcc2144e41b5c366838d1b0ec4&keyword=肉&page=1&rows=10")
	//fmt.Println(jsonStr)

	//解析json为成语map
	ParseJson2Idioms(jsonStr)
	fmt.Println(idiomsMap)

	//数据的本地持久化
	//WriteIdioms2File("d:/成语大全.json")

	//查询已有成语的详细解释，解析
	time.Sleep(1 * time.Second)
	for title, _ := range idiomsMap {
		//查询每个成语的详细解释
		jsonStr, err := GetJson("http://route.showapi.com/1196-2?showapi_appid=19988&showapi_sign=968ad4fcc2144e41b5c366838d1b0ec4&keyword=" + title + "&page=1&rows=10")
		if err != nil {
			fmt.Printf("查询%s释义失败\n", title)
			continue
		}

		//将jsonStr解析为Idiom对象
		var idiom = Idiom{Title: title}
		var temp map[string]interface{}
		json.Unmarshal([]byte(jsonStr), &temp)
		temp = temp["showapi_res_body"].(map[string]interface{})["data"].(map[string]interface{})
		idiom.Spell = temp["spell"].(string)
		idiom.Sample = temp["samples"].(string)
		idiom.Derivation = temp["derivation"].(string)
		idiom.Content = temp["content"].(string)

		//替换之前的idiom
		idiomsMap[title] = idiom
		fmt.Printf("%s的释义更新成功\n", title)
		time.Sleep(1 * time.Second)
	}

	//数据的本地持久化
	WriteIdioms2File("d:/成语大全.json")

}

func init() {
	ReadIdiomsFromFile()
}

//从文件加载成语数据
func ReadIdiomsFromFile() {
	idiomsMap = make(map[string]Idiom)
	//读入json文件数据
	dstFile, _ := os.OpenFile(DBFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	defer dstFile.Close()
	decoder := json.NewDecoder(dstFile)
	err := decoder.Decode(&idiomsMap)
	if err != nil {
		fmt.Println("加载数据失败！err=", err)
	} else {
		fmt.Println("成功加载数据！")
		fmt.Println("idiomsMap=", idiomsMap)
	}
}

func main() {
	//获得命令行参数
	//idiom.exe -cmd idioms -keyword 肉
	//name value usage
	cmdInfo := [3]interface{}{"cmd", "未知命令", "ambiguous=模糊查询，accurate=精确查询"}
	keywordInfo := [3]interface{}{"keyword", "未知关键字", "查询关键字"}
	retValuesMap := GetCmdlineArgs(cmdInfo, keywordInfo)
	fmt.Println(retValuesMap)

	//执行查询
	cmd := retValuesMap["cmd"].(string)
	keyword := retValuesMap["keyword"].(string)
	if cmd == "ambiguous" {
		//模糊查询
		fmt.Println("开始模糊查询...")
		fmt.Println(idiomsMap)
		titles := make([]string, 0)
		for title, _ := range idiomsMap {

			if strings.Contains(title, keyword) {
				titles = append(titles, title)
			}
		}
		if len(titles) == 0 {
			fmt.Println("没有查询到结果！")
		} else {
			fmt.Printf("查询到%d条结果:\n", len(titles))
			for _, title := range titles {
				fmt.Println(title)
			}
		}

	} else if cmd == "accurate" {
		idiom := idiomsMap[keyword]
		PrintIdiom(idiom)
	} else {
		fmt.Println("非法命令：", cmd)
	}

}

func main0() {
	init0()
}

func PrintIdiom(idiom Idiom) {
	if idiom.Title != "" {
		fmt.Printf("Title:%s\n", idiom.Title)
		fmt.Printf("Spell:%s\n", idiom.Spell)
		fmt.Printf("Sample:%s\n", idiom.Sample)
		fmt.Printf("Derivation:%s\n", idiom.Derivation)
		fmt.Printf("Content:%s\n", idiom.Content)
	} else {
		fmt.Println("未找到成语！")
	}
}

//将成语map写出到指定文件
func WriteIdioms2File(path string) {
	dstFile, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer dstFile.Close()

	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(idiomsMap)
	if err != nil {
		fmt.Println("写出json文件失败，err=", err)
		return
	}
	fmt.Println("写出json文件成功！")
}

//解析json为成语map
func ParseJson2Idioms(jsonStr string) {
	//将json转换为go数据
	tempMap := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &tempMap)
	//fmt.Println(tempMap)
	dataSlice := tempMap["showapi_res_body"].(map[string]interface{})["data"].([]interface{})
	//fmt.Printf("type=%T,value=%v",dataSlice,dataSlice)
	for _, v := range dataSlice {
		title := v.(map[string]interface{})["title"].(string)
		idiom := Idiom{Title: title}
		idiomsMap[title] = idiom
	}
}

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
