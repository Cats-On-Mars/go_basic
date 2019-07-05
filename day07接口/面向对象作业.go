package main

import (
	"fmt"
	"flag"
)

type IPerson interface{
	Eat()
	Drink()
	Sleep()
}

type Person struct{}
func (p *Person)Eat(){}
func (p *Person)Drink(){}
func (p *Person)Sleep(){}



type Worker interface{
	Work() string
	Rest()
}

type Coder struct{Person}
func (w *Coder)Work()(output string){
	output = "代码"
	return
}
func (w *Coder)Rest(){
	fmt.Println("程序员休息")
}


type Teacher struct{Person}
func (w *Teacher)Work()(output string){
	output = "教案"
	return
}
func (w *Teacher)Rest(){
	fmt.Println("老师休息")
}


type Farmer struct{Person}
func (w *Farmer)Work()(output string){
	output = "农作物"
	return
}
func (w *Farmer)Rest(){
	fmt.Println("农民休息")
}


var Workers = make([]Worker,0)

func main() {

	var coder0 =Coder{Person{}}
	Workers = append(Workers,&coder0)
	Workers = append(Workers,new(Coder))
	Workers = append(Workers,new(Teacher))
	Workers = append(Workers,new(Farmer))

	today := getCmd()
	//var today int
	//fmt.Scan(&today)
	if today>0&&today<6{
		for _,x := range Workers{
			fmt.Println(x.Work())
		}
	}else if today == 6||today == 7{
		for _,v := range Workers{
			if val,ok := v.(*Coder);ok{
				fmt.Println(val.Work())
			}else{
				v.Rest()
			}
			//switch v.(type){
			//case *Coder:
			//	fmt.Println(v.Work())
			//default:
			//	v.Rest()
			//}
		}
	}else{
		fmt.Println("vaild cmd")
	}

}


func getCmd()(today int){
	flag.IntVar(&today,"今天星期几",0,"今天星期几")
	flag.Parse()
	return
}