package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio4/user"
)

func TestSuma(t *testing.T) {
	var expected int
	var result int
	result = user.Suma(3, 5)
	expected = 8
	if result != expected {
		t.Errorf("se esperaba %d, se obtuvo %d", expected, result)
	}
}
