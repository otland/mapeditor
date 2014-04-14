package ot

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type OtbLoader struct {
	items        []ItemType
	majorVersion uint32
	minorVersion uint32
}

func (otbLoader *OtbLoader) Load(fileName string) (err error) {
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

	var root BinaryNode
	if err = root.parse(reader); err != nil {
		return
	}

	root.skip(1) // first byte always 0
	if signature, err = root.getLong(); err != nil || signature != 0 {
		return errors.New("Invalid signature in OTB file!")
	}

	if attr, err := root.getByte(); err != nil || attr != 0x01 {
		return errors.New("Invalid otb attr root version")
	}

	if size, err := root.getShort(); err != nil || size != 4+4+4+128 {
		return errors.New("Invalid otb attr root version size")
	}

	if otbLoader.majorVersion, err = root.getLong(); err != nil {
		return
	}

	if otbLoader.minorVersion, err = root.getLong(); err != nil {
		return
	}

	root.skip(4 + 128)
	for i := range root.children {
		var itemType ItemType
		itemType.unserialize(&root.children[i])
		otbLoader.items = append(otbLoader.items, itemType)
	}

	return
}
