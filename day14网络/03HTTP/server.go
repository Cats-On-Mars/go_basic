package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("我有一头小毛驴我从来也不骑"))
	})

	http.HandleFunc("/hello2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RemoteAddr))
		writer.Write([]byte(request.Method))
	})

	http.HandleFunc("/hello3", func(writer http.ResponseWriter, request *http.Request) {
		data, _ := ioutil.ReadFile("/Users/shenhaihui/go/src/learn/goBiliblili/day14网络/HTTP/肉.html")
		writer.Write(data)

	})

	http.ListenAndServe("127.0.0.1:8080",nil)
}
