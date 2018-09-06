package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Ejercicio #1
	fmt.Println("Hello, 世界 🍺️")

	fmt.Println("")

	// Ejercicio #2
	if len(os.Args) != 2 {
		fmt.Println("Error: Se requiere un solo argumento de tipo int")
		os.Exit(1)
	}

	p := os.Args[1]
	fmt.Println("El argumento es: ", p)
	vi, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Error: El argumento no es casteable a int")
		os.Exit(1)
	}

	add := 0
	arr := make([]int,0)
	for i := 1; i <= vi; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			arr = append(arr, i)
			add =+ i
		}
	}

	fmt.Println("Lista: ", arr)
	fmt.Println("Suma:  ", add)
}