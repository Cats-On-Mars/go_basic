package main

import "fmt"

//战士接口
type Soldier interface {
	Attack() (bloodLoss int)
	Defend()
}

//骑兵
type Rider struct{}

//法师
type Master struct{}

func (r *Rider) Attack() (bloodLoss int) {
	fmt.Println("发动战争践踏，撞死你，踩死你，扎死你")
	return 50
}
func (r *Rider) Defend() {
	fmt.Println("不能跑，我得快点走")
}

func (r *Master) Attack() (bloodLoss int) {
	fmt.Println("天崩地裂，电闪雷鸣，群星坠落，万籁俱寂，遍布下去了...")
	return 10000
}
func (r *Master) Defend() {
	fmt.Println("回城")
}

func main() {
	soldiers := make([]Soldier, 0)
	soldiers = append(soldiers, new(Rider))
	soldiers = append(soldiers, new(Master))

	//强调共性时
	/*	fmt.Println("全体进攻")
		for _, f := range soldiers {
			//所有子类实现都调用共同的父类方法
			f.Attack()
		}*/

	//强调个性时
	fmt.Println("敌寇势大，法师进攻，骑士防守")
	for _, s := range soldiers{

		/*		//类型断言方式1
				switch s.(type) {
				case *Rider:
					s.Defend()
				case *Master:
					s.Attack()
				default:
					fmt.Println("来了个什么鬼...")
				}*/

		//类型断言方式2
		//判断是当前s实例是不是骑士
		// 如果是:ok为true，riderPtr有值
		//如果不是：ok为false，riderPtr为nil
		if riderPtr,ok:=s.(*Rider);ok{
			riderPtr.Defend()
		}
		//判断是当前s实例是不是法师
		// 如果是:ok为true，masterPtr有值
		//如果不是：ok为false，masterPtr为nil
		if masterPtr,ok:=s.(*Master);ok{
			masterPtr.Attack()
		}

	}

}
