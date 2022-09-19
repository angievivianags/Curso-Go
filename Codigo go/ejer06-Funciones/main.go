package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(1)

	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(tres(5))

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5, 15, 58, 23, 10))
	fmt.Println(calculo(5, 6, 20, 45))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 4, 6, 7, 9))

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func tres(numero int) (z int) {
	z = numero * 2
	return
}

func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
