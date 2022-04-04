package main

import (
	"fmt"
	"log"
	"strconv"
)

//https://programmers.co.kr/learn/courses/30/lessons/12947?language=go

func main() {
	var x int = 63
	fmt.Println(solution(x))
	fmt.Println(solution2(x))
}

func solution(x int) bool {
	var divide int
	for _, runne := range strconv.Itoa(x) { // converstion from Int to String
		add, err := strconv.Atoi(string(runne)) // converstion from rune to string, string to int
		if err != nil {
			log.Fatal(err)
		}
		divide += add
	}
	return x%divide == 0

}

func solution2(x int) bool {
	n := x
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return x%sum == 0
}
