### Variable Declaration
```go
package main

import "fmt"

func main() {
	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2
	fmt.Println(firstNumber)

	// another way, declare type and name and assign value
	var secondNumber = 5

	// one step variable: declare name, assign value, and let Go figure out the type
	subtraction := 7

}

```


### Guess number game
```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready." //this can never be changed

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //to get number 2-10 because you don't want 0 and 1
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2

	reader := bufio.NewReader(os.Stdin)

	//display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n') //use \n for enter key to be pressed

	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	answer = firstNumber*secondNumber - subtraction // answer is always three.
	fmt.Println("The answer is", answer)

}
```


### Making function more task specific and consise
Let's refactor it so main function contains less components. 
```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready." //this can never be changed

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //to get number 2-10 because you don't want 0 and 1
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	answer := firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)

}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n') //use \n for enter key to be pressed

	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("The answer is", answer)

	// Different 
	var a int
	fmt.Scanln(&a)

	fmt.Println(a)

}

```


