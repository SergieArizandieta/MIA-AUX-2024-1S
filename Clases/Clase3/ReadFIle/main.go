package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("===Start===")

	// Abrir el archivo
	file, err := os.Open("lec.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}

	// Crear un escáner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Recorrer el archivo línea por línea
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}

	defer file.Close()

}
