package main

import "fmt"

// Here, Type Assertions are being used to extract the underlying value of an interface variable.
// Since the value of interface is a string, it works perfectly for first two instances
// In third instance, since the value of interface is being asserted to a float the assertion fails, but due to the comma ok assertion it assigns 0 to f and value of ok is false
// In the fourth instance, since there is not a comma ok assertion the program panics
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
