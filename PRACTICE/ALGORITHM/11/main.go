package main

import (
	"fmt"
	"strconv"
)

func main() {
	inp := 987
	fmt.Println(solution(inp))
	fmt.Println(solution2(inp))
}

func solution(n int) int {
	answer := 0

	for n > 0 {
		answer += n % 10
		n /= 10
	}

	return answer
}

func solution2(n int) int {
	answer := 0
	s := strconv.Itoa(n)
	for _, v := range s {
		answer += (int(v) - 48) // ascii! 49 == 1
	}
	return answer
}
