package main

import "fmt"

//가장 작은 수 제거

func main() {

	arr := []int{4, 3, 0, 1}
	fmt.Println(solution(arr))
}

func solution(arr []int) []int {
	// Assuming there is no built-in min() function in Go...
	min := arr[0]
	index := 0
	for idx, val := range arr {
		if val <= min {
			min = val
			index = idx
		}
	}
	answer := append(arr[:index], arr[index+1:]...)
	if len(answer) == 0 {
		return []int{-1}
	}
	return answer
}
