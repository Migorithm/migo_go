package main

import "fmt"

//https://programmers.co.kr/learn/courses/30/lessons/12926

func main() {
	fmt.Println(solution("a B Z", 4))
}

func solution(s string, n int) string {
	fmt.Println('a') // 97
	fmt.Println('z') // 122
	fmt.Println('A') // 65
	fmt.Println('Z') // 90

	answer := ""
	for _, val := range s {
		//66 ~ 89
		str := int(val) + n
		if val > 96 {
			if str > 122 {
				answer += string(rune(96 + str - 122))
			} else {
				answer += string(rune(str))
			}
		} else if val > 64 {
			if str > 90 {
				answer += string(rune(64 + str - 90))
			} else {
				answer += string(rune(str))
			}

		} else {
			answer += " "

		}
	}

	return answer
}
