package main
import "fmt"

type Planet struct {
	name string
	galaxyname string
	numberOfMoons int
}

type Rotater interface {
	Rotate(degrees float32) string
}

func (planet Planet) Rotate(degrees float32) string{
	return fmt.Sprintf("Planet %s is rotating at %f degrees", planet.name, degrees)
}

func main() {
	var rotatable Rotater
	rotatable = Planet{name: "Earth", galaxyname: "MilkyWay", numberOfMoons: 1}
	fmt.Println(rotatable.Rotate(70.0))

	rotatable = Planet{name: "Jupiter", galaxyname: "MilkyWay", numberOfMoons: 21}
	fmt.Println(rotatable.Rotate(60.3))
}