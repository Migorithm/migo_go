package main

//https://programmers.co.kr/learn/courses/30/lessons/12943?language=go

import "fmt"

func main() {
	var arr int = 16
	fmt.Println(solution(arr))
}

func solution(num int) int {
	cnt := 0
	for num != 1 && cnt < 500 {
		if num%2 == 0 {
			num /= 2
			cnt += 1
		} else {
			num = num*3 + 1
			cnt += 1
		}
	}
	if cnt < 500 {
		return cnt
	} else {
		return -1
	}

}
