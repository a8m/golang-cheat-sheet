// Goroutines are lightweight threads (managed by Go, not OS threads). go f(a, b) starts a new goroutine which runs f (given f is a function).

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
