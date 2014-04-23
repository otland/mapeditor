package client

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

func (loader *SpriteLoader) Load(filename string) error {
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

func (loader *SpriteLoader) GetSprite(id uint32) []byte {
	idx := loader.spriteIndex[id]
	if idx == 0 {
		return nil
	}

	length := uint32(binary.LittleEndian.Uint16(loader.data[idx:]))
	data := loader.data[idx+2 : idx+2+length]

	// NRGBA sprite
	sprite := make([]byte, 32*32*4)
	sp := 0

	for len(data) > 0 {
		transparentPixels := int(binary.LittleEndian.Uint16(data))
		coloredPixels := int(binary.LittleEndian.Uint16(data[2:]))

		data = data[4:]

		sp += transparentPixels * 4

		for i := 0; i < coloredPixels; i++ {
			sprite[sp] = data[0]
			sprite[sp+1] = data[1]
			sprite[sp+2] = data[2]
			sprite[sp+3] = 255

			data = data[3:]
			sp += 4
		}
	}

	return sprite
}
