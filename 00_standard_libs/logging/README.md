### Log Package

```
TRACE: 2009/11/10 23:00:00.000000 /tmpfs/gosa/main.go:14: message
```
You see a log entry produced by the log package. This log entry contains:
- prefix
- a datetime stamp
- full path to the source code 
<br>

Let's look at a program that allows you to configure the log package to write<br>
such a line:
```go
package main

import (
	"log"
)

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	log.Panicln("panic message")
}
```
<br>

Let's break down the code:
```go
func init(){
    log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
```
We have a function called init(). This function is exeucted before main() as part of the program initialization. 
