package main

import (
	"fmt"
	"strings"
)

//https://programmers.co.kr/learn/courses/30/lessons/12930?language=go

func main() {
	fmt.Println(solution("  try  hello   a   world"))
	fmt.Println(solution("Bot eqew asds  qwe    wewe "))
	fmt.Println(solution(" jump  space one two  three  "))
}

func solution(s string) string {
	answer := ""
	cnt := 0
	for _, val := range s {
		if val == ' ' {
			if cnt == 0 {
				continue
			} else {
				answer += string(val)
				cnt = 0
			}
		} else {
			if cnt%2 == 0 {
				answer += strings.ToUpper(string(val))
			} else {
				answer += strings.ToLower(string(val))
			}
			cnt++
		}

	}
	return strings.Trim(answer, " ")
}
