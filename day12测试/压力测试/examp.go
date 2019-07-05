package main

/*func main() {
	fmt.Println(fib1(5))
	fmt.Println(fib2(5))
}*/

func fib1(i int) int{
	if i==0||i==1{
		return i
	}else{
		return fib1(i-1)+fib1(i-2)
	}
}


func fib2(n int) int{
	x,y := 1,1
	for i := 1; i < n ; i++ {
		x,y = y,x+y
	}
	return x
}