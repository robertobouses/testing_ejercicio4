package user

import "fmt"

type Pais struct {
	Name       string
	Population int
	Continent  string
}

func (P Pais) BigPais() {
	if P.Population > 100 {
		fmt.Printf("%s es uno de los mayores países del mundo\n", P.Name)
	} else {
		fmt.Printf("%s no es uno de los países más grandes\n", P.Name)
	}
}

func (P Pais) PlacePais() {
	switch P.Continent {
	case "Europa":
		fmt.Printf("El país %s pertenece al continente %s\n", P.Name, P.Continent)
	case "Asia":
		fmt.Printf("El país %s pertenece al continente %s\n", P.Name, P.Continent)
	default:
		fmt.Printf("%s es un país con el que nuestra empresa no trabaja\n", P.Name)
	}
}
