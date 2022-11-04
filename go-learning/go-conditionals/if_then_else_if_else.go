package main
import "fmt"

func main() {
	if accountBalance := 1000001; accountBalance < 1000 {
		fmt.Println("Close Account!")
	} else if(accountBalance > 1000000) {
		fmt.Println("Please find a Europe tour cruise package in your mailbox.")
	} else {
		fmt.Println("We love having you with us!")
	}
}