//  package main
package main 

// formatter imported
import "fmt"


// split function declarations
func split(sum int) (x, y int) {
	// x defined
	x = sum * 54 / 2
	// y defined
	y = sum - x
	//  Returns values
	return
}

// Main Function starts here
func main(){
	// format and print in new line the split value for the number provided
	fmt.Println(split(17))
}