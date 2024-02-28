package FileSystem

import (
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
    

   numerator := TempMBR.Partitions[index].Size - int32(binary.Size(Structs.Superblock{}))
   denominator := 4 + int32(binary.Size(Structs.Inode{})) + 3 * int32(binary.Size(Structs.Fileblock{}))
   var temp int
   if strings.Contains(fs_, "2fs") {
      temp = 0
   } else {
      temp = int(binary.Size(Structs.Journaling{}))
   }

   fmt.Println("temp:", temp)
   fmt.Println("numerator:", numerator)
   fmt.Println("denominator:", denominator)


   fmt.Println("======End MKFS======")
}