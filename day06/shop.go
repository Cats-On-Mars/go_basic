package main

import (
	"fmt"
	"game"
)

type book struct{
	gid int
	name string
	cid int
	price float32
}

var books = make([]book,7)


func main() {

	books[0]= book{0,"货币金融学",1,85.00}
	books[1] = book{1,"经融学",1,59.50}
	books[2] = book{2,"Go并发编程实战",2,79.00}
	books[3] = book{3,"计算机网络",2,64.50}
	books[4] = book{4,"MySQL必知必会",2,94.00}
	books[5] = book{5,"区块链BlockChain",3,150.00}
	books[6] = book{6,"失控",4,84.50}

	cmd, gid, cid, order := game.GetCmd()

	switch cmd{
	case "single":
		if gid>-1 && gid<len(books){
			fmt.Println(books[gid])
		}else{
			fmt.Println("vaild gid")
		}

	case "category":
		fmt.Println(orderFunc(cid,order))

	default:
		fmt.Println("vaild command")
	}

}



func orderFunc(cid,order int) []book{

	newBooks := make([]book,0)
	for i := 0;i<len(books);i++  {
		if books[i].cid == cid{
			newBooks = append(newBooks,books[i])
		}
	}

	switch order{
	case 0:
		for i:=0;i<len(newBooks)-1;i++{
			for j:= i+1;j<len(newBooks);j++{
				if newBooks[i].gid>newBooks[j].gid{
					newBooks[i],newBooks[j]=newBooks[j],newBooks[i]
				}
			}
		}

	case 1:
		for i:=0;i<len(newBooks)-1;i++{
			for j:= i+1;j<len(newBooks);j++{
				if newBooks[i].gid<newBooks[j].gid{
					newBooks[i],newBooks[j]=newBooks[j],newBooks[i]
				}
			}
		}
	case 2:
		for i:=len(newBooks)-1;i>0;i--{
			for j:=0;j<i;j++{
				if newBooks[j].price>newBooks[j+1].price{
					newBooks[j],newBooks[j+1]=newBooks[j+1],newBooks[j]
				}
			}
		}
	case 3:
		for i:=len(newBooks)-1;i>0;i--{
			for j:=0;j<i;j++{
				if newBooks[j].price<newBooks[j+1].price{
					newBooks[j],newBooks[j+1]=newBooks[j+1],newBooks[j]
				}
			}
		}
	}
	return newBooks
}



