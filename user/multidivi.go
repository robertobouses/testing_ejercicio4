package user

func MultiplicacionDivision(x, y int, opcion string) int {
	//fmt.Println("indica la operación que quieres hacer (M/D)")
	//fmt.Scanln(opcion)
	if opcion == "M" {
		return x * y
	} else {
		return x / y
	}
}
