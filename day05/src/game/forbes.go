package game

import (
	"fmt"
	"strconv"
)

func Forbes(){
	fmt.Println("请按提示输入，结束请输入over")

	var name,fortune string

	type person struct{
		name string
		fortune int
	}

	var p person
	var ps []person

	for {
		fmt.Print("请输入姓名:")
		fmt.Scanln(&name)
		if name == "over"{
			break
		}

		fmt.Print("请输入财富:")
		fmt.Scanln(&fortune)
		if fortune == "over"{
			break
		}

		ret,_ := strconv.Atoi(fortune)
		p = person{name,ret}
		ps = append(ps,p)
	}

	for i:=0;i<len(ps)-1;i++{
		for j:=i+1;j<len(ps);j++{
			if ps[i].fortune<ps[j].fortune{
				ps[i],ps[j]=ps[j],ps[i]
			}
		}
	}

	//fmt.Println(ps)
	for _,x := range ps{
		fmt.Println(x)
	}

}