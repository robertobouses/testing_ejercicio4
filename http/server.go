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

	router.GET("/", handler.Index)
	return Server{
		engine:  router,
		handler: handler,
	}
}

func (h Handler) Index(c *gin.Context) {
	valor := user.Suma(5, 3)
	response := fmt.Sprintf("El resultado de la suma es: %d", valor)
	c.String(http.StatusOK, response)
}

func (s *Server) Run(puerto string) error {
	return s.engine.Run(puerto)
}
