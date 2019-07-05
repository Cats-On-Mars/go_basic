package main

import (
	"sync"
	"fmt"
	"time"
)

//必须保证并发安全的数据
type Account struct {
	name  string
	money float32

	//定义该数据的互斥锁
	mt    sync.Mutex
}

//本方法不能被并发执行——并发安全的
func (a *Account) saveGet(amount float32) {
	//先将资源锁起来
	a.mt.Lock()

	//执行操作
	fmt.Println("操作前：", a.money)
	a.money += amount
	fmt.Println("操作后：", a.money)
	<-time.After(3 * time.Second)

	//释放资源
	a.mt.Unlock()
}

//本方法可以被并发执行——不是并发安全的,无此必要
func (a *Account) getName() string {
	return a.name
}

func main() {
	a := Account{name: "张全蛋", money: 1000}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		//调用一个加锁的方法（同步）
		a.saveGet(500)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//调用一个加锁的方法（同步）
		a.saveGet(-500)
		wg.Done()
	}()

	for i:=0;i<3 ;i++  {
		wg.Add(1)
		go func() {
			//调用一个普通的没有访问锁的方法（异步）
			fmt.Println(a.getName())
			wg.Done()
		}()
	}

	wg.Wait()
}
