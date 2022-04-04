package main

//https://programmers.co.kr/learn/courses/30/lessons/12935?language=go

import (
	"fmt"
)

func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(solution(arr))
}

func solution(arr []int) []int {
	// There is no built-in min()
	min := arr[0]
	index := 0
	for idx, val := range arr {
		if val <= min {
			min = val
			index = idx
		}
	}
	answer := append(arr[:index], arr[index+1:]...) //concatenate two slices using three dot notation

	if len(answer) == 0 {
		return []int{-1}
	}
	return answer
}
