
// In the above example, 
// if the account balance is greater than or equal to 1000, 
// we are printing a warm message.
package main
import "fmt"

func main() {
	if accountBalance := 1001; accountBalance < 1000 {
		fmt.Println("Close Account!")
	} else {
		fmt.Println("We love having you with us!")
	}
}