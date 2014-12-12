# Go Cheat Sheet

## Credits

Most example code taken from [A Tour of Go](http://tour.golang.org/), which is an excellent introduction to Go.
If you're new to Go, do that tour. Seriously.

## Go in a Nutshell

* Imperative language
* Statically typed
* Syntax similar to Java/C/C++, but less parantheses and no semicolons
* Compiles to native code (no JVM)
* No classes, but structs with methods
* Interfaces
* No implementation inheritance. There's [type embedding](http://golang.org/doc/effective%5Fgo.html#embedding), though.
* Functions are first class citizens
* Functions can return multiple values
* Go has closures
* Pointers, but not pointer arithmetic
* Built-in concurrency primitives: Goroutines and Channels

# Basic Syntax

## Hello World
File `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```
`$ go run hello.go`


## Declarations
* Type goes after identifier! 
* `var foo int // declaration without initialization`
* `var foo int = 42 // declaration with initialization`
* `var foo, bar int = 42, 1302 // declare and init multiple vars at once`
* `var foo = 42 // type omitted, will be inferred`
* `foo := 42 // shorthand, only in func bodies, omit var keyword, type is always implicit `
* `const constant = "This is a constant"`

## Functions
```go
// a simple function
func functionName() {}

// function with parameters (again, types go after identifiers)
func functionName(param1 string, param2 int) {}

// multiple parameters of the same type
func functionName(param1, param2 int) {}

// return type declaration
func functionName() int {
    return 42
}

// Can return multiple values at once
func returnMulti() (int, string) {
    return 42, "foobar"
}
var x, str = returnMulti()

// Return multiple named results simply by return
func returnMulti2() (n int, s string) {
    n = 42
    s = "foobar"
    // n and s will be returned
    return
}
var x, str = returnMulti2()

```

### Functions As Values And Closures
```go
func main() {
    // assign a function to a name
    add := func(a, b int) int {
        return a + b
    }
    // use the name to call the function
    fmt.Println(add(3, 4))
}

// Closures: Functions can access values that were in scope when defining the
// function

// adder returns an anonymous function with a closure containing the variable sum
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x // sum is declared outside, but still visible
        return sum
    }
}
```

### Variadic Functions
```go
func main() {
	fmt.Println(adder(1, 2, 3)) // return 6
	fmt.Println(adder(9, 9)) // return 18
}

// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
// The function is invoked like any other function except we can pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	for _, v := range args { // Iterates over the arguments whatever the number.
		total += v
	}
	return total
}
```

## Built-in Types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 ~= a character (Unicode code point) - very Viking

float32 float64

complex64 complex128
```

## Type Conversions
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// alternative syntax
i := 42
f := float64(i)
u := uint(f)
```

## Packages 
* package declaration at top of every source file
* executables are in package `main`
* convention: package name == last name of import path (import path `math/rand` => package `rand`)
* upper case identifier: exported (visible from other packages)
* Lower case identifier: private (not visible from other packages) 

## Control structures

### If
```go
    if x > 0 {
        return x
    } else {
        return -x
    }

    // You can put one statement before the condition
    if a := b + c; a < 42 {
        return a
    } else {
        return a - 42
    }
```

### Loops
```go
    // There's only `for`, no `while`, no `until`
    for i := 1; i < 10; i++ {
    }
    for ; i < 10;  { // while - loop
    }
    for i < 10  { // you can omit semicolons if there is only a condition
    }
    for { // you can omit the condition ~ while (true)
    }
```

### Switch
```go
    // switch statement
    switch operatingSystem {
    case "darwin":
        fmt.Println("Mac OS Hipster")
        // cases break automatically, no fallthrough by default
    case "linux":
        fmt.Println("Linux Geek")
    default:
        // Windows, BSD, ...
        fmt.Println("Other")
    }

    // as with for and if, you can have an assignment statement before the switch value 
    switch os := runtime.GOOS; os {
    case "darwin": ...
    }
```

## Arrays, Slices, Ranges

### Arrays
```go
var a [10]int // declare an int array with length 10. Array length is part of the type!
a[3] = 42     // set elements
i := a[3]     // read elements

// declare and initialize
a := [2]int{1, 2}
a := [...]int{1, 2} // elipsis -> Compiler figures out array length
```

### Slices
```go
var a []int // declare a slice - similar to an array, but length is unspecified
var a = []int {1, 2, 3, 4} // declare and initialize a slice (backed by the array given implicitly)
a := []int{ 1, 2, 3, 4 } // shorthand


var b = a[lo:hi] // creates a slice (view of the array) from index lo to hi-1
var b = a[1:4] // slice from index 1 to 3
var b = a[:3] // missing low index implies 0
var b = a[3:] // missing high index implies len(a)

// create a slice with make
a = make([]byte, 5, 5) // first arg length, second capacity
a = make([]byte, 5) // capacity is optional

```

### Operations on Arrays and Slices
`len(a)` gives you the length of an array/a slice. It's a built-in function, not a attribute/method on the array.

```go
// loop over an array/a slice
for i, e := range a {
    // i is the index, e the element
}

// you'll get a compiler error if you're not using i and e. If you only need e:
for _, e := range a {
    // e is the element
}

// ...and if you only need the index
for i := range a {
}

```

## Maps

```go
var m map[string]int
m = make(map[string]int)
m["key"] = 42
fmt.Println(m["key"])

delete(m, "key")

elem, ok := m["key"] // test if key "key" is present and retrieve it, if so

// map literal
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

```

## Structs

There are no classes, only structs. Structs can have methods.
```go
// A struct is a type. It's also a collection of fields 

// Declaration
type Vertex struct {
    X, Y int
}

// Creating
var v = Vertex{1, 2}

// Accessing members
v.X = 4

// You can declare methods on structs. The struct you want to declare the
// method on (the receiving type) comes between the the func keyword and
// the method name. The struct is copied on each method call(!)
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Call method
v.Abs()

// For mutating methods, you need to use a pointer (see below) to the Struct
// as the type. With this, the struct value is not copied for the method call.
func (v *Vertex) add(n float64) {
    v.X += n
    v.Y += n
}

```

## Pointers
```go
p := Vertex{1, 2}  // p is a Vertex
q := &p            // q is a pointer to a Vertex
r := &Vertex{1, 2} // r is also a pointer to a Vertex

// The type of a pointer to a Vertex is *Vertex

var s *Vertex = new(Vertex) // new creates a pointer to a new struct instance 
```

## Interfaces
```go
// interface declaration
type Awesomizer interface {
    Awesomize() string
}

// types do *not* declare to implement interfaces
type Foo struct {}

// instead, types implicitly satisfy an interface if they implement all required methods
func (foo Foo) Awesomize() string {
    return "Awesome!"
}
```

## Type Embedding
TODO

## Errors
There is no exception handling. Functions that might produce an error just declare an additional return value of type `Error`. This is the `Error` interface:
```go
type error interface {
    Error() string
}
```

A function that might return an error:
```go
func doStuff() (int, error) {
}

func main() {
    result, error := doStuff()
    if (error != nil) {
        // handle error
    } else {
        // all is good, use result
    }
}
```

# Concurrency

## Goroutines
Goroutines are lightweight threads (managed by Go, not OS threads). `go f(a, b)` starts a new goroutine which runs `f` (given `f` is a function).

```go
// just a function (which can be later started as a goroutine)
func doStuff(s string) {
}

func main() {
    // using a named function in a goroutine
    go doStuff("foobar")

    // using an anonymous inner function in a goroutine
    go func (x int) {
        // function body goes here
    }(42)
}
```

## Channels
```go
ch := make(chan int) // create a channel of type int
ch <- 42             // Send a value to the channel ch.
v := <-ch            // Receive a value from ch

// Non-buffered channels block. Read blocks when no value is available, write blocks if a value already has been written but not read.

// Create a buffered channel. Writing to a buffered channels does not block if less than <buffer size> unread values have been written.
ch := make(chan int, 100)

close(c) // closes the channel (only sender should close)

// read from channel and test if it has been closed
v, ok := <-ch

// if ok is false, channel has been closed

// Read from channel until it is closed
for i := range ch {
    fmt.Println(i)
}

// select blocks on multiple channel operations, if one unblocks, the corresponding case is executed
func doStuff(channelOut, channelIn chan int) {
    select {
    case channelOut <- 42:
        fmt.Println("We could write to channelOut!")
    case x := <- channelIn:
        fmt.Println("We could read from channelIn")
    }
}
```
# Snippets

## HTTP Server
```go
package main

import (
    "fmt"
    "net/http"
)

// define a type for the response
type Hello struct{}

// let that type implement the ServeHTTP method (defined in interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}

// Here's the method signature of http.ServeHTTP:
// type Handler interface {
//     ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
```
