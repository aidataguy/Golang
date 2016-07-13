// Simple Addition program in GoLang
// Adds package main
package main

// Imports formatting 
import "fmt"

// when two named functions shares same data type omit it from all but last
func add(x, y int) int {
	// function returned
	return x + y
}

// time to run the add function
func main() {
	// formatted line print to output the result of add function from above
	fmt.Println(add(50, 50))
}
// End of the program