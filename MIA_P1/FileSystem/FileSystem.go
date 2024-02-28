package FileSystem

import (
	 "os"
    "fmt"
    "strings"
    "encoding/binary"
    "MIA_P1/Utilities"
    "MIA_P1/Structs"
)

func Mkfs(id string, type_ string, fs_ string){
   fmt.Println("======Start MKFS======")
   fmt.Println("Id:", id)
   fmt.Println("Type:", type_)
   fmt.Println("Fs:", fs_)

   driveletter := string(id[0])

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
   // Iterate over the partitions
   for i := 0; i < 4; i++ {
      if TempMBR.Partitions[i].Size != 0 {
         if strings.Contains(string(TempMBR.Partitions[i].Id[:]), id) {
            fmt.Println("Partition found")
            if strings.Contains(string(TempMBR.Partitions[i].Status[:]), "1") {
               fmt.Println("Partition is mounted")
               index = i
            }else{
               fmt.Println("Partition is not mounted")
               return
            }
            break
         }
      }
   }

   if index != -1 {
      Structs.PrintPartition(TempMBR.Partitions[index])
   }else{
      fmt.Println("Partition not found")
      return
   }

   // numerador = (partition_montada.size - sizeof(Structs::Superblock)
   // denrominador base = (4 + sizeof(Structs::Inodes) + 3 * sizeof(Structs::Fileblock))
   // temp = "2" ? 0 : sizeof(Structs::Journaling)
   // denrominador = base + temp
   // n = floor(numerador / denrominador)

   numerador := int32(TempMBR.Partitions[index].Size - int32(binary.Size(Structs.Superblock{})))
   denrominador_base := int32(4 + int32(binary.Size(Structs.Inode{})) + 3 * int32(binary.Size(Structs.Fileblock{})))
   var temp int32 = 0
   if fs_ == "2fs" {
      temp = 0
   }else{
      temp = int32(binary.Size(Structs.Journaling{}))
   }
   denrominador := denrominador_base + temp
   n := int32(numerador / denrominador)

   fmt.Println("N:", n)

   // var newMRB Structs.MRB
   var newSuperblock Structs.Superblock
   newSuperblock.S_inodes_count = 0
   newSuperblock.S_blocks_count = 0

   newSuperblock.S_free_blocks_count = 3 * n
   newSuperblock.S_free_inodes_count = n

   copy(newSuperblock.S_mtime[:], "28/02/2024")
   copy(newSuperblock.S_umtime[:], "28/02/2024")
   newSuperblock.S_mnt_count = 0

   if fs_ == "2fs" {
      create_ext2(n, TempMBR.Partitions[index] ,newSuperblock, "28/02/2024", file)
   }else{
      fmt.Println("EXT3")
   }

   // Close bin file
   defer file.Close()

   fmt.Println("======End MKFS======")
}


