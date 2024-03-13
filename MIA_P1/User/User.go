package User

import (
	//  "os"
	"MIA_P1/Structs"
	"MIA_P1/Utilities"
	"MIA_P1/UtilitiesInodes"
	"encoding/binary"
	"fmt"
	"strings"
)

// login -user=root -pass=123 -id=A119
func Login(user string, pass string, id string) {
	fmt.Println("======Start LOGIN======")
	fmt.Println("User:", user)
	fmt.Println("Pass:", pass)
	fmt.Println("Id:", id)

	driveletter := string(id[0])

	// Open bin file
	filepath := "./test/" + strings.ToUpper(driveletter) + ".bin"
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
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			if strings.Contains(string(TempMBR.Partitions[i].Id[:]), id) {
				fmt.Println("Partition found")
				if strings.Contains(string(TempMBR.Partitions[i].Status[:]), "1") {
					fmt.Println("Partition is mounted")
					index = i
				} else {
					fmt.Println("Partition is not mounted")
					return
				}
				break
			}
		}
	}

	if index != -1 {
		Structs.PrintPartition(TempMBR.Partitions[index])
	} else {
		fmt.Println("Partition not found")
		return
	}

	var tempSuperblock Structs.Superblock
	// Read object from bin file
	if err := Utilities.ReadObject(file, &tempSuperblock, int64(TempMBR.Partitions[index].Start)); err != nil {
		return
	}

	// initSearch /users.txt -> regresa no Inodo
	// initSearch -> 1
	indexInode := UtilitiesInodes.InitSearch("/users.txt", file, tempSuperblock)

	// indexInode := int32(1)

	var crrInode Structs.Inode
	// Read object from bin file
	if err := Utilities.ReadObject(file, &crrInode, int64(tempSuperblock.S_inode_start+indexInode*int32(binary.Size(Structs.Inode{})))); err != nil {
		return
	}

	// getInodeFileData -> Iterate the I_Block n concat the data

	var Fileblock Structs.Fileblock
	// Read object from bin file
	if err := Utilities.ReadObject(file, &Fileblock, int64(tempSuperblock.S_block_start+crrInode.I_block[0]*int32(binary.Size(Structs.Fileblock{})))); err != nil {
		return
	}

	fmt.Println("Fileblock------------")
	data := string(Fileblock.B_content[:])
	// Dividir la cadena en líneas
	lines := strings.Split(data, "\n")

	// Iterar a través de las líneas
	for _, line := range lines {
		// Imprimir cada línea
		fmt.Println(line)

	}

	// Print object
	fmt.Println("Inode", crrInode.I_block)

	// Close bin file
	defer file.Close()

	fmt.Println("======End LOGIN======")
}
