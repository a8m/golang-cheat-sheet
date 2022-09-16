// There is no exception handling. Instead, functions that might produce an error just declare an additional return value of type error. This is the error interface:

// The error built-in interface type is the conventional interface for representing an error condition,
// with the nil value representing no error.
type error interface {
    Error() string
}

// Here's an example
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func main() {
	val, err := sqrt(-1)
	if err != nil {
		// handle error
		fmt.Println(err) // negative value
		return
	}
	// All is good, use `val`.
	fmt.Println(val)
}
