// IF
if x > 10 {
    return x
} else if x == 10 {
    return 10
} else {
    return -x
}

// You can put one statement before the condition
if a := b + c; a < 42 {
    return a
} else {
    return a - 42
}

// Type assertion inside if
var val interface{} = "foo"
if str, ok := val.(string); ok {
    fmt.Println(str)
}

// LOOPS, There's only `for`, no `while`, no `until`
for i := 1; i < 10; i++ {
}
for ; i < 10;  { // while - loop
}
for i < 10  { // you can omit semicolons if there is only a condition
}
for { // you can omit the condition ~ while (true)
}

// use break/continue on current loop
// use break/continue with label on outer loop
here:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if i == 0 {
                continue here
            }
            fmt.Println(j)
            if j == 2 {
                break
            }
        }
    }

there:
    for i := 0; i < 2; i++ {
        for j := i + 1; j < 3; j++ {
            if j == 1 {
                continue
            }
            fmt.Println(j)
            if j == 2 {
                break there
            }
        }
    }

// SWITCH
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

// you can also make comparisons in switch cases
number := 42
switch {
    case number < 42:
        fmt.Println("Smaller")
    case number == 42:
        fmt.Println("Equal")
    case number > 42:
        fmt.Println("Greater")
}

// cases can be presented in comma-separated lists
var char byte = '?'
switch char {
    case ' ', '?', '&', '=', '#', '+', '%':
        fmt.Println("Should escape")
}
