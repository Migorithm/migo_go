### Scope
```go

var mynum = "Three" //3
func main() {
    var one = "this is a block level variable" //4 
	var one = "One"  // 1
	fmt.Println(one)

	myFunc()
}

func myFunc() {
	var one = "the number one" //2
	fmt.Println(one)
    fmt.Println(mynum)
}
```
1,2 : There are block level variable, available inside the function where they're declared. 
3 : **package level** variable. 


### Public/private var
- Private var that starts The starting **lowercase** menas it is only available inside the package
- Public var is available to any other package that import package one.

Every time you start a new Go project, you should create module by:
```sh
go init mod <module_name>
```

**packageone/package.go**
```go
package packageone
var privateVar = "I am private"//1
var PublicVar = "I am public (for exported)"//2

func notExported() {  //3

}

func Exported() { //4
}

```
1. This is available as it starts with capital letter.
2. This is not available here. 
3. Not available for the same reason as first one
4. Available anywhere

<br><br>

main.go:
```go
package main
import (
	"fmt"
	"myapp/packageone"
)
var mynum = "Two"
func main() {
	var mynum = "this is a block level variable."
	//var one = "One"
	fmt.Println(mynum)
	myFunc()
	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)
	packageone.Exported()
	//packageone.notExported()
	fmt.Println(packageone.PublicVar)
}
func myFunc() {
	var one = "the number one"
	fmt.Println(one)
	fmt.Println(mynum)
}
```

### Scope challenge
Let's delete everything in main.go and packageone/package.go and do the following tasks
```go
package main

func main(){
	//variables
	//declare a package level variable for the main pakcage named myVar

	//declare a block level variable for the main function called blockVar

	//declare a package level variable in the packageone package named PackageVar

	//create an exported function in packageone called PrintMe

	// in the main function, print out the values of myVar, blockVar, and PackageVar
	// on one line using the PrintMe function in packageone
}
```


#### Solution
```go
package main

import (
	"myapp/packageone"
)

var myVar = "This is a package level variable" //1

func main() {

	var blockVar = "This is the block level variable"
	packageone.PrintMe(myVar, blockVar)

}

```
1. This can NOT be 'myVar:="This is..."' 


```go
package packageone

import "fmt"
var PackageVar = "Package level Variable from packageone" //2

func PrintMe(s1,s2 string ){
	fmt.Println(s1,s2,PackageVar)
}
```
1. This can NOT be 'PackageVar:="Package level Variable from packageone"'

