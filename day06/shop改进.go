package main

import (
	"fmt"
	"game"
)

const(
	ECONOMICS = 1+iota
	COMPUTER
	BLOCKCHAIN
	LITERATURE
)

type Goods struct{
	gid int
	name string
	cid int
	price float32
}

//var allGoods = make([]Goods,7)
var allGoods []*Goods    //传指针

//工厂模式，gid自动生成，如果要修改产品生成逻辑，只需要修改一处
func createGoods(name string,cid int,price float32) *Goods{

	gPtr := new(Goods)

	gPtr.gid = cid*10000+len(allGoods)+1
	gPtr.name = name
	gPtr.cid = cid
	gPtr.price = price

	return gPtr
}

//初始化函数，在main之前执行
func init(){
	allGoods = make([]*Goods,0)
	allGoods = append(allGoods,createGoods("信息论",COMPUTER,67.00))
}


func main() {

	allGoods[0]= &Goods{0,"货币金融学",ECONOMICS,85.00}
	allGoods[1] = &Goods{1,"经融学",ECONOMICS,59.50}
	allGoods[2] = &Goods{2,"Go并发编程实战",COMPUTER,79.00}
	allGoods[3] = &Goods{3,"计算机网络",COMPUTER,64.50}
	allGoods[4] = &Goods{4,"MySQL必知必会",COMPUTER,94.00}
	allGoods[5] = &Goods{5,"区块链BlockChain",BLOCKCHAIN,150.00}
	allGoods[6] = &Goods{6,"失控",LITERATURE,84.50}

	cmd, gid, cid, order := game.GetCmd()

	switch cmd{
	case "single":
		if gid>-1 && gid<len(allGoods){
			fmt.Println(allGoods[gid])
		}else{
			fmt.Println("vaild gid")
		}

	case "category":
		fmt.Println(newOrderFunc(cid,order))

	default:
		fmt.Println("vaild command")
	}

}



func newOrderFunc(cid,order int) []Goods{

	newAllGoods := make([]Goods,0)
	for i := 0;i<len(allGoods);i++  {
		if allGoods[i].cid == cid{
			newAllGoods = append(newAllGoods,allGoods[i])
		}
	}

	switch order{
	case 0:
		for i:=0;i<len(newAllGoods)-1;i++{
			for j:= i+1;j<len(newAllGoods);j++{
				if newAllGoods[i].gid>newAllGoods[j].gid{
					newAllGoods[i],newAllGoods[j]=newAllGoods[j],newAllGoods[i]
				}
			}
		}

	case 1:
		for i:=0;i<len(newAllGoods)-1;i++{
			for j:= i+1;j<len(newAllGoods);j++{
				if newAllGoods[i].gid<newAllGoods[j].gid{
					newAllGoods[i],newAllGoods[j]=newAllGoods[j],newAllGoods[i]
				}
			}
		}
	case 2:
		for i:=len(newAllGoods)-1;i>0;i--{
			for j:=0;j<i;j++{
				if newAllGoods[j].price>newAllGoods[j+1].price{
					newAllGoods[j],newAllGoods[j+1]=newAllGoods[j+1],newAllGoods[j]
				}
			}
		}
	case 3:
		for i:=len(newAllGoods)-1;i>0;i--{
			for j:=0;j<i;j++{
				if newAllGoods[j].price<newAllGoods[j+1].price{
					newAllGoods[j],newAllGoods[j+1]=newAllGoods[j+1],newAllGoods[j]
				}
			}
		}
	}
	return newAllGoods
}




