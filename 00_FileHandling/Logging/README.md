### Logging
The purpose of logging is to get a trace of what the program is doing, where it's happening and when.<br>
For example:
```
TRACE: 2009/11/10 23:00:00.0000000 /tmpfs/gosandbox-/prog.go:14: message
```
The above message contains:
- prefix(TRACE)
- datatime stamp
- full path to the source code writing to the log
- the line of code 
- message
<br>

Let's look at a program that allows you to configure the log package.
```go 
package main

import (
    "log"
)

func init(){
    log.SetPrefix("TRACE: ")
    log.SetFlags(log.Ldata | log.Lmicroseconds | log.Llongfile)
}

func main() {
    //Println writes to the standard logger.
    log.Println("message")

    // Fatalln is Println() followed by a call to os.Exit(1) so when you call this, the program will end. 
    log.Fatalln("fatal message")

    // panicln is Println() followed by a call to panic(), so this will iuse a panic which unless recovered, will cause the program to terminate and stack trace
    log.Panicln("panic message")
}
```

When you run the program, you should get something close to the example above.<br>
Let's break down the code and see how it works.<br>
```go
func init(){
    log.SetPrefix("TRACE: ")
    log.SetFlags(log.Ldata | log.Lmicroseconds | log.Llongfile)
}
```
<br>

**This init() function is executed before main()** as a part of program initialization.<br><br>

It's common to set the log configration inside of init() so the log package can be used immediately<br><br>

when the program starts.<br><br>


#### Flags
There are several flags associated with the log package:
```go
const (
    //these items: 
        // 2009/01/23  01:23:23.123123 /a/b/c/d.go:23: message

    // the data : 2009/01/23
    Ldata = 1 << iota

    // the time : 01:23:23
    Ltime

    // microsecond resolution : 01:23:23.123123
    Lmicroseconds

    // full file name and line number: /a/b/c/d.go:23
    Llongfile

    // final file name element and line number: d.go:23
    // this overrides Llongfile
    Lshortfile

    //initial values for the standard logger
    LstdFlags = Ldate | time

)
```
The first constant in this block is called Ldata, which is declared with special syntax: "iota"<br><br>

The **iota** keyword has a special purpose when it comes to declaring a block of constants.<br><br>

It instructs the compiler to duplicate the expression for every constant until the block ends or an assignment statement is found.<br><br>

Plus, the value of iota for each preceding constant gets incremented by 1. Let's take a look at this more closely.<br>
```go
const (
    Ldata =1 <<iota //1 << 0 = 000000001 =1
    Ltime           //1 << 1 = 000000010 =2
    Lmicroseconds   //1 << 2 = 000000100 =4
    Llongfile       //1 << 3 = 000001000 =8
    Lshortfile      //1 << 4 = 000010000 =16 
)
```
This is what's happening behind the scenes.<br><br>

<< operator performs a left bitwise shift and in each case the bit pattern for the value of 1 is shifted to the left iota position.<br><br>

The net effect of this is givving each constant its own unique bit position, which is perfect when working with flags.<br><br>

The LstdFlags constant shows the purpose behind giving each constant its own unique bit position.<br>
```go
const (
    ...
    LstdFlags = Ldate(1) | Ltime(2) = 00000011 = 3 
)
```
You see LstdFlags constant is assigned the value of 3, thanks to the fact that tthe pipe operator(|) is used to "or bits" together.<br>
So this technique can be used when you want to set the flags that are to be applied. 
```go
func init(){
    ...
    log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile )
}
```

#### loggers are multigoroutine-safe
Multiple goroutines can call these functions from the same logger value at the same time without the writes colliding with each other.<br>

The standard logger and any customized logger you may create will have this attribute. 


### Customized loggers
Creating customized loggers require that you create your own Logger type values.<br>
Each logger you create can be configured for a unique destination and set with its own prefix and flags.<br>
Let's look at an example program that creates different Logger type pointer variables to support different logging levels.<br>
```go
package main

import (
    "io"
    "io/ioutil"
    "log"
    "os"
)

var (
    Trace   *log.Logger //just about anything
    Info    *log.Logger //Important information
    Warning *log.Logger //Be concerned
    Error   *log.Logger //Criticcal problem
)

func init(){
    file, err := os.Openfile("errors.txt",
        os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil{
        log.Fatalln("Failed to open error log file:" , err)
    }

    Trace = log.New(ioutil.Discard, 
            "TRACE: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Info = log.New(os.Stdout,
            "INFO: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Warning = log.New(os.Stdout,
            "WARNING: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Error = log.New(io.MultiWriter(file, os.Stderr), //file object 
            "ERROR: ",
            log.Ldate | log.Ltime | log.Lshortfile)
}

func main(){
    Trace.Println("I have something standard to say")
    Info.Println("Special information")
    Warning.Println("There is something you need to know about")
    Error.Println("Something has failed.")
}
```
Here, the program creates four different Logger type pointer variables, each of which is configured differenty.<br><br>

