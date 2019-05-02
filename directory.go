package s7

import (
	"fmt"
)

// Copyright 2018 Trung Hieu Le. All rights reserved.
const (
	// Block type byte
	BlockOB  = 56
	BlockDB  = 65
	BlockSDB = 66
	BlockFC  = 67
	BlockSFC = 68
	BlockFB  = 69
	BlockSFB = 70
)

//S7BlocksList Block List
type S7BlocksList struct {
	OBList  []int
	FBList  []int
	FCList  []int
	SFBList []int
	SFCList []int
	DBList  []int
	SDBList []int
}

//implement list block
func (mb *client) PGListBlocks() (list S7BlocksList, err error) {
	list.OBList, err = mb.pgBlockList(BlockOB)
	//debug
	fmt.Printf("%v", list.DBList)
	list.DBList, err = mb.pgBlockList(BlockDB)
	list.FCList, err = mb.pgBlockList(BlockFC)
	list.OBList, err = mb.pgBlockList(BlockOB)
	list.FBList, err = mb.pgBlockList(BlockFB)
	list.SDBList, err = mb.pgBlockList(BlockSDB)
	list.SFBList, err = mb.pgBlockList(BlockSFB)
	list.SFCList, err = mb.pgBlockList(BlockSFC)
	return
}

func (mb *client) pgBlockList(blockType byte) (arr []int, err error) {
	bl := make([]byte, len(s7PGBlockListTelegram))
	copy(bl, s7PGBlockListTelegram)
	bl = append(bl, make([]byte, 1)...)
	switch blockType {
	case BlockDB:
		bl[len(bl)-1] = BlockDB
	case BlockOB:
		bl[len(bl)-1] = BlockOB
	case BlockSDB:
		bl[len(bl)-1] = BlockSDB
	case BlockFC:
		bl[len(bl)-1] = BlockFC
	case BlockSFC:
		bl[len(bl)-1] = BlockSFC
	case BlockFB:
		bl[len(bl)-1] = BlockFB
	case BlockSFB:
		bl[len(bl)-1] = BlockSFB
	default:
		return
	}
	request := NewProtocolDataUnit(bl)
	//send
	response, err := mb.send(&request)
	if err == nil {
		res := make([]byte, len(response.Data)-33) //remove first 26 byte function and 7 byte header
		copy(res, response.Data[33:len(response.Data)])
		arr = dataToBlocks(res)
	}
	return
}
func dataToBlocks(data []byte) []int {
	arr := make([]int, len(data)/4)
	for i := 0; i <= len(data)/4-1; i++ {
		arr[i] = int(data[i*4])*256 + int(data[i*4+1])
	}
	return arr
}
