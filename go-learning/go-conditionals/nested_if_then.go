package main
import "fmt"

func main() {
	if accountBalance := 501; accountBalance < 1000 {
		if(accountBalance < 500) {
			fmt.Println("Close Account!")
		} else {
			fmt.Println("You better maintain a minimum balance Pal! You got 5 days time")
		}
	} else if(accountBalance > 1000000) {
		fmt.Println("Please find a Europe tour cruise package in your mailbox.")
	} else {
		fmt.Println("We love having you with us!")
	}
}