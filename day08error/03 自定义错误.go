package main

import (
	"time"
	"fmt"
)

func main() {
	ret, err := getAreaIII(-5)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(ret)
	}




	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getAreaIV(-5)
}

/*实现的error接口的自定义异常*/
type InvalidArgError struct {          //定义更复杂的错误，不止有字符串信息
	info string
	when time.Time
	user string
}

func (iae *InvalidArgError)Error() string{
	return fmt.Sprintln(iae.info,iae.when,iae.user)
}

func NewInvalidArgError(info string) error {
	iaePtr := new(InvalidArgError)
	iaePtr.info = info
	iaePtr.when = time.Now()
	iaePtr.user = "bill"
	return iaePtr
}

func getAreaIII(r float32)(area float32,err error){
	if r < 0{
		err = NewInvalidArgError("半径不能为负")
		return
	}
	area = 3.14 * r * r
	return
}




/*不实现的error接口的自定义异常*/
type MyError struct {
	pain interface{}
	info string
}
func (me *MyError)String() string {
	return fmt.Sprintf("%s:%v\n",me.info,me.pain)
}

func getAreaIV(r float32)(area float32) {
	if r < 0{
		panic(&MyError{r,"傻吊半径不能为负数"})   //直接抛一个结构体（的字符串华输出方式）
	}
	area = 3.14 * r * r
	return
}