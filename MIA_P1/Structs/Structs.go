package Structs

import (
	"fmt"
)


type MRB struct {
	MbrSize int32
	CreationDate [10]byte
	Signature int32
	Fit [1]byte
	Partitions [4]Partition
}


func PrintMBR(data MRB){
	fmt.Println(fmt.Sprintf("CreationDate: %s, fit: %s, size: %d", string(data.CreationDate[:]), string(data.Fit[:]), data.MbrSize))
	for i := 0; i < 4; i++ {
		fmt.Println(fmt.Sprintf("Partition %d: %s, %s, %d, %d", i, string(data.Partitions[i].Name[:]), string(data.Partitions[i].Type[:]), data.Partitions[i].Start, data.Partitions[i].Size))
	}
}

type Partition struct {
	Status [1]byte
	Type [1]byte
	Fit [1]byte
	Start int32
	Size int32
	Name [16]byte
	Correlative int32
	Id [4]byte
}