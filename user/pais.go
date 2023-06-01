package user

import "fmt"

type Pais struct {
	Name       string
	Population int
	Continent  string
}

func (P Pais) BigPais() {
	if P.Population > 100 {
		fmt.Println("%d es de los mayores países del mundo", P.Name)
	}else {fmt.Println("%d no es un país de los más grandes", P.Name)}
}

func (P Pais)
