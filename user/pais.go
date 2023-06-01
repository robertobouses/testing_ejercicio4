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

func (P Pais) PlacePais (){
	switch P.Continent{
	case Europa {
		fmt.Println("El pais %d pertenece al contienente %d", P.Name, P.Continent)
	}
case Asia {fmt.Println("El pais %d pertenece al contienente %d", P.Name, P.Continent)}
default{fmt.Println("%d se trata de un país con el que no trabaja nuestra empresa", P.Name)}

}
