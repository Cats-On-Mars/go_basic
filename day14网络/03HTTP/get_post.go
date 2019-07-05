package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	//get()
	post()
}

func get(){
	url := "http://wwww.baidu.com/s?wd=肉"

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(bytes))

}

func post(){
	url := "https://httpbin.org/post"

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("id=小慧慧"))
	checkError(err)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(bytes))

}

func checkError(err error){
	if err != nil{
		fmt.Println("错误",err.Error())
		os.Exit(1)
	}
}