package main

const (
	Toy       = 1 + iota //1
	Weapon               //2
	BigmanToy            //3
)

type Goods struct {
	gid   int
	name  string
	cid   int
	price float32
}

/*
流水线套路化生产产品——工厂模式
*/
func CreateGoods(name string, cid int, price float32) *Goods {
	gPtr := new(Goods)
	gPtr.name = name
	gPtr.cid = cid
	gPtr.price = price

	//丢入分类map
	cateGoodsMap[cid] = append(cateGoodsMap[cid], gPtr)

	//按算法生成gid
	//gPtr.gid = cid*100000 + (len(allGoods)+1)
	gPtr.gid = cid*100000 + (len(cateGoodsMap[cid]))

	//丢入单品map
	allGoodsMap[gPtr.gid] = gPtr

	return gPtr
}
