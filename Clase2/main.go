package main 

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

type Data struct {
	Name [25]byte
	ID int32
}

func PrintData(data Data){
	fmt.Println(fmt.Sprintf("Name: %s, id: %d", string(data.Name[:]), data.ID))
}

func main(){
	var fileName string
	fileName = "./test/test.bin"
	
	// Instance new Data
	var newData Data

	//Call function Print
	PrintData(newData)

	// Create bin file
	if err := CreateFile(fileName); err != nil {
		return
	}

	// Open bin file
	file, err := OpenFile(fileName)
	if err != nil {
		return
	}

	// Writing 10 objects to the file
	for i := 0; i < 10; i++ {
		// Set values to newData
		copy(newData.Name[:], "Sergie")
		newData.ID = int32(i+1)

		// Write object in bin file
		if err := WriteObject(file,newData,int64(i * binary.Size(newData))); err != nil {
			return
		}
	}

	// Read 10 objects from bin file
	for i := 0; i < 10; i++ {
		var TempData Data
		// Read object from bin file
		if err := ReadObject(file, &TempData, int64(i * binary.Size(TempData) + 0)); err != nil {
			return
		}

		// Print object
		PrintData(TempData)
	}
	

	// Close bin file
	defer file.Close()

	// Print end of program
	fmt.Println("End of program") 
	
}


// Funtion to create bin file
func CreateFile(name string) error {
	//Ensure the directory exists
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Err CreateFile dir==",err)
		return err
	}

	// Create file
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Err CreateFile create==",err)
			return err
		}
		defer file.Close()
	}
	return nil
}

// Funtion to open bin file in read/write mode
func OpenFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Err OpenFile==",err)
		return nil, err
	}
	return file, nil
}

// Function to Write an object in a bin file
func WriteObject(file *os.File, data interface{}, position  int64) error {
	file.Seek(position, 0)
	err := binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err WriteObject==",err)
		return err
	}
	return nil
}


// Function to Read an object from a bin file
func ReadObject(file *os.File, data interface{}, position  int64) error {
	file.Seek(position, 0)
	err := binary.Read(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err ReadObject==",err)
		return err
	}
	return nil
}
		