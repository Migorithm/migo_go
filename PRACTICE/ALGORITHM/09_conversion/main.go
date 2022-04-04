package main

//https://programmers.co.kr/learn/courses/30/lessons/12933?language=go

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution(118372))
	fmt.Println(solution2(118372))
}

func solution(n int64) int64 {
	str_num := strconv.Itoa(int(n)) //There won't be error when you convert int into string
	lst_string := []string{}
	fmt.Println(lst_string)
	for _, runne := range str_num {
		lst_string = append(lst_string, string(runne))
	}
	//sort.Reverse(sort.StringSlice(lst_string))
	sort.Sort(sort.Reverse(sort.StringSlice(lst_string)))
	answer, err := strconv.Atoi(strings.Join(lst_string, "")) // there might be an error when you convert string into int
	if err != nil {
		log.Fatal(err)
	}
	return int64(answer)

}

func solution2(n int64) int64 {

	arr := make([]int, 0, 16)

	for n != 0 {
		arr = append(arr, int(n%10))
		n /= 10
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	var ret int64

	for _, v := range arr {
		ret *= 10
		ret += int64(v)
	}

	return ret
}