func create_ext2 (n int32, partition Structs.Partition ,newSuperblock Structs.Superblock, date string, file *os.File){
   fmt.Println("======Start CREATE EXT2======")
   fmt.Println("N:", n)
   fmt.Println("Superblock:", newSuperblock)
   fmt.Println("Date:", date)


   newSuperblock.S_filesystem_type = 2
   newSuperblock.S_bm_inode_start = partition.Start + int32(binary.Size(Structs.Superblock{}))
   newSuperblock.S_bm_block_start = newSuperblock.S_bm_inode_start + n 
   newSuperblock.S_inode_start = newSuperblock.S_bm_block_start + 3 * n
   newSuperblock.S_block_start = newSuperblock.S_inode_start + n * int32(binary.Size(Structs.Inode{}))

   newSuperblock.S_free_inodes_count -= 1
   newSuperblock.S_free_blocks_count -= 1
   newSuperblock.S_free_inodes_count -= 1
   newSuperblock.S_free_blocks_count -= 1
 
	for i  := int32(0); i < n; i++ {
		err := Utilities.WriteObject(file, byte(0), int64(newSuperblock.S_bm_inode_start + i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

   for i  := int32(0); i < 3 * n; i++ {
      err := Utilities.WriteObject(file, byte(0), int64(newSuperblock.S_bm_block_start + i))
      if err != nil {
         fmt.Println("Error: ", err)
      }
   }

   var newInode Structs.Inode
   for i  := int32(0); i < 15; i++ {
      newInode.I_block[i] = -1
   }

   for i  := int32(0); i < n; i++ {
      err := Utilities.WriteObject(file, newInode, int64(newSuperblock.S_inode_start + i * int32(binary.Size(Structs.Inode{}))))
      if err != nil {
         fmt.Println("Error: ", err)
      }
   }

   var newFileblock Structs.Fileblock
   for i  := int32(0); i < 3 * n; i++ {
      err := Utilities.WriteObject(file, newFileblock, int64(newSuperblock.S_block_start + i * int32(binary.Size(Structs.Fileblock{}))))
      if err != nil {
         fmt.Println("Error: ", err)
      }
   }

   var Inode0 Structs.Inode //Inode 0
   Inode0.I_uid = 1
   Inode0.I_gid = 1
   Inode0.I_size = 0
   copy(Inode0.I_atime[:], date)
   copy(Inode0.I_ctime[:], date)
   copy(Inode0.I_mtime[:], date)
   copy(Inode0.I_perm[:], "0")
   copy(Inode0.I_perm[:], "664")

   for i  := int32(0); i < 15; i++ {
      Inode0.I_block[i] = -1
   }

   Inode0.I_block[0] = 0

   // . | 0
   // .. | 0
   // users.txt | 1
   //

   var Folderblock0 Structs.Folderblock //Bloque 0 -> carpetas
   Folderblock0.B_content[0].B_inodo = 0
   copy(Folderblock0.B_content[0].B_name[:], ".")
   Folderblock0.B_content[1].B_inodo = 0
   copy(Folderblock0.B_content[1].B_name[:], "..")
   Folderblock0.B_content[1].B_inodo = 1
   copy(Folderblock0.B_content[1].B_name[:], "users.txt")

   var Inode1 Structs.Inode //Inode 1
   Inode1.I_uid = 1
   Inode1.I_gid = 1
   Inode1.I_size = int32(binary.Size(Structs.Folderblock{}))
   copy(Inode1.I_atime[:], date)
   copy(Inode1.I_ctime[:], date)
   copy(Inode1.I_mtime[:], date)
   copy(Inode1.I_perm[:], "0")
   copy(Inode1.I_perm[:], "664")

   for i  := int32(0); i < 15; i++ {
      Inode1.I_block[i] = -1
   }

   Inode0.I_block[0] = 1


   data := "1,G,root\n1,U,root,root,123\n"
   var Fileblock1 Structs.Fileblock  //Bloque 1 -> archivo
   copy(Fileblock1.B_content[:], data)

   // Inodo 0 -> Bloque 0 -> Inodo 1 -> Bloque 1
   // Crear la carpeta raiz /
   // Crear el archivo users.txt "1,G,root\n1,U,root,root,123\n"


   // write superblock
   err := Utilities.WriteObject(file, newSuperblock, int64(partition.Start))


   // write bitmap inodes
   err = Utilities.WriteObject(file, byte(1), int64(newSuperblock.S_bm_inode_start))
   err = Utilities.WriteObject(file, byte(1), int64(newSuperblock.S_bm_inode_start + 1))

   // write bitmap blocks
   err = Utilities.WriteObject(file, byte(1), int64(newSuperblock.S_bm_block_start))
   err = Utilities.WriteObject(file, byte(1), int64(newSuperblock.S_bm_block_start + 1))

   // write inodes
   err = Utilities.WriteObject(file, Inode0, int64(newSuperblock.S_inode_start)) //Inode 0
   err = Utilities.WriteObject(file, Inode1, int64(newSuperblock.S_inode_start + int32(binary.Size(Structs.Inode{}))) ) //Inode 1

   // write blocks
   err = Utilities.WriteObject(file, Folderblock0, int64(newSuperblock.S_block_start)) //Bloque 0
   err = Utilities.WriteObject(file, Fileblock1, int64(newSuperblock.S_block_start + int32(binary.Size(Structs.Fileblock{}))) ) //Bloque 1  

   if err != nil {
      fmt.Println("Error: ", err)
   }

	//mkfs -type=full -id=A119

   fmt.Println("======End CREATE EXT2======")
}