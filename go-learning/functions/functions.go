// Functions: Functions are set of statements that process input and return output.
// In Go, functions are defined using func keyword. In the below example, 
// we created a function which prints "Hello World!" and 
// then called the function from the main function.




package main
import "fmt"

func userDefinedFunction() {
	fmt.Println("Hello World!")
}

func main() {
	userDefinedFunction()
}