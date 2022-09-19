package main

import "fmt"

func main() {

	//Primer For
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//For con  i en la misma instruccion
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//For sin la variable  i
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var x = 0

	for x < 10 {
		fmt.Printf("\n Valor de i: %d", x)
		if x == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			x = x * 2
			continue
		}
		fmt.Printf(" Paso por aqui \n")
		x++
	}

	var y int = 0
RUTINA:
	for y < 10 {
		if y == 4 {
			y = y + 2
			fmt.Println("Voy a RUTINA sumando 2 a y")
			goto RUTINA
		}
		fmt.Printf("Valor de y: %d\n", y)
		y++
	}
}
