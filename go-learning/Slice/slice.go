package main
import "fmt"

func main(){
	var i [5]int
	var sliced []int = i[1:3]

	fmt.Println(sliced)
}