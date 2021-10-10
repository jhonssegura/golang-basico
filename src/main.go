package main

import "fmt"

func main() {

	// Declaraci{on de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Calcular el area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("EL area del cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Printf

	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf(message)

}
