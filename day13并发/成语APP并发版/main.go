package main

import (
	"time"
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main() {

	fq := make(chan string)
	fq2 := make(chan string)
	pq := make(chan string)
	save := make(chan bool)

	verse := getCmd()

	go func(){
		for _,c := range verse{
			idiom := fuzzyQuery(string(c))
			fq <- idiom
			fq2 <- idiom
		}
	}()

	go func(){
		for{
			words := <- fq2
			paraphrase := preciseQuery(words)
			pq <- paraphrase
		}

	}()

	go func(){
		timer := time.NewTimer(20 * time.Second)
		<- timer.C
		save <- true
	}()

	var idioms []string
	var paraphrases []string

	for{
		time.NewTicker(time.Second)
		select{
		case fqret := <- fq:
			idioms= append(idioms,fqret)
		case pqret := <- pq:
			paraphrases= append(paraphrases,pqret)
		case sq := <- save:
			fmt.Println("结束了吗？",sq)
			encod2json(idioms,paraphrases)
		    break
		}
	}
}


func getCmd() (verse string) {
	flag.StringVar(&verse,"verse","","输入诗句")
	flag.Parse()
	return
}

func fuzzyQuery(word string)(idiom string){
	resp, err := http.Get("http://route.showapi.com/1196-1?showapi_appid=86850&showapi_sign=0014d67cea944b63ad8f4c5dc6f70917&keyword="+word)
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
	jsonStr := string(bytes)

	var tempData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &tempData)
	val,_ := tempData["showapi_res_body"].(map[string]interface {})
	val1,_ :=val["data"].([]interface{})
	val2,_ := val1[0].(map[string]interface{})
	idiom = val2["title"].(string)
	return
}

func preciseQuery(words string)(paraphrase string){
	resp, err := http.Get("http://route.showapi.com/1196-2?showapi_appid=86850&showapi_sign=0014d67cea944b63ad8f4c5dc6f70917&keyword="+words)
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
	jsonStr := string(bytes)

	var tempData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &tempData)
	val,_ := tempData["showapi_res_body"].(map[string]interface {})
	val1,_ :=val["data"].(map[string]interface{})
	paraphrase = val1["content"].(string)
	return
}

func encod2json(idioms,paraphrases []string){
	/*maps := make(map[string]string)
	for _,v1 := range idioms{
		for _,v2 := range paraphrases{
			maps[v1] = v2
		}
	}*/
	file, _ := os.OpenFile("/Users/shenhaihui/go/src/learn/goBiliblili/13 并发/综合案例/成语APP并发版文件.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	defer file.Close()

	encoder := json.NewEncoder(file)
	//encoder.Encode(maps)
	encoder.Encode(idioms)
	encoder.Encode(paraphrases)
}