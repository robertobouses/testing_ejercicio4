package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio4/user"
)

func TestResta(t *testing.T) {
	expected := 10
	value := user.Resta(12, 2)
	if expected != value {
		t.Errorf("el valor obtenido es %d, y debería ser %d", value, expected)
	}
}
