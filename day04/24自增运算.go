package main

import "fmt"

func main() {
	var a int = 10

	a += 3
	a -= 3
	a *= 3
	a /= 3
	a %= 3

	a &= 3
	a |= 3
	a ^= 3
	a <<= 3
	a >>= 3

	fmt.Println(a)
}
