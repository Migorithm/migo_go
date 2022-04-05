package main

//https://programmers.co.kr/learn/courses/30/lessons/12928?language=go
import "fmt"

func main() {
	var n int = 5
	fmt.Println(solution(n))
}

func solution(n int) int {
	answer := 0
	for i := n; i > 0; i-- {
		if n%i == 0 {
			answer += i
		}
	}
	return answer

}
