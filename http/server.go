package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/testing_ejercicio4/user"
)

type Server struct {
	engine  *gin.Engine
	handler Handler
}

type Handler struct {
}

/*
	Pais1 := user.Pais{
		Name: "Portugal",

Population: 38
Continent: "Europa"
}

	Pais2:= user.Pais{
		Name: "Japón",

Population: 102
Continent: "Asia"
}

	Pais3:= user.Pais{
		Name: "México",

Population: 174
Continent: "América"
}

	Pais2:= user.Pais{
		Name: "Camerún",

Population: 93
Continent: "África"
}
*/
func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	router.GET("/suma", handler.PaginaSuma)
	router.GET("/resta", handler.PaginaResta)
	router.GET("/multidivi", handler.PaginaMultiDivi)
	router.GET("/pais", handler.PaginaPais)

	return Server{
		engine:  router,
		handler: handler,
	}
}

func (h Handler) PaginaSuma(c *gin.Context) {
	valor := user.Suma(10, 5)
	response := fmt.Sprintf("El resultado de la suma es: %d", valor)
	c.String(http.StatusOK, response)
}

func (h Handler) PaginaResta(c *gin.Context) {
	valor := user.Resta(10, 5)
	response := fmt.Sprintf("El resultado de la resta es: %d", valor)
	c.String(http.StatusOK, response)
}

func (h Handler) PaginaMultiDivi(c *gin.Context) {
	opcion := "M"
	var response string
	valor := user.MultiplicacionDivision(10, 5, opcion)
	if opcion == "M" {
		response = fmt.Sprintf("El resultado de la multiplicación es: %d", valor)
	} else {
		response = fmt.Sprintf("El resultado de la división es: %d", valor)
	}
	c.String(http.StatusOK, response)
}

func (h Handler) PaginaPais(c *gin.Context) {
	pais := user.Pais{
		Name:       "Portugal",
		Population: 38,
		Continent:  "Europa",
	}

	// Capture the output of BigPais() and PlacePais() in variables
	bigPaisOutput := captureOutput(pais.BigPais)
	placePaisOutput := captureOutput(pais.PlacePais)

	// Print the captured outputs
	fmt.Printf("BigPais Output: %s\n", bigPaisOutput)
	fmt.Printf("PlacePais Output: %s\n", placePaisOutput)

	c.String(http.StatusOK, "Pais:", pais.Name, bigPaisOutput, placePaisOutput)
}

// Helper function to capture the output of a function
func captureOutput(fn func()) string {
	output := ""
	// Redirect standard output to capture the function's output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn() // Call the function

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old
	output = string(out)

	return output
}

func (s *Server) Run(puerto string) error {
	return s.engine.Run(puerto)
}
