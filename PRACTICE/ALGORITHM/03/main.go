package main

import (
	"fmt"
	"strings"
)

//https://programmers.co.kr/learn/courses/30/lessons/12948?language=go

func main() {
	var parameter string
	parameter = "01033334444"
	fmt.Println(solution(parameter))

}

func solution(phone_number string) string {
	//concatenation of *

	return strings.Repeat("*", len(phone_number)-4) + phone_number[len(phone_number)-4:]
}

// lesson learned!
// String can be repeated by strings.Repeat function

// https://www.geeksforgeeks.org/string-package-in-golang/
