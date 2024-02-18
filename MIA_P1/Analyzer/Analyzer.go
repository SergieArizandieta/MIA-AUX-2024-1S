package Analyzer

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	"encoding/binary"
	"MIA_P1/Utilities"
	"MIA_P1/Structs"
)

func AnalyzeType() {
	// define flags
	// size := flag.Int("size", 0, "Tamaño")
	// fit := flag.String("fit", "f", "Ajuste")
	// unit := flag.String("unit", "m", "Unidad")

	// define flags
	size := flag.Int("size", 0, "Tamaño")
	driveletter := flag.String("driveletter", "", "Letra")
	name := flag.String("name", "", "Nombre")
	unit := flag.String("unit", "m", "Unidad")
	type_ := flag.String("type", "p", "Tipo")
	fit2 := flag.String("fit", "f", "Ajuste")

	

	// Parse the command line into the defined flags. 
	flag.Parse()

	// Command line input "-size=3000 -unit=\"K a\""
	// input := "-size=3000 -unit=K -fit=\"BF\"" 
	
	input := "-size=300 -driveletter=A -name=Particion1" 
	processInputFdisk(input, size, driveletter, name, unit, type_, fit2)
	fdisk(*size, *driveletter, *name, *unit, *type_, *fit2)

	// Proccess the input string and set the values of the flags
	// processInputMkdisk(input, size, fit, unit)

	// Luego puedes llamar a tu función mkdisk y pasarle los valores necesarios
	// mkdisk(*size, *fit, *unit)

}

func fdisk(size int, driveletter string, name string, unit string, type_ string, fit string) {
	fmt.Println("======Start FDISK======")
	fmt.Println("Size:", size)
	fmt.Println("Driveletter:", driveletter)
	fmt.Println("Name:", name)
	fmt.Println("Unit:", unit)
	fmt.Println("Type:", type_)
	fmt.Println("Fit:", fit)

	// validate fit equals to b/w/f
	if fit != "b" && fit != "w" && fit != "f" {
		fmt.Println("Error: Fit must be b, w or f")
		return
	}

	// validate size > 0
	if size <= 0 {
		fmt.Println("Error: Size must be greater than 0")
		return
	}

	// validate unit equals to b/k/m
	if unit != "b" && unit != "k" && unit != "m" {
		fmt.Println("Error: Unit must be b, k or m")
		return
	}

	// validate type equals to p/e/l
	if type_ != "p" && type_ != "e" && type_ != "l" {
		fmt.Println("Error: Type must be p, e or l")
		return
	}

	// Set the size in bytes
	if unit == "k" {
		size = size * 1024
	}else{
		size = size * 1024 * 1024
	}

	// Open bin file
	filepath := "./test/" + driveletter + ".bin"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}

	var TempMBR Structs.MRB
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)

	fmt.Println("-------------")

	var count = 0
	var gap = int32(0)
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			count++
			gap = TempMBR.Partitions[i].Start + TempMBR.Partitions[i].Size
		}
	}

		
	

	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size == 0 {
			TempMBR.Partitions[i].Size = int32(size)

			if count == 0 {
				TempMBR.Partitions[i].Start = int32(binary.Size(TempMBR))
			}else{
				TempMBR.Partitions[i].Start = gap
			}
			
			copy(TempMBR.Partitions[i].Name[:], name)
			copy(TempMBR.Partitions[i].Fit[:], fit)
			copy(TempMBR.Partitions[i].Status[:], "0")
			copy(TempMBR.Partitions[i].Type[:], type_)		
			TempMBR.Partitions[i].Correlative = int32(count + 1)
			break
		}
	}

	// Overwrite the MBR
	if err := Utilities.WriteObject(file,TempMBR,0); err != nil {
		return
	}

	var TempMBR2 Structs.MRB
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR2, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR2)

	// Close bin file
	defer file.Close()

	fmt.Println("======End FDISK======")
}

func mkdisk(size int, fit string, unit string) {
	fmt.Println("======Start MKDISK======") 
	fmt.Println("Size:", size)
	fmt.Println("Fit:", fit)
	fmt.Println("Unit:", unit)

	// validate fit equals to b/w/f
	if fit != "b" && fit != "w" && fit != "f" {
		fmt.Println("Error: Fit must be b, w or f")
		return
	}

	// validate size > 0
	if size <= 0 {
		fmt.Println("Error: Size must be greater than 0")
		return
	}

	// validate unit equals to k/m
	if unit != "k" && unit != "m" {
		fmt.Println("Error: Unit must be k or m")
		return
	}

	// Create file
	err := Utilities.CreateFile("./test/A.bin")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Set the size in bytes
	if unit == "k" {
		size = size * 1024
	}else{
		size = size * 1024 * 1024
	}

	// Open bin file
	file, err := Utilities.OpenFile("./test/A.bin")
	if err != nil {
		return
	}

	// Write 0 binary data to the file

	// create array of byte(0)
	for i := 0; i < size; i++ {
		err := Utilities.WriteObject(file, byte(0), int64(i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}


	// Create a new instance of MRB
	var newMRB Structs.MRB
	newMRB.MbrSize = int32(size)
	newMRB.Signature = 10 // random
	copy(newMRB.Fit[:], fit)
	copy(newMRB.CreationDate[:], "2021-08-20")

	// Write object in bin file
	if err := Utilities.WriteObject(file,newMRB,0); err != nil {
		return
	}


	var TempMBR Structs.MRB
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)

	// Close bin file
	defer file.Close()

	fmt.Println("======End MKDISK======") 

}

func processInputMkdisk(input string, size *int, fit *string, unit *string) {
	re := regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		// Delete quotes if they are present in the value
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size":
			sizeValue := 0
			fmt.Sscanf(flagValue, "%d", &sizeValue)
			*size = sizeValue
		case "fit":
			flagValue = flagValue[:1]
			flagValue = strings.ToLower(flagValue)
			*fit = flagValue
		case "unit":
			flagValue = strings.ToLower(flagValue)
			*unit = flagValue
		default:
			fmt.Println("Error: Flag not found")
		}
	}
}

func processInputFdisk(input string, size *int, driveletter *string, name *string, unit *string, type_ *string, fit *string) {
	re := regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		// Delete quotes if they are present in the value
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size":
			sizeValue := 0
			fmt.Sscanf(flagValue, "%d", &sizeValue)
			*size = sizeValue
		case "fit":
			flagValue = flagValue[:1]
			flagValue = strings.ToLower(flagValue)
			*fit = flagValue
		case "unit":
			flagValue = strings.ToLower(flagValue)
			*unit = flagValue

		case "driveletter":
			flagValue = strings.ToUpper(flagValue)
			*driveletter = flagValue
		
		case "name":
			flagValue = strings.ToLower(flagValue)
			*name = flagValue
		
		case "type":
			flagValue = strings.ToLower(flagValue)
			*type_ = flagValue

		default:
			fmt.Println("Error: Flag not found")
		}
	}
}