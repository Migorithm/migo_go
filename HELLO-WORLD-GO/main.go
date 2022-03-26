package main

import (
	"bufio" //1
	"fmt"
	"myapp/doctor"
	"os" //2
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(doctor.Intro())
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) //1
		userInput = strings.Replace(userInput, "\n", "", -1)   //2

		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}
