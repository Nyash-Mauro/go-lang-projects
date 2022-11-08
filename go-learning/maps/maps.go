package main
import "fmt"


func main(){
	var countryCapitals map[string]string = map[string]string{
		"Kenya": "Nairobi",
		"Uganda":"Kampala",
		"South Africa":"Johannesburg",
	}
	fmt.Println(countryCapitals)
	fmt.Println(countryCapitals["Kenya"])
}