package client

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type DatLoader struct {
	things []DatThing
}

type DatThing struct {
	Width  uint8
	Height uint8

	Frames uint8
	XDiv   uint8
	YDiv   uint8
	ZDiv   uint8

	AnimationLength uint8

	Sprites []uint32
}

func (loader *DatLoader) Load(filename string) error {
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

	var items, creatures, magicEffects, distanceEffects uint16
	if err := binary.Read(reader, binary.LittleEndian, &items); err != nil {
		return err
	} else if err := binary.Read(reader, binary.LittleEndian, &creatures); err != nil {
		return err
	} else if err := binary.Read(reader, binary.LittleEndian, &magicEffects); err != nil {
		return err
	} else if err := binary.Read(reader, binary.LittleEndian, &distanceEffects); err != nil {
		return err
	}

	loader.things = make([]DatThing, items+1)

	clientID := uint16(100)
	for clientID <= items {
		if err := loader.readAttributes(reader); err != nil {
			return err
		}

		var width, height uint8
		if width, err = reader.ReadByte(); err != nil {
			return err
		} else if height, err = reader.ReadByte(); err != nil {
			return err
		}

		if width > 1 || height > 1 {
			if _, err := reader.ReadByte(); err != nil {
				return err
			}
		}

		var frames, xdiv, ydiv, zdiv, animationLength uint8
		if frames, err = reader.ReadByte(); err != nil {
			return err
		} else if xdiv, err = reader.ReadByte(); err != nil {
			return err
		} else if ydiv, err = reader.ReadByte(); err != nil {
			return err
		} else if zdiv, err = reader.ReadByte(); err != nil {
			return err
		} else if animationLength, err = reader.ReadByte(); err != nil {
			return err
		}

		spriteCount := uint16(width) * uint16(height) *
			uint16(xdiv) * uint16(ydiv) * uint16(zdiv) *
			uint16(frames) * uint16(animationLength)

		dt := DatThing{width, height, frames, xdiv, ydiv, zdiv, animationLength,
			make([]uint32, spriteCount)}

		for i := uint16(0); i < spriteCount; i++ {
			var spriteID uint32
			if err := binary.Read(reader, binary.LittleEndian, &spriteID); err != nil {
				return err
			}

			dt.Sprites[i] = spriteID
			//loader.sprit[clientID] = append(loader.spriteIDs[clientID], spriteID)
		}

		loader.things[clientID] = dt
		clientID++
	}
	return nil
}

func (loader *DatLoader) readAttributes(reader *bufio.Reader) (err error) {
	for {
		var attr byte
		if attr, err = reader.ReadByte(); err != nil {
			return err
		}

		switch attr {
		case DatAttributeGround:
			var groundSpeed uint16
			if err = binary.Read(reader, binary.LittleEndian, &groundSpeed); err != nil {
				return err
			}
		case DatAttributeFirstTopOrder:
		case DatAttributeSecondTopOrder:
		case DatAttributeThirdTopOrder:
		case DatAttributeContainer:
		case DatAttributeStackable:
		case DatAttributeUnknown1:
		case DatAttributeUseTarget:
		case DatAttributeWritable:
			var maxLength uint16
			if err = binary.Read(reader, binary.LittleEndian, &maxLength); err != nil {
				return err
			}
		case DatAttributeReadable:
			var maxLength uint16
			if err = binary.Read(reader, binary.LittleEndian, &maxLength); err != nil {
				return err
			}
		case DatAttributeFluidContainer:
		case DatAttributeSplash:
		case DatAttributeBlockSolid:
		case DatAttributeImmovable:
		case DatAttributeBlockProjectile:
		case DatAttributeBlockPathfind:
		case DatAttributeNoMoveAnimation:
		case DatAttributePickupable:
		case DatAttributeHangable:
		case DatAttributeHorizontalHangable:
		case DatAttributeVerticalHangable:
		case DatAttributeRotatable:
		case DatAttributeLight:
			var level, color uint16
			if err = binary.Read(reader, binary.LittleEndian, &level); err != nil {
				return err
			} else if err = binary.Read(reader, binary.LittleEndian, &color); err != nil {
				return err
			}
		case DatAttributeUnknown2:
		case DatAttributeFloorChange:
		case DatAttributeUnknown3:
			var unknown uint32
			if err = binary.Read(reader, binary.LittleEndian, &unknown); err != nil {
				return err
			}
		case DatAttributeHeight:
			var height uint16
			if err = binary.Read(reader, binary.LittleEndian, &height); err != nil {
				return err
			}
		case DatAttributeUnknown4:
		case DatAttributeUnknown5:
		case DatAttributeMinimapColor:
			var color uint16
			if err = binary.Read(reader, binary.LittleEndian, &color); err != nil {
				return err
			}
		case DatAttributeUnknown6:
			var unknown uint16
			if err = binary.Read(reader, binary.LittleEndian, &unknown); err != nil {
				return err
			}
		case DatAttributeFullTile:
		case DatAttributeLookThrough:
		case DatAttributeUnknown7:
			var unknown uint16
			if err = binary.Read(reader, binary.LittleEndian, &unknown); err != nil {
				return err
			}
		case DatAttributeMarket:
			var category, tradeAs, showAs, strLength uint16
			if err = binary.Read(reader, binary.LittleEndian, &category); err != nil {
				return err
			} else if err = binary.Read(reader, binary.LittleEndian, &tradeAs); err != nil {
				return err
			} else if err = binary.Read(reader, binary.LittleEndian, &showAs); err != nil {
				return err
			} else if err = binary.Read(reader, binary.LittleEndian, &strLength); err != nil {
				return err
			}

			name := make([]byte, strLength)
			pos := uint16(0)
			for pos < strLength {
				var n int
				if n, err = reader.Read(name[pos:]); err != nil {
					return err
				}
				pos += uint16(n)
			}

			var profession, level uint16
			if err := binary.Read(reader, binary.LittleEndian, &profession); err != nil {
				return err
			} else if err := binary.Read(reader, binary.LittleEndian, &level); err != nil {
				return err
			}
		case DatAttributeDefaultAction:
			var unknown uint16
			if err := binary.Read(reader, binary.LittleEndian, &unknown); err != nil {
				return err
			}
		case DatAttributeUsable:
		case DatAttributeEnd:
			return nil
		default:
			return fmt.Errorf("unknown dat attribute %d", attr)
		}
	}
}

func (loader *DatLoader) GetThing(clientID uint16) DatThing {
	return loader.things[clientID]
}
