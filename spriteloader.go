package main

import (
	"bufio"
	"encoding/binary"
	"os"
)

type SpriteLoader struct {
	file        *os.File
	spriteIndex []uint32
}

func (loader *SpriteLoader) load(filename string) error {
	var err error
	if loader.file, err = os.Open(filename); err != nil {
		return err
	}

	reader := bufio.NewReader(loader.file)

	var signature uint32
	if err := binary.Read(reader, binary.LittleEndian, &signature); err != nil {
		return err
	}

	var num_sprites uint32
	if err := binary.Read(reader, binary.LittleEndian, &num_sprites); err != nil {
		return err
	}

	loader.spriteIndex = make([]uint32, num_sprites+1)
	for i := uint32(1); i <= num_sprites; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &loader.spriteIndex[i]); err != nil {
			return err
		}
		loader.spriteIndex[i] += 3
	}
	return nil
}

func (loader *SpriteLoader) close() {
	loader.file.Close()
}

func (loader *SpriteLoader) getSprite(id uint32) ([]byte, error) {
	loader.file.Seek(int64(loader.spriteIndex[id]), os.SEEK_SET)
	reader := bufio.NewReader(loader.file)

	var length uint16
	if err := binary.Read(reader, binary.LittleEndian, &length); err != nil {
		return nil, err
	}

	data := make([]byte, length)

	var pos uint16
	for pos != length {
		if n, err := reader.Read(data[pos:]); err != nil {
			return nil, err
		} else {
			pos += uint16(n)
		}
	}
	return data, nil
}
