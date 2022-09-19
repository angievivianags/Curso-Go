package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo en Go")
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	//paises["Mexico"] = "D.F."

	fmt.Println(paises)
	fmt.Println(paises["Mexico"])

	//Crear Map
	paises5 := make(map[string]string, 5)
	paises5["Mexico"] = "D.F."

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)

	//Agregar una clave y su valor
	campeonato["River Plate"] = 25

	//Actualizar un valor
	campeonato["Chivas"] = 28
	fmt.Println(campeonato)

	//eliminar un elemento
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	//Recorrer el mapa

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un untaje de : %d \n", equipo, puntaje)
	}

	//Buscar
	puntaje, existe := campeonato["Mineiro"]

	fmt.Printf("El porcentaje capturado es %d, y el equiopo exise %t \n", puntaje, existe)
}
