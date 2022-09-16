// ARRAYS
var a [10]int // declare an int array with length 10. Array length is part of the type!
a[3] = 42     // set elements
i := a[3]     // read elements
// declare and initialize
var a = [2]int{1, 2}
a := [2]int{1, 2}   //shorthand
a := [...]int{1, 2} // elipsis -> Compiler figures out array length

// SLICES
var a []int                               // declare a slice - similar to an array, but length is unspecified
var a = []int{1, 2, 3, 4}                 // declare and initialize a slice (backed by the array given implicitly)
a := []int{1, 2, 3, 4}                    // shorthand
chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"]

var b = a[lo:hi]     // creates a slice (view of the array) from index lo to hi-1
var b = a[1:4]       // slice from index 1 to 3
var b = a[:3]        // missing low index implies 0
var b = a[3:]        // missing high index implies len(a)
a = append(a, 17, 3) // append items to slice a
c := append(a, b...) // concatenate slices a and b

// create a slice with make
a = make([]byte, 5, 5) // first arg length, second capacity
a = make([]byte, 5)    // capacity is optional

// create a slice from an array
x := [3]string{"evanka", "avanka", "cevanka"}
s := x[:] // a slice referencing the storage of x

// OPERATIONS ON ARRAYS AND SLICES
// len(a) gives you the length of an array/a slice. It's a built-in function, not a attribute/method on the array.
// loop over an array/a slice
for i, e := range a {
	// i is the index, e the element
}

// if you only need e:
for _, e := range a {
	// e is the element
}

// ...and if you only need the index
for i := range a {
}

// In Go pre-1.4, you'll get a compiler error if you're not using i and e.
// Go 1.4 introduced a variable-free form, so that you can do this
for range time.Tick(time.Second) {
	// do it once a sec
}

