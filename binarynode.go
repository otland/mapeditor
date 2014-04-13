package main

import (
	"bufio"
	"errors"
	"fmt"
)

const (
	ESCAPE_CHAR = 0xFD
	NODE_START  = 0xFE
	NODE_END    = 0xFF
)

type BinaryNode struct {
	data     []byte
	children []BinaryNode
}

func (binaryNode *BinaryNode) unpackShort() uint16 {
	return uint16(binaryNode.data[0]) | uint16(binaryNode.data[1])<<8
}

func (binaryNode *BinaryNode) getByte() (uint8, error) {
	if len(binaryNode.data) < 1 {
		return 0, errors.New("Out of data!")
	}

	b := binaryNode.data[0]
	binaryNode.data = binaryNode.data[1:]
	return b, nil
}

func (binaryNode *BinaryNode) getShort() (uint16, error) {
	if len(binaryNode.data) < 2 {
		return 0, errors.New("Out of data!")
	}

	ret := binaryNode.unpackShort()
	binaryNode.data = binaryNode.data[2:]

	return ret, nil
}

func (binaryNode *BinaryNode) getLong() (uint32, error) {
	if len(binaryNode.data) < 4 {
		return 0, errors.New("Out of data!")
	}

	u16 := binaryNode.unpackShort()
	binaryNode.data = binaryNode.data[2:]
	ret := uint32(u16) | uint32(binaryNode.unpackShort())<<16
	binaryNode.data = binaryNode.data[2:]

	return ret, nil
}

func (binaryNode *BinaryNode) getString() (string, error) {
	var length uint16
	var err error

	if length, err = binaryNode.getShort(); err != nil {
		return "", err
	}

	if length == 0 || length == 0xFFFF {
		return "", errors.New("String length cannot be 0 or equal to 0xFFFF")
	}

	if len(binaryNode.data) < int(length) {
		return "", errors.New("Out of data")
	}

	ret := string(binaryNode.data[:int(length)])
	binaryNode.data = binaryNode.data[int(length):]
	return ret, nil
}

func (binaryNode *BinaryNode) skip(length int) error {
	if len(binaryNode.data) < length {
		return fmt.Errorf("Cannot skip %d bytes", length)
	}

	binaryNode.data = binaryNode.data[length:]
	return nil
}

func (binaryNode *BinaryNode) parse(reader *bufio.Reader) error {
	if startByte, err := reader.ReadByte(); err != nil {
		return err
	} else if startByte != NODE_START {
		return fmt.Errorf("Unable to read root node start byte (should be 0x%x got 0x%x\n", NODE_START, startByte)
	}

	return binaryNode.unserialize(reader)
}

func (binaryNode *BinaryNode) unserialize(reader *bufio.Reader) (err error) {
	for {
		var what uint8
		if what, err = reader.ReadByte(); err != nil {
			// Most likely an EOF and we don't care.
			return nil
		}

		switch what {
		case NODE_START:
			var newNode BinaryNode
			if err = newNode.unserialize(reader); err != nil {
				return
			}

			binaryNode.children = append(binaryNode.children, newNode)
		case NODE_END:
			return nil
		case ESCAPE_CHAR:
			var b uint8
			if b, err = reader.ReadByte(); err != nil {
				return
			}

			binaryNode.data = append(binaryNode.data, b)
		default:
			binaryNode.data = append(binaryNode.data, what)
		}
	}

	return nil
}
