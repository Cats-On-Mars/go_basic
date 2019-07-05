package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	//必须保证并发安全的数据
	type Account struct {
		money float32
	}

	var wg sync.WaitGroup
	account := Account{1000}
	fmt.Println(account)

	//资源互斥锁(谁抢到锁，谁先访问资源，其他人阻塞等待)
	//全局就这么一把锁，谁先抢到谁操作，其他人被阻塞直到锁释放
	var mt sync.Mutex

	//银行卡取钱
	wg.Add(1)
	go func() {
		//拿到互斥锁
		mt.Lock()

		//加锁的访问
		fmt.Println("取钱前：",account.money)
		account.money -= 500
		time.Sleep(time.Nanosecond)
		fmt.Println("取钱后：",account.money)
		wg.Done()

		//释放互斥锁
		mt.Unlock()
	}()

	//存折存钱
	wg.Add(1)
	go func() {
		//拿到互斥锁（如果别人先抢到，则阻塞等待）
		mt.Lock()

		fmt.Println("存钱前：",account.money)
		account.money += 500
		time.Sleep(time.Nanosecond)
		fmt.Println("存钱后：",account.money)
		wg.Done()

		//释放互斥锁
		mt.Unlock()
	}()

	wg.Wait()
}

