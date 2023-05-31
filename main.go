package main

import (
	"fmt"

	"github.com/robertobouses/testing_ejercicio4/http"
)

func main() {
	fmt.Println("HELLOO")

	server := http.NewServer()
	server.Run(":8080")
}
