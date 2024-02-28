package DiskManagement

import (
	"fmt"
   "strings"
   "strconv"
	"encoding/binary"
	"MIA_P1/Utilities"
	"MIA_P1/Structs"
)

func Mount(driveletter string, name string) {
	fmt.Println("======Start MOUNT======")
	fmt.Println("Driveletter:", driveletter)
	fmt.Println("Name:", name)

	// Open bin file
	filepath := "./test/" + strings.ToUpper(driveletter)  + ".bin"
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

	var index int = -1
	var count = 0
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			count++
			if strings.Contains(string(TempMBR.Partitions[i].Name[:]), name) {
				index = i
				break
			}
		}
	}

	if index != -1 {
		fmt.Println("Partition found")
		Structs.PrintPartition(TempMBR.Partitions[index])
	}else{
		fmt.Println("Partition not found")
		return
	}

	// id = DriveLetter + Correlative + 19

	id := strings.ToUpper(driveletter) + strconv.Itoa(count) + "19"

	copy(TempMBR.Partitions[index].Status[:], "1")
	copy(TempMBR.Partitions[index].Id[:], id)

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

	fmt.Println("======End MOUNT======")
}

func Fdisk(size int, driveletter string, name string, unit string, type_ string, fit string) {
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
	filepath := "./test/" + strings.ToUpper(driveletter)  + ".bin"
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

func Mkdisk(size int, fit string, unit string) {
	fmt.Println("======Start MKDISK======") 
	fmt.Println("Size:", size)
	fmt.Println("Fit:", fit)
	fmt.Println("Unit:", unit)

	// validate fit equals to b/w/f
	if fit != "bf" && fit != "wf" && fit != "ff" {
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

