package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a rune
	a += 97

	fmt.Println(string(a))

	var b int
	b += 97
	fmt.Println(string(b))
	fmt.Println(reflect.TypeOf(b).Kind())
	fmt.Println(&b)

}
