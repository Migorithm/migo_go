package main

import "fmt"

func main() {
	var x int = 5748
	var p *int
	p = &x
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Addresses x =", &x, "Address p = ", &p)
	x = 31

	fmt.Println("Value stored in p = ", *p)
	*p = 503

	fmt.Println("Value stored in x = ", x)
}
