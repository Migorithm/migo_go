package main

import (
	"fmt"
	"math"
)

//https://programmers.co.kr/learn/courses/30/lessons/12934?language=go

func main() {

	fmt.Println(solution(121))
	fmt.Println(solution2(121))

}

func solution(n int64) int64 {
	square := math.Pow(float64(n), 0.5)
	remainder := math.Mod(square, 1)
	if remainder > 0 {
		return -1
	} else {
		var answer int64 = int64(math.Pow(float64(square+1), float64(2)))
		return answer
	}
}

//A little cleaner
func solution2(n int64) int64 {
	sqrt := int64(math.Pow(float64(n), 0.5))
	if sqrt*sqrt == n {
		sqrt++
		return sqrt * sqrt
	}
	return -1
}
