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


### Real Introduction to Go
Let's go through a little more about doctor.go package.<br><br>

Top of the file:
```go
package doctor //1

import ( //2
    "log"
    "math/rand"
    "regexp"
    "strings"
    "time"
)

/* This code is a port of the original  "Eliza" program to Go. The program appears to have a conversation with the user,
using string replacement and basic parsing, to give the appearance of understanding what the user types. The
responses are based on a Python program by Joe Strout, Jeff Epler and Jez Higgins. The Python code is available at
https://github.com/jezhiggins/eliza.py. This code is licensed under the terms of the MIT License.
*/ //3

```
1. package declaration. As you might've noticed, when we added the doctor.go file to our project, we put them in a folder named "doctor", which is convention.  
2. Most Go files will import packages either from standard libraries or from packages that other developers(or we) define.
3. Multiline comments (//* *//)


```go
var matches = []string{ //1
    "life",
    "i need",
    "why don't",
    "why can't",
    "i can",
    "i am",
...
}

var reflections = map[string]string{ //2
    "am":     "are",
    "was":    "were",
    "i":      "you",
    "i'd":    "you would",
    "i'll":   "you will",
    "my":     "your",
    "are":    "am",
    "you've": "I have",
    "your":   "my",
    "yours":  "mine",
    "you":    "me",
    "i'm":    "your",
}


var responses = [][]string{ //3
    {"Life? Don't talk to me about life.", "At least you have a life, I'm stuck inside this computer.", "Life can be good. Remember, 'this, too, will pass'."},
    {"Why do you need %1?", "Would it really help you to get %1?", "Are you sure you need %1?"}, //4
    {"Do you really think I don't %1?", "Perhaps eventually I will %1.", "Do you really want me to %1?"}
}
```
1. This variable declaration 'matches' doesn't have the plain string types. This is a slice of strings

2. Then we have another variable called 'reflections'. This is not a string but entirely new time called a **map**.

3. A slice of slice of strings. So the first entry in the slice itself is a slice having 3entries. 

4. this '%1' means what? 


```go
func Intro() string {
    return ` 
I'm Eliza
---------
Talk to the program by typing in plain English, using normal upper
and lower-case letters and punctuation.  Enter 'quit' when done.

Hello. How are you feeling today?` //1
}
```
1. Here, for return we didn't use double quote(") but backtick(`). It allows you to construct a string that spans multiple lines. 
<br><br>


```go
// Response builds a response based on input and sends back string
func Response(userInput string) string {
    // declare the two strings we need for output
    var output, remainder string

    // seed the random number generator.
    rand.Seed(time.Now().UnixNano())

    // Sanitize user input by setting up a regular expression to strip out punctuation
    reg, err := regexp.Compile("[^a-zA-Z0-9]+")
    if err != nil {
        log.Fatal(err)
    }

    // strip out punctuation from user input
    userInput = reg.ReplaceAllString(userInput, " ")

    // Loop through the matches list. If there's a match, strip it out. Change words in the remainder (if any)
    // of the input with the corresponding entry from the reflections map
    for i := 0; i < len(matches); i++ {
        match := matches[i]
        position := strings.Index(strings.ToLower(userInput), match)

        if position > -1 {
            // we found the word in matches in the user input string, so now we need to
            // figure out how much to delete that input
            tmp := strings.ToLower(userInput)[position+len(match):]
            tmp = strings.Trim(tmp, " ")
            exploded := strings.Split(tmp, " ")

            // loop through slice of words in our exploded variable, looking for one that
            // matches the key in the reflections map. This will change pronouns and verbs into words
            // appropriate for our response.
            for index, word := range exploded {
                for key, value := range reflections {
                    if strings.ToLower(key) == strings.ToLower(word) {
                        // we found one, so swap the old value for the one
                        // from our reflections map and break out of this loop.
                        exploded[index] = value
                        break
                    }
                }
            }

            // turn the slice of words into a single string, and strip off extra spaces from beginning/end
            remainder = strings.Join(exploded, " ")
            remainder = strings.Trim(remainder, " ")

            // get a random number to use to pick a random response
            randomIndex := rand.Intn(len(responses[i]))
            // assign randomly selected response to output
            output = responses[i][randomIndex]
            break
        }

    }

    // If there wasn't a match, use the last item in the responses slice.
    if output == "" {
        randomIndex := rand.Intn(len(responses[len(responses)-1]))
        output = responses[len(responses)-1][randomIndex]
    }

    // Build our final response and send it back. If the response contains %1, replace that with the remainder
    // of the input string.
    output = strings.ReplaceAll(output, "%1", remainder) //1

    return output
}

```
1. This is where you replace %1 with remainder in output. 


### Building program to an executable file
Instead of executing "go run main.go" every single time, you may want to build program in an executable file.<br><br>

**On Mac or Linux**
```sh
go build -o name_of_file main.go
```

**On Windows**
```shell
go build -o name_of_file.exe main.go
```

**For linux(ubuntu)**
```shell
env GOOS=linux GOARCH=amd64 go build -o eliza main.go
```