package main
import "fmt"

func main() {
	var doesUniverseExist bool
	doesUniverseExist = true
	fmt.Printf("Data = %v, Type = %T", doesUniverseExist, doesUniverseExist);

	var numberOfStars uint
	numberOfStars = 456678
	fmt.Printf("Data = %v, Type = %T", numberOfStars, numberOfStars);

	var complexNumber complex64
	complexNumber = complex(2, 3)
	fmt.Printf("Data = %v, Type = %T", complexNumber, complexNumber);

	var weight float32
	weight = 17.23
	fmt.Printf("Data = %v, Type = %T", weight, weight);
}
