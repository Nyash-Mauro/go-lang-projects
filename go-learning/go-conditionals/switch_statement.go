package main
import "fmt"

func main(){
	var alphabet string = "A"
	switch alphabet {
		case "A":
			fmt.Println("Apple")
		case "B":
			fmt.Println("Ball")
		case "C":
			fmt.Println("Cat")
		case "D":
			fmt.Println("Doll")
		default:
			fmt.Println("EZ!")
	}
}