package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type OtbLoader struct {
	items        []ItemType
	majorVersion uint32
	minorVersion uint32
}

func (otbLoader *OtbLoader) load(fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var signature uint32
	if err = binary.Read(reader, binary.LittleEndian, &signature); err != nil {
		return
	}

	if signature != 0 {
		fmt.Println("Invalid OTB signature")
		return
	}

	var nodeStart uint8
	if err = binary.Read(reader, binary.LittleEndian, &nodeStart); err != nil || nodeStart != NODE_START {
		return fmt.Errorf("Failed to read node start")
	}

	root := &BinaryNode{}
	if err = root.unserialize(reader); err != nil {
		return
	}

	root.pos += 1 // first byte always 0
	if signature, err = root.getLong(); err != nil || signature != 0 {
		log.Print(err)
		return fmt.Errorf("Invalid signature in OTB file!")
	}

	if attr, err := root.getByte(); err != nil || attr != 0x01 {
		return fmt.Errorf("Invalid otb attr root version")
	}

	if size, err := root.getShort(); err != nil || size != 4+4+4+128 {
		return fmt.Errorf("Invalid otb attr root version size")
	}

	if otbLoader.majorVersion, err = root.getLong(); err != nil {
		return
	}

	if otbLoader.minorVersion, err = root.getLong(); err != nil {
		return
	}

	root.pos += 4 + 128

	lastId := 99
	for _, binaryNode := range root.children {
		var itemType ItemType
		itemType.unserialize(&binaryNode, otbLoader, lastId)
		otbLoader.items = append(otbLoader.items, itemType)
	}

	return
}
