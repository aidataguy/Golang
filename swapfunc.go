// package main declared
package main

// importing fmt aka formatter

import "fmt"

// declared Swap Function 
func swap(x, y string) (string, string) {
	return y, x
}

// declared main function
func main() {
	a, b := swap("Hello", "World!") 
	fmt.Println(b, a)
}
