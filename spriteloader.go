package main

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/otland/mmap-go"
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

	var signature, numSprites uint32
	if err := binary.Read(reader, binary.LittleEndian, &signature); err != nil {
		return err
	} else if err := binary.Read(reader, binary.LittleEndian, &numSprites); err != nil {
		return err
	}

	offset := (int64(numSprites) * 4) - 8
	spriteIndexOffset := uint32(offset - 3)

	loader.spriteIndex = make([]uint32, numSprites+1)
	for i := uint32(1); i <= numSprites; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &loader.spriteIndex[i]); err != nil {
			return err
		} else if loader.spriteIndex[i] != 0 {
			loader.spriteIndex[i] -= spriteIndexOffset
		}
	}

	if fi, err := file.Stat(); err != nil {
		return err
	} else if loader.data, err = mmap.MapRegion(file, int(fi.Size()-offset), mmap.RDONLY, 0, offset); err != nil {
		return err
	}
	return nil
}

func (loader *SpriteLoader) getSprite(id uint32) []byte {
	idx := loader.spriteIndex[id]
	if idx == 0 {
		return nil
	}

	length := uint32(loader.data[idx]) | uint32(loader.data[idx+1])<<8
	return loader.data[idx+2 : idx+2+length]
}
