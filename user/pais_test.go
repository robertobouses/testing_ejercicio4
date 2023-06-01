package user

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBigPais(t *testing.T) {

	pais := Pais{
		Name:       "Estados Unidos",
		Population: 328,
		Continent:  "América",
	}

	output := captureOutput(func() {
		pais.BigPais()
	})

	expectedOutput := "Estados Unidos es uno de los mayores países del mundo\n"
	if output != expectedOutput {
		t.Errorf("BigPais test failed. Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestPlacePais(t *testing.T) {
	pais := Pais{
		Name:       "España",
		Population: 47,
		Continent:  "Europa",
	}

	output := captureOutput(func() {
		pais.PlacePais()
	})

	expectedOutput := "El país España pertenece al continente Europa\n"
	if output != expectedOutput {
		t.Errorf("PlacePais test failed. Expected output: %q, got: %q", expectedOutput, output)
	}
}

func captureOutput(fn func()) string {
	output := ""
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	output = string(out)

	return output
}
