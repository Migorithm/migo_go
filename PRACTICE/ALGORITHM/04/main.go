package main

import "fmt"

//https://programmers.co.kr/learn/courses/30/lessons/12950?language=go

func main() {
	inp := make([][]int, 2)
	inp2 := make([][]int, 2)

	//var inp2 [][]int
	inp[0] = append(inp[0], 1, 2)
	inp[1] = append(inp[1], 2, 3)
	inp2[0] = append(inp2[0], 3, 4)
	inp2[1] = append(inp2[1], 5, 6)
	// fmt.Println(len(inp))
	// fmt.Println(len(inp[1]))

	fmt.Println(solution(inp, inp2))

	fmt.Println(solution2(inp, inp2))
}

// Solution1, creating empty slice of slice and put it.
func solution(arr [][]int, arr2 [][]int) [][]int {
	answer := make([][]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[1]); j++ {
			answer[i] = append(answer[i], arr[i][j]+arr2[i][j])
		}
	}
	return answer
}

// Solution2, adding the value directly
func solution2(arr1 [][]int, arr2 [][]int) [][]int {
	for i, _ := range arr1 { //here, i is index. _ is value but not used.
		for j, _ := range arr2 {
			arr1[i][j] += arr2[i][j]
		}
	}
	return arr1
}
