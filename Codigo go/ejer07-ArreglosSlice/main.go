package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

func main() {

	//listas
	tabla[0] = 1
	tabla[5] = 15

	fmt.Println(tabla)

	otratabla := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(otratabla); i++ {
		fmt.Println(otratabla[i])
	}
	fmt.Println(otratabla)

	//Matriz

	matriz[3][5] = 1
	fmt.Println(matriz)

	matrizOtra := []int{2, 5, 9, 3, 4}
	fmt.Println(matrizOtra)

	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]

	porcion2 := elementos[2:4]

	fmt.Println(porcion)
	fmt.Println(porcion2)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

}

func variante4() {
	nums := make([]int, 0, 0)
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))

}