Let's break down the code so you can learn how all this works.<br>

#### Logger type pointer variables
```go
var (
    Trace   *log.Logger //just about anything
    Info    *log.Logger //Important information
    Warning *log.Logger //Be concerned
    Error   *log.Logger //Criticcal problem
)

func init(){
    file, err := os.Openfile("errors.txt",
        os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil{
        log.Fatalln("Failed to open error log file:" , err)
    }

    Trace = log.New(ioutil.Discard, 
            "TRACE: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Info = log.New(os.Stdout,
            "INFO: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Warning = log.New(os.Stdout,
            "WARNING: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Error = log.New(io.MultiWriter(file, os.Stderr), //file object 
            "ERROR: ",
            log.Ldate | log.Ltime | log.Lshortfile)
```
To create each logger, we use the ***New*** function from the log package, which creates a properly initialized Logger type value.<br><br>
As you can imagine, the **New function returns the address to the newly created value**.<br>
golang.org/src/log/log.go:
```go 
//New creates a new logger. The out variable sets the destination to which log data will be written.
func New(out io.Writer, prefix string, flag int) *Logger {
    return &Logger{out:out, prefix:prefix, flag:flag}
}
```
First parameter is the destination we want to logger to write to, the second is the prefix, and the last is log flags.<br><br>

##### Trace - ioutil.Discard
```go
Trace = log.New(ioutil.Discard,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)
```
The Discard variable has some very interesting properties. See the details:<br>
golang.org/src/io/ioutil/ioutil.go:
```go
//devNull is a named type using int as its base type
type devNull int

// Discard is an io.Writer on which all write aclls succeed without doing anything
var Discard io.Writer = devNull(0)

// Implementation of the io.Writer interface
func (devNull) Write(p []byte) (int, error){
    return len(p), nil
}
```
The Discard variable is declared to be of interface type io.Writer and is given a value of 0 of type devNull.<br>
Anything written to this variable is discarded based on the implementation of the Write method for the devNull type.<br>
Simply put, using the Discard variable is a technique you can use to disable a logging level when the output for that level is not required.<br>

#### Info, Warning - os.Stdout
```go
    ...
    Info = log.New(os.Stdout,
            "INFO: ",
            log.Ldate | log.Ltime | log.Lshortfile)
    Warning = log.New(os.Stdout,
            "WARNING: ",
            log.Ldate | log.Ltime | log.Lshortfile)
```
The declaration of the Stdout variable is also interesting.<br>
golang.org/src/os/file.go:
```go
//Stdin, Stdout, Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
var (
    Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

os/file_unix.go
// NewFile returns a new File with the given file descriptor and name 
func NewFile(fd uintptr, name string) *File {
    ...
}
```

#### Error - io.MultiWriter
```go
    Error = log.New(io.MultiWriter(file,os.Stderr), 
            "ERROR: ",
            log.Ldate|log.Ltime|log.Lshortfile)
```
You can see that the first parameter to the New function comes from a special function called MultiWriter from the io package. 

#### Declaration of the different logging methods
The following shows all the methods that have been implemented for the Logger type
```go
func (l *Logger) Fatal(v ...interface{})
func (l *Logger) Fatalf(format string, v ...interface{})
func (l *Logger) Fatalln(v ...interface{})
func (l *Logger) Flags() int
func (l *Logger) Output(calldepth int, s string) error
func (l *Logger) Panic(v ...interface{})
func (l *Logger) Panicf(format string, v ...interface{})
func (l *Logger) Panicln(v ...interface{})
func (l *Logger) Print(v ...interface{})
func (l *Logger) Printf(format string, v ...interface{})
func (l *Logger) Println(v ...interface{})
func (l *Logger) SetFlags(flag int )
func (l *Logger) SetPrefix(prefix string)


```