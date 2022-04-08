package main

import (
	"fmt"
	"sort"
)

//https://programmers.co.kr/learn/courses/30/lessons/12915?language=go

func main() {
	inp := []string{"sun", "bed", "car"}
	inp2 := []string{"abce", "abcd", "cdx"}
	solution(inp, 1)
	solution(inp2, 2)
	fmt.Println(solution2(inp2, 2))
}

func solution(strings []string, n int) []string {
	var dic = make(map[string][]string)
	str := ""
	// make a bucket for letter at index 'n'
	for _, val := range strings {
		dic[string(val[n])] = append(dic[string(val[n])], val)
	}

	// sort the values inside the bucket and store the key
	for key, _ := range dic {
		sort.Strings(dic[key])
		str += key
	}
	into_bytes := []byte(str)

	//sort the key
	sort.Slice(into_bytes, func(a, b int) bool { return into_bytes[a] < into_bytes[b] })

	answer := []string{}
	for _, val := range into_bytes {
		answer = append(answer, dic[string(val)]...)
	}
	fmt.Println(answer)
	return answer
}

func solution2(strings []string, n int) []string {

	sort.Slice(strings, func(i, j int) bool {
		strA := []byte(strings[i]) // take the first word
		strB := []byte(strings[j]) // take the second word
		if strA[n] == strB[n] {
			return strings[i] < strings[j] //you can compare string to string ! (abc < abb returns false)
		}
		return strA[n] < strB[n]
	})

	return strings
}
