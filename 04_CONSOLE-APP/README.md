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

2. It will replace \\r\\n with "" on any occurrence(when the last value < 0 ) (For windows, this should come first)

3. It will replace \\n with "" on any occurrence(when the last value < 0 ) (For mac and linux)


## What if you want to look for a single key press(like esc button)
You can do that in Go, but it's not as easy as in C# or C.<br>
So instead of writing hundred lines of code, let's use a third-party package from the following link<br>

	https://github.com/eiannone/keyboard


### go get it 
The way you install go package is different from the way you do in Python or Java.<br>

**Installation:**

	go get -u github.com/eiannone/keyboard

The -u flag instructs 'get' to update modules providing dependencies of packages named on the command line<br>
to use newer minor or patch releases when available.<br><br>

By now, you will see go.sum is there in the project directory, the contents of which you don't want to fiddle with. And more notably,<br>
go.mod:
```plain
module myapp

go 1.18

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807 // indirect
	golang.org/x/sys v0.0.0-20220329152356-43be30ef3008 // indirect
)
```

### use of third-party module

```go

package main

import (
	"fmt"
	"log"
	"github.com/eiannone/keyboard"
)
func main() {
	err := keyboard.Open() //1
	if err != nil {  //2
		log.Fatal(err) //3
	}
	defer func(){  //4 
		_ = keyboard.Close() //5
	}()  

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey() //6
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key) //7
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc { //8
			break
		}

	}
	fmt.Println("Program exiting...")
}
```
2. nil is equal to None in Python

3. built-in log function

4. whatever follows the **defer** keyword won't execute immediately. Instead, it will execute as soon as the current function(main function) ends. Plus, this is the way you define anonymous function. Instead of using a "finally" block, we use the "defer" statement in Go. When we use "defer", the statement will run just before the function exits. The good thing about this is that we can put a "defer" almost anywhere in the code.

5. as soon as the main function finishes, it will close the keyboard. It might throw an error, but we ignore it. 

6. If you get your mouse hovering over GetSingleKey, you will see the return types are Rune, keyboard.key and err. Rune literals are just 32-bit integer values. For example, the rune literal 'a' is actually the number 97.

7. When you pressed Esc, return(enter) and so on, the key is not 0 with char of them being 0. 

8. keyboard.KeyEsc will detect enter key. 


#### About defer keyword
What if there is more than one "defer" in a function? Well, they all run, but they run in the **reverse order** that they are declared. Here's an example:

```go
  func main() {
    defer fmt.Println("Goodbye")
    defer fmt.Println("Farewell")
    fmt.Println("Hello, World!")
  }
```
And here's the output:
	Hello, World!
	Farewell
	Goodbye


#### What about then error handling? 
Go has a different philosophy toward error handling.<br>
In Go, there are no "try / catch" blocks. <br>
Instead, there are errors that CAN be dealt with (an "error" returned from a function)<br> 
and errors that CANNOT be dealt with (a "panic" in Go).


#### "defer" and "panic" 
The reason this is important to us is that we need to know what happens to deferred items when a panic happens.<br>
The good news is that "defer" still runs.<br><br>

If a function panics, any defers in that function will still run.<br>
If that function was called by another function, defers will run in that calling function.<br>
This goes all the way up to the main goroutine.<br>
So if a function panics, all of the deferred items will still run.<br>


## Changing the program
So far, this only prints out the number of char(superset of ASCII) which is essentially a number.<br>
Let's change it to display a "menu" that the user can choose from with a single key press. 

```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open() 
	if err != nil {        
		log.Fatal(err) 

	defer func() { 
		_ = keyboard.Close() 
	}()

	
	coffees := make(map[int]string) // dictionary!
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program") //Instead of ESC, now we are looking Q being pressed

	for {
		char, _, err := keyboard.GetSingleKey() 
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' { //2 
			break
		}

		//How can you convert the rune to human-readable character? - By the following!
		t := fmt.Sprintf("You chose %q", char) // This returns a string. fmt.Sprintf replaces %q with char
		fmt.Println(t)

		//What about converting string to integer?
		i, _ := strconv.Atoi(string(char))
		t2 := fmt.Sprintf("You chose %d", i) //placeholder for integer is 'd'
		fmt.Println(t2)

		t3 := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t3)

		//What if you press things that are not numeric? -> it will return 0

	}
	fmt.Println("Program exiting...")
}
```
1.  maps are specialized data structure so you can't just say "var coffees = map()". Instead, you declare it:
```go
variable = make(map[string]int) // [type_of_key]type_of_value
```

2. Unlike python, Go distinguishes "" and ''. "" is used for string whereas '' is for rune
