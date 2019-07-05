package main

import "fmt"

//定义结构体
type book struct {
	//书名属性
	name string
	//价格属性
	price float64
}

func main() {
	//demo91()
	//demo92()
	//demo93()

	demo94()
}

//通过结构体对象，或结构体指针访问结构体属性
func demo94() {
	b1 := book{"三国", 45.67}
	showBookInfo(b1)
	showBookInfo2(&b1)

	bp := new(book)
	bp.name = "三国"
	bp.price = 45.67
	showBookInfo(*bp)
	showBookInfo2(bp)
}


//通过对象查看结构体属性
func showBookInfo(b book)  {
	fmt.Println(b.name)
	fmt.Println(b.price)
}

//通过指针查看结构体属性
func showBookInfo2(b *book)  {
	fmt.Println(b.name)
	fmt.Println(b.price)
}

//创建对象时直接对属性赋值
func demo91()  {
	b := book{"水许传",34.56}
	fmt.Printf("type=%T，value=%#v\n",b,b)
}

//创建对象指针并给属性赋值
func demo93() {
	//创建book指针
	bookPtr := new(book)
	fmt.Printf("type=%T，value=%#v,address=%p\n", bookPtr, *bookPtr, bookPtr)

	//结构体指针给属性赋值的方式，与结构体对象一模一样
	bookPtr.name = "水许传"
	bookPtr.price = 34.56

	fmt.Println(*bookPtr)
}

//创建空白对象，并逐一对属性赋值
func demo92() {
	b := book{}
	b.name = "水许传"
	b.price = 34.56
	fmt.Printf("type=%T，value=%#v\n", b, b)
}