package main

import "fmt"

func main() {
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	// For while
	counter := 0
	fmt.Println("For While")
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For Range
	fmt.Println("For Range")
	listPairNumbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listPairNumbers {
		fmt.Printf("Posicion %d numero par: %d \n", i, par)
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
