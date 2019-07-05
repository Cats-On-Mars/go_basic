package goBiliblili

import (
	"net/http"
	"flag"
)

type myHandler struct{}

func (mh *myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	phrase(w,r)
}


func main() {
	handler := myHandler{}
	server := http.Server{
		Addr:"http://route.showapi.com/1196-2",
		Handler: &handler,
	}
	server.ListenAndServe()
}

func getCmd()(showapi_appid,showapi_sign,keyword string){
	flag.StringVar(&showapi_appid,"showapi_appid","","易源应用id")
	flag.StringVar(&showapi_sign,"showapi_sign","","传递调用者的数字签名")
	flag.StringVar(&keyword,"keyword","","要搜索的成语")

	flag.Parse()
	return
}


func phrase(w http.ResponseWriter,r *http.Request) {
	showapi_appid,showapi_sign,keyword := getCmd()
	r.URL.RawQuery ="showapi_appid="+showapi_appid+"showapi_sign="+showapi_sign+"keyword="+keyword

}