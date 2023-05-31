package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/testing_ejercicio4/user"
)

type Server struct {
	engine  *gin.Engine
	handler Handler
}

type Handler struct {
}

func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	router.GET("/suma", handler.PaginaSuma)
	router.GET("/resta", handler.PaginaResta)
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

func (s *Server) Run(puerto string) error {
	return s.engine.Run(puerto)
}
