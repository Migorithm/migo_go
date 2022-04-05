package main

import (
	"fmt"
	"strconv"
)

//https://programmers.co.kr/learn/courses/30/lessons/12932

func main() {
	var str int64 = 12345
	fmt.Println(solution(str))
	fmt.Println(solution2(str))
}

func solution(n int64) []int {
	str := strconv.Itoa(int(n))
	var ints = make([]int, len(str))
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		ivar, _ := strconv.Atoi(string(str[j])) // each element in string is treated as rune.
		jvar, _ := strconv.Atoi(string(str[i]))
		ints[i] = ivar
		ints[j] = jvar

	}

	return ints
}

//better solution!
func solution2(n int64) []int {
	var answer []int
	temp := n
	for temp > 0 {
		answer = append(answer, int(temp%10))
		temp /= 10
	}
	return answer
}
