package main

import "fmt"

//이 문제에는 표준 입력으로 두 개의 정수 n과 m이 주어집니다.
//별(*) 문자를 이용해 가로의 길이가 n, 세로의 길이가 m인 직사각형 형태를 출력해보세요.
//5 3
//*****
//*****
//*****

func main() {
	var a, b int
	fmt.Scan(&a, &b) // This is how you get two inputs with a space in between
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
