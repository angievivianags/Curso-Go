package main

import (
	"fmt"
	"strconv"
)

//Variables globales
var numero int
var texto string
var status bool

var statusTrue bool = true

//numeros con coma flotante
var numeroConComa1 float32
var numeroConComa2 float64

//enteros sin signo
var numeroSinSigno uint

func main() {
	//Variables local
	var numeroLocal int

	//variables sin decir var
	numero3 := 3

	numero4 := 4

	var moneda int64 = 0

	// convertir a otro tpo de datos
	numero3 = int(moneda)
	texto = fmt.Sprintf("%d", moneda)
	texto = strconv.Itoa(int(moneda))

	var numeroLocal3, numeroLocal4 int
	numeroLocal3, numeroLocal4, text := 3, 4, "Este es el text"

	fmt.Println(numero3, numero4, numeroLocal, numeroLocal3, numeroLocal4, text)
	fmt.Println(status)
}
