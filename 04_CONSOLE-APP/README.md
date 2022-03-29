# Modulation

    go mod init myapp

# Get started!
```go
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') //1
		userInput = strings.Replace(userInput, "\r\n", "", -1) //2
        userInput = strings.Replace(userInput,"\n","",-1) //3

        if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}
	}
}
```
1. What is this underscore(_) for? If you roll your mouse over .ReadString, it says:
- it takes one parameter which is delimiter of type byte
- it returns two things - string, error. 
So essentially by putting underscore you're saying that "I don't care about the second thing" 

2. It will replace \\r\\n with "" on any occurrence(when the last value < 0 ) (For windows)

3. It will replace \\n with "" on any occurrence(when the last value < 0 ) (For mac and linux)


## What if you want to look for a single key press(like esc button)