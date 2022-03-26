### Hello world!
The very first line in any file must be a package declaration<br>
And the convention for your main part of your program is to name the package "main"<br>

```go
package main
```
<br>

and every main package has to have a functio named main

```go
package main
import "fmt"
func main() {
	//fmt itself is a package that is available in GO standard library
	fmt.Println("Hello, world!")
}
```

#### Running go program
To run go program, type the follow on terminal
```sh
go run main.go
```



### Structure of a Go Program
```go
package main
import "fmt"

func main(){
    sayHelloWorld("Hello, world again")
}

func sayHelloWorld(whatToSay string){ //1
    fmt.Println(whatToSay)
}
```
1. This is how you define the type of argument


### Variables and Dot Notation
variable is as in other programming languages,<br>
nothing more than a place to store things.

```go
package main
import "fmt"

func main(){
    //var whatToSay string //1
    //whatToSay = "Hello world, again!"//2

    whatToSay := "Hello world, again!" //3
    sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string){
    fmt.Print.ln(whatToSay)
}
```
1. This is how you define the type of variable. 

2. Go is quite picky bout variables and it won't let you have any "unused" variables in your program. So you have to assign a value and do something with it. 

3. This is a shorthand that says to Go, 
- "I want a new variable called whatToSay" 
- "I want you to store value on the right side in that variable" 
- "I want you do figure out what type this is"

 
 ### Setting up the program 
 If you want your program to be "importable" like "fmt", you can use "Eliza".  

 ```sh
go mod init myapp # you can call whatever you want
 ```
 
 If you run the command above, you will see "go.mod" file. <br><br>

Suppose you have subdirectory called **doctor** with **doctor.go** file inside of it.<br><br>

Inside this file you have a function called intro that returns string as follows:
```go
package doctor

func Intro() string {
    return `
I'm Eliza
---------
Talk to the program by typing in plain English, using normal upper
and lower-case letters and punctuation.  Enter 'quit' when done.

Hello. How are you feeling today?`
}
```
<br>

If you want to get the return value of this function, what you do is:

```go
package main 

import ( //1
	"fmt"
	"myapp/doctor" //2
)

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}
```
1. When you have mulitple modules to be imported, use this. 
2. We named our program as "myapp"  

#### Structure of Go 
- package: As recall, every Go **"file"** must have a package declaration. 
- import : And it is followed by import statement.
- main function : every Go **"program"** has to have main function as well. Note that this is not talking about Go files. So you need just one main function for the entire program. 


### User input
```go
import (
    "bufio" //1
    "fmt"
    "myapp/doctor"
    "os"  //2 
)

func main(){
    reader := bufio.NewReader(os.Stdin) //2 

    userInput, _ := reader.ReadString('\n') //3 

    fmt.Println(userInput)
}

```
1. You import bufio package to get user input 

2. Where you are getting your input? What should you be reading? - In this case it's from os's standard input. 

3. ReadString allows us to capture whatever the user types. But you have to be able to tell when the user is done typing - which is in our case when user presses the enter('\n'). 
<br>

### Allowing users to type multiple things in(For loop)
For this, we need to repeat getting user inputs and printing them into the screen.<br><br>

In Go, there is only one kind of loop.(for loop)<br> 
The syntax of that is as follows:
```go
func main() {
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println(doctor.Intro())
	for {
		userInput, _ := reader.ReadString('\n') 
		fmt.Println(userInput)
	}

}
```
With this, it will execute forever as it has no condition to be finished.<br><br>

Well, it works as expected but we don't want the program to simply echo back all the time. Let's get some meaningful response from Eliza.(from one of the function in doctor.go<br><br>

```go
func main() {
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println(doctor.Intro())
	for {
        fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') 
		response := doctor.Response(userInput)
        fmt.Println(response)
	}
}
```

### Trapping for a certain word (if statement)
```go
func main() {
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println(doctor.Intro())
	for {
        fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') 
        if userInput == "quit"{ //1
            break  //2
        }else{
            response := doctor.Response(userInput)
            fmt.Println(response)
        }
	}
}
```
1. if statement! note that there is no open parenthesis for condition
2. break to get out of loop.

#### Caveat : User input is not actually just 'quit'
When you type 'quit' to get out of the loop, you'll see it does NOT actually work. And that's because there is the actual return that's been appended to letters quit.<br><br>

To fix this, you need to strip off the part where the user press the enter(carriage return). And make sure you do that twice : once for windows and once for everything else.
```go
func main() {
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println(doctor.Intro())
	for {
        fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') 

        userInput = strings.Replace(userInput,"\r\n","", -1 ) //1 
        userInput = strings.Replace(userInput,"\n","",-1) //2

        if userInput == "quit"{ 
            break  
        }else{
            response := doctor.Response(userInput)
            fmt.Println(response)
        }
	}
}
```
1. In windows, return key is a carriage return(\\r) and a line feed(\\n). You want to replace \\r\\n with empty string "", and '-1' here means every occurrence of the letters. 

2. This is for other OSs like macOs or linux. 




