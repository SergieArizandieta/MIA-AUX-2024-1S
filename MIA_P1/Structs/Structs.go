package Structs

import (
	"fmt"
)

//  =============================================================

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
		PrintPartition(data.Partitions[i])
	}
}

//  =============================================================

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


func PrintPartition(data Partition){
	fmt.Println(fmt.Sprintf("Name: %s, type: %s, start: %d, size: %d, status: %s, id: %s", string(data.Name[:]), string(data.Type[:]), data.Start, data.Size, string(data.Status[:]), string(data.Id[:])))
}

//  =============================================================

type Superblock struct {
	S_filesystem_type int32
	S_inodes_count int32
	S_blocks_count int32
	S_free_blocks_count int32
	S_free_inodes_count int32
	S_mtime [17]byte
	S_umtime [17]byte
	S_mnt_count int32
	S_magic int32
	S_inode_size int32
	S_block_size int32
	S_fist_ino int32
	S_first_blo int32
	S_bm_inode_start int32
	S_bm_block_start int32
	S_inode_start int32
	S_block_start int32
}

//  =============================================================

type Inode struct {
	I_uid int32
	I_gid int32
	I_size int32
	I_atime [17]byte
	I_ctime [17]byte
	I_mtime [17]byte
	I_block [15]int32
	I_type [1]byte
	I_perm [3]byte
}

//  =============================================================

type Fileblock struct {
	B_content [64]byte
}

//  =============================================================

type Content struct {
	B_name [12]byte
	B_inodo int32
}

type Folderblock struct {
	B_content [4]Content
}

//  =============================================================

type Pointerblock struct {
	B_pointers [16]int32
}

//  =============================================================

type Content_J struct {
	Operation [10]byte
	Path [100]byte
	Content [100]byte
	Date [17]byte
}

type Journaling struct {
	Size int32
	Ultimo int32
	Contenido [50]Content_J
}
