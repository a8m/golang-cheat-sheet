fmt.Println("Hello") // basic print, plus newline
p := struct{ X, Y int }{17, 2}
fmt.Println("My point:", p, "x coord=", p.X)       // print structs, ints, etc
s := fmt.Sprintln("My point:", p, "x coord=", p.X) // print to string variable

fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e", 17, 17, 17, 17.0, 17.0) // c-ish format
s2 := fmt.Sprintf("%d %f", 17, 17.0)                                // formatted print to string variable

hellomsg := `
 "Hello" in Chinese is ('Ni Hao')
 "Hello" in Hindi is ('Namaste')
` // multi-line string literal, using back-tick at beginning and end
