package main

import (
	"fmt"
)

func main() {
	// var str = make([]string, 5)
	// str = append(str, "1")
	// str = append(str, "2")
	// str = append(str, "3")
	var str [5]string

	fmt.Println(str)

	increase(str)
	fmt.Println(str)

}

func increase(str [5]string) {
	str[1] = "1"
	fmt.Println(str)

}
