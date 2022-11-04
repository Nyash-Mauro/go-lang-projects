// Logical AND (&&):

package main
import "fmt"

func main() {
	var isPlanet bool = false
	var isComet bool = false
	var isMeteor bool = false

	if(!isPlanet && !isComet) {
		isMeteor = true
		fmt.Println("isPlanet = ", isPlanet)
		fmt.Println("isComet = ", isComet)
		fmt.Println("isMeteor = ", isMeteor)
	}
}