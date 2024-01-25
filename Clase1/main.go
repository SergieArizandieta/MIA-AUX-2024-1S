package main

import (
	"encoding/binary"
	"fmt"
	// "io"
	"os"
)

type Profesor struct {
	Name [25]byte
	Age int32
}


func main() {

	// Instance new Profesor
	var newProfesor Profesor

	// Set values to newProfesor
	newProfesor.Name = [25]byte{'J', 'o', 's', 'e'}
	newProfesor.Age = 20

	// var name string
	// var age int32
	// name = "Sergie"
	// age = 21
	// Set values to newProfesor
	// copy(newProfesor.Name[:], name)
	// newProfesor.Age = age

	// Print Profesor
	fmt.Println("Profesor")
	fmt.Println("Nombre: ", string(newProfesor.Name [:]))
	fmt.Println("Edad: ", newProfesor.Age)

	// Create bin file
	if _, err := os.Stat("register.bin"); os.IsNotExist(err) {
		arch, err := os.Create("register.bin")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer arch.Close()
	}

	//Open bin file in read/write mode
	file, err := os.OpenFile("register.bin", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Move cursor to start file
	file.Seek(0, 0)

	// Explication Seek(0,0)
	// Seek moves the current offset in the file to a new position.
	// (0,0) 
	// first 0 is the offset
	// second 0 is the position, 0 = start file, 1 = current position, 2 = end file


	// Write Profesor in bin file
	binary.Write(file, binary.LittleEndian, &newProfesor)

	// Print success message
	fmt.Println("Se ha escrito el Profesor en el archivo")

	// Move cursor to start file
	file.Seek(0, 0)

	// Read Profesor from bin file
	var newProfesor2 Profesor
	err = binary.Read(file, binary.LittleEndian, &newProfesor2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Profesor
	fmt.Println("Profesor 2")
	fmt.Println("Nombre: ", string(newProfesor2.Name [:]))
	fmt.Println("Edad: ", newProfesor2.Age)

	// Close bin file
	defer file.Close()
}
