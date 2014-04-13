package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

const (
	ESCAPE_CHAR = 0xFD
	NODE_START  = 0xFE
	NODE_END    = 0xFF
)

type BinaryNode struct {
	pos      int
	data     []byte
	children []BinaryNode
}

func (binaryNode *BinaryNode) unpackShort() uint16 {
	return uint16(binaryNode.data[binaryNode.pos]) | uint16(binaryNode.data[binaryNode.pos+1])<<8
}

func (binaryNode *BinaryNode) getByte() (uint8, error) {
	if binaryNode.pos+1 > len(binaryNode.data) {
		return 0, fmt.Errorf("Out of data (pos: %d, size: %d)", binaryNode.pos, len(binaryNode.data))
	}

	b := binaryNode.data[binaryNode.pos]
	binaryNode.pos += 1
	return b, nil
}

func (binaryNode *BinaryNode) getShort() (uint16, error) {
	if binaryNode.pos+2 > len(binaryNode.data) {
		return 0, fmt.Errorf("Out of data (pos: %d, size: %d)", binaryNode.pos, len(binaryNode.data))
	}

	ret := binaryNode.unpackShort()
	binaryNode.pos += 2

	return ret, nil
}

func (binaryNode *BinaryNode) getLong() (uint32, error) {
	if binaryNode.pos+4 > len(binaryNode.data) {
		return 0, fmt.Errorf("Out of data (pos: %d, size: %d)", binaryNode.pos, len(binaryNode.data))
	}

	u16 := binaryNode.unpackShort()
	binaryNode.pos += 2
	ret := uint32(u16) | uint32(binaryNode.unpackShort())<<16
	binaryNode.pos += 2

	return ret, nil
}

func (binaryNode *BinaryNode) getString() (string, error) {
	var length uint16
	if length, err := binaryNode.getShort(); err != nil || length == 0 || length == 0xFFFF {
		return "", err
	}

	if binaryNode.pos+int(length) > len(binaryNode.data) {
		return "", fmt.Errorf("Out of data")
	}

	ret := string(binaryNode.data[binaryNode.pos : binaryNode.pos+int(length)])
	binaryNode.pos += int(length)
	return ret, nil
}

func (binaryNode *BinaryNode) unserialize(reader *bufio.Reader) error {
	for {
		var what uint8

		if err := binary.Read(reader, binary.LittleEndian, &what); err != nil {
			// Most likely an EOF and we don't care.
			return nil
		}

		switch what {
		case NODE_START:
			var newNode BinaryNode
			if err := newNode.unserialize(reader); err != nil {
				return err
			}

			binaryNode.children = append(binaryNode.children, newNode)
		case NODE_END:
			return nil
		case ESCAPE_CHAR:
			var b uint8
			if err := binary.Read(reader, binary.LittleEndian, &b); err != nil {
				return err
			}

			binaryNode.data = append(binaryNode.data, b)
		default:
			binaryNode.data = append(binaryNode.data, what)
		}
	}

	return nil
}
