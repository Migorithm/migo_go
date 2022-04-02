package main

//https://programmers.co.kr/learn/courses/30/lessons/12954?language=go
import "fmt"

func main() {
	var x int
	var n int
	fmt.Scan(&x, &n)

	fmt.Println(solution(x, n))
}

func solution(x int, n int) []int64 {
	var answer []int64
	for i := 1; i <= n; i++ {
		to_be_added := i * x
		answer = append(answer, int64(to_be_added)) // this is how you append stuff
	}
	return answer
}

// The type specification for a slice is []T, where T is the type of the elements of the slice.
// Unlike an array type, a slice type has no specified length.

// A slice can be created with the built-in function called make, which has the signature,
// func make([]T,len,cap) []T
// The make function takes a type, a length, and an optional capacity.
