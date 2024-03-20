package FileManager

import (
	"MIA_P1/Global"
	"MIA_P1/Structs"
	"MIA_P1/Utilities"
	"MIA_P1/UtilitiesInodes"

	"encoding/binary"
	"fmt"

	// "os"
	"strings"
)

// login -user=root -pass=123 -id=A119
// mkusr -user=user1 -pass=usuario -grp=usuarios
func Mkusr(user string, pass string, grp string) {
	fmt.Println("======Start MKUSR======")
	fmt.Println("User:", user)
	fmt.Println("Pass:", pass)
	fmt.Println("Grp:", grp)

	if !Global.Usuario.Status {
		fmt.Println("User already logged in")
		return
	}

	driveletter := string(Global.Usuario.ID[0])

	// Open bin file
	filepath := "./test/" + strings.ToUpper(driveletter) + ".bin"
	fmt.Println("Filepath:", filepath)
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}

	var TempMBR Structs.MRB
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	index := int(Global.Usuario.ID[1]) // read corelative

	fmt.Println("ID:", string(Global.Usuario.ID[:]))
	fmt.Println("index:", index)

	var tempSuperblock Structs.Superblock
	// Read object from bin file
	if err := Utilities.ReadObject(file, &tempSuperblock, int64(TempMBR.Partitions[index].Start)); err != nil {
		return
	}

	// initSearch /users.txt -> regresa no Inodo
	// initSearch -> 1
	indexInode := UtilitiesInodes.InitSearch("/users.txt", file, tempSuperblock)

	var crrInode Structs.Inode
	// Read object from bin file
	if err := Utilities.ReadObject(file, &crrInode, int64(tempSuperblock.S_inode_start+indexInode*int32(binary.Size(Structs.Inode{})))); err != nil {
		return
	}

	// read file data
	data := UtilitiesInodes.GetInodeFileData(crrInode, file, tempSuperblock)

	// UID , Tipo , Grupo , Nombre , Contraseña

	// read number of users

	// UID -> read number of users + 1

	// write new user -> validate if data > 64 -> create new block

	fmt.Println("indexInode:", indexInode)

	fmt.Println("======End MKUSR======")
}
