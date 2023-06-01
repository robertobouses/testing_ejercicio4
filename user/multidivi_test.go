package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio4/user"
)

func TestMultiplicacionDivision(t *testing.T) {
	expected := 100 // Expected value for multiplication
	var result int
	opcion := "M"

	result = user.MultiplicacionDivision(10, 10, opcion)
	if result != expected {
		t.Errorf("Error in multiplication. Expected %v, but got %v", expected, result)
	}

	expected = 1 // Expected value for division
	opcion = "D"
	result = user.MultiplicacionDivision(10, 10, opcion)
	if result != expected {
		t.Errorf("Error in division. Expected %v, but got %v", expected, result)
	}
}
