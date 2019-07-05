package main

import "fmt"

/*
switch x.(tpye){
case
}
*/


/*
val = x.(T)
vai,ok = x.(T)
if val,ok := x.(T);ok{....}

*/

type AI interface{
	funcAI1()
	funcAI2()
}


type AO1 struct{}

func (ao *AO1)funcAI1(){
	fmt.Println("funcAO11")
}
func (ao *AO1)funcAI2(){
	fmt.Println("funcAO12")
}


type AO2 struct{}

func (ao *AO2)funcAI1(){
	fmt.Println("funcAO21")
}
func (ao *AO2)funcAI2(){
	fmt.Println("funcAO22")
}



func main() {

	ais :=make([]AI,0)
	ais = append(ais,new(AO1))
	ais = append(ais,new(AO2))

	//共性
	//for _,x := range ais{
	//	x.funcAI1()
	//}

	//个性
	//类型断言方式1      适用于想判断所有类型的
	for _,x := range ais{
		switch x.(type){
		case *AO1:
			x.funcAI1()
		case *AO2:
			x.funcAI2()
		}
	}

	//类型断言方式2      适用于想判断个别类型的
	for _,x := range ais{
		if val,ok := x.(*AO1);ok{
			val.funcAI1()
		}
		if val,ok := x.(*AO2);ok{
			val.funcAI2()
		}
	}

}
