package main

import "fmt"

func main() {
	fmt.Printf("%T\n",1.23e9)

	//科学记数法的浮点表示
	fmt.Printf("%f\n",1.234e9)
	fmt.Printf("%f\n",1.234E9)
	fmt.Printf("%.15f\n",1.234e-9)
	fmt.Printf("%.15f\n",1.234E-9)

	//浮点数的科学计数法表示
	fmt.Printf("%e\n",123458668.0)
	fmt.Printf("%E\n",0.00000018646354)
	fmt.Printf("%.2e\n",1343728656.0)
	fmt.Printf("%.2E\n",0.000000001245)

}
