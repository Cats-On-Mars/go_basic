package main

import "fmt"

//0=按gid升序，1=按gid降序，2=按price升序，3=按price降序
func sortGoods(glist []*Goods, order int) {
	for i := 0; i < len(glist)-1; i++ {
		for j := i + 1; j < len(glist); j++ {
			switch order {
			case 0:
				if glist[j].gid < glist[i].gid {
					glist[i], glist[j] = glist[j], glist[i]
				}
			case 1:
				if glist[j].gid > glist[i].gid {
					glist[i], glist[j] = glist[j], glist[i]
				}
			case 2:
				if glist[j].price < glist[i].price {
					glist[i], glist[j] = glist[j], glist[i]
				}
			case 3:
				if glist[j].price > glist[i].price {
					glist[i], glist[j] = glist[j], glist[i]
				}
			default:
				fmt.Println("order必须是：0,1,2,3之一!")
			}

		}
	}
}
