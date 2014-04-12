package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"syscall"
)

type SpriteLoader struct {
	data        []byte
	spriteIndex []uint32
}

func (loader *SpriteLoader) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var signature uint32
	if err := binary.Read(reader, binary.LittleEndian, &signature); err != nil {
		return err
	}

	var num_sprites uint32
	if err := binary.Read(reader, binary.LittleEndian, &num_sprites); err != nil {
		return err
	}

	offset := (num_sprites * 4) - 4
	spriteIndexOffset := offset - 3

	loader.spriteIndex = make([]uint32, num_sprites+1)
	for i := uint32(1); i <= num_sprites; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &loader.spriteIndex[i]); err != nil {
			return err
		}
		loader.spriteIndex[i] -= spriteIndexOffset
	}

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	if loader.data, err = syscall.Mmap(int(file.Fd()), int64(offset), int(fi.Size()-int64(offset)), syscall.PROT_READ, syscall.MAP_PRIVATE); err != nil {
		return err
	}
	return nil
}

func (loader *SpriteLoader) getSprite(id uint32) []byte {
	idx := loader.spriteIndex[id]
	length := uint16(loader.data[idx]) | (uint16(loader.data[idx+1]) << 8)
	return loader.data[idx+2 : idx+2+uint32(length)]
}
