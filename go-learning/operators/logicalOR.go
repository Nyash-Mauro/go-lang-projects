// logical OR (||)
package main
import "fmt"

func main(){
	var isPerson bool = false
	var isProgrammer bool = true
	var isPython bool = true

	if (isPerson || isProgrammer) {
		isPython = false
		fmt.Println("isPerson =", isPerson)
		fmt.Println("isProgrammer =", isProgrammer)
		fmt.Println("isPython =", isPython)
	}
}