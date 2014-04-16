package ot

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

func (otMap *Map) ReadOTBM(fileName string, otbLoader *OtbLoader) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var identifier [4]byte
	if err = binary.Read(reader, binary.LittleEndian, &identifier); err != nil {
		return
	}

	if !bytes.Equal(identifier[:4], []byte{'\x00', '\x00', '\x00', '\x00'}) && !bytes.Equal(identifier[:4], []byte{'O', 'T', 'B', 'M'}) {
		return fmt.Errorf("Corrupt OTBM file; OTBM Identifier must either be 0's or \"OTBM\" (got: %q)", identifier)
	}

	var root BinaryNode
	if err = root.parse(reader); err != nil {
		return
	}

	var property uint8
	if property, err = root.getByte(); err != nil {
		return
	}

	if property != 0 {
		return errors.New("Unable to read OTBM Root property!")
	}

	var headerVersion uint32
	if headerVersion, err = root.getLong(); err != nil {
		return
	}

	if headerVersion > 3 {
		return errors.New("Unknown OTBM Version detected")
	}

	if otMap.width, err = root.getShort(); err != nil {
		return
	}

	if otMap.height, err = root.getShort(); err != nil {
		return
	}

	var majorItemsVersion uint32
	if majorItemsVersion, err = root.getLong(); err != nil {
		return
	}

	if majorItemsVersion > otbLoader.majorVersion {
		return fmt.Errorf("This map was saved with a different OTB version (OTB version: %d, got: %d)", otbLoader.majorVersion, majorItemsVersion)
	}

	var minorItemsVersion uint32
	if minorItemsVersion, err = root.getLong(); err != nil {
		return
	}

	if minorItemsVersion > otbLoader.minorVersion {
		return fmt.Errorf("This map needs an updated OTB (OTB version: %d, got: %d)", otbLoader.minorVersion, minorItemsVersion)
	}

	var nodeType uint8

	mapDataNode := root.children[0]
	if nodeType, err = mapDataNode.getByte(); err != nil {
		return
	}

	if nodeType != OTBMNodeMapData {
		return fmt.Errorf("Corrupt OTBM file (Expected OTBMNodeMapData got: 0x%X)", nodeType)
	}

	for len(mapDataNode.data) != 0 {
		var attribute uint8
		var tmp string

		if attribute, err = mapDataNode.getByte(); err != nil {
			return
		}

		if tmp, err = mapDataNode.getString(); err != nil {
			return
		}

		switch attribute {
		case OTBMAttrDescription:
			otMap.description += tmp

		case OTBMAttrSpawnFile:
			otMap.spawnFile = tmp

		case OTBMAttrHouseFile:
			otMap.houseFile = tmp

		default:
			return fmt.Errorf("Corrupt OTBM file (got unknown attribute: 0x%X", attribute)
		}
	}

	for k := range mapDataNode.children {
		node := mapDataNode.children[k]
		if nodeType, err = node.getByte(); err != nil {
			return
		}

		if nodeType == OTBMNodeTileArea {
			var basePos Position
			if basePos, err = node.getPosition(); err != nil {
				return err
			}

			//fmt.Printf("read basePos: %d %d %d\n", basePos.x, basePos.y, basePos.z)
			for t := range node.children {
				nodeTile := node.children[t]
				if nodeType, err = nodeTile.getByte(); err != nil {
					return
				}

				if nodeType != OTBMNodeTile && nodeType != OTBMNodeHouseTile {
					return fmt.Errorf("Corrupt OTBM File, Tile node should be either OTBMTile or OTBMHouseTile (got: 0x%X)", nodeType)
				}

				var tile Tile

				var x, y uint8
				if x, err = nodeTile.getByte(); err != nil {
					return
				}
				tile.pos.x = uint16(x) + basePos.x

				if y, err = nodeTile.getByte(); err != nil {
					return
				}
				tile.pos.y = uint16(y) + basePos.y
				tile.pos.z = basePos.z

				//fmt.Printf("read tilePos: %d %d %d\n", tile.pos.x, tile.pos.y, tile.pos.z)
				house := &House{}
				if nodeType == OTBMNodeHouseTile {
					var id uint32
					if id, err = nodeTile.getLong(); err != nil {
						return err
					}

					fmt.Printf("read house id: %d\n", id)
					if tmpHouse := &otMap.houses[id]; tmpHouse == nil {
						otMap.houses = append(otMap.houses, *tmpHouse)
					} else {
						house = tmpHouse
					}

					house.id = id
					tile.flags |= TileFlagHouse
					house.tiles = append(house.tiles, tile)
				}

				for len(nodeTile.data) != 0 {
					var tileAttribute uint8
					if tileAttribute, err = nodeTile.getByte(); err != nil {
						return
					}

					switch tileAttribute {
					case OTBMAttrTileFlags:
						if tile.flags, err = nodeTile.getLong(); err != nil {
							return
						}

						//fmt.Printf("read tileflags: %d\n", tile.flags)
					case OTBMAttrItem:
						// This is the ground item, it's always the bottom-level item in the
						// tile.items array.  So to access it just use the 0 index.

						var tileItem Item
						if tileItem.serverId, err = nodeTile.getShort(); err != nil {
							return
						}

						tile.items = append(tile.items, tileItem)
						//fmt.Printf("read ground item: %d\n", tileItem.serverId)

					default:
						return fmt.Errorf("Unknown tile attribute: 0x%X", tileAttribute)
					}
				}

				for i := range nodeTile.children {
					nodeItem := nodeTile.children[i]
					if nodeType, err = nodeItem.getByte(); err != nil {
						return
					}

					if nodeType != OTBMNodeItem {
						return fmt.Errorf("Corrupt OTBM file, expected OTBMItem node in OTBMTile node! (got: 0x%X)", nodeType)
					}

					var item Item
					if err = item.unserialize(&nodeItem); err != nil {
						return
					}

					//fmt.Printf("finished reading item properities\n")
					if item.isContainer() {
						fmt.Println("CONTAINER")
						for c := range nodeItem.children {
							nodeContainerItem := nodeItem.children[c]
							if nodeType, err = nodeContainerItem.getByte(); err != nil {
								return
							}

							if nodeType != OTBMNodeItem {
								return fmt.Errorf("Corrupt OTBM file, expected OTBMItem node as child of a container (got: 0x%X)", nodeType)
							}

							var containerItem Item
							if err = containerItem.unserialize(&nodeContainerItem); err != nil {
								return
							}

							item.children = append(item.children, containerItem)
							/*
								if house != nil && newItem.isMovable {
									fmt.Printf("Warning: Movable item found in house (x: %d, y: %d, z: %d)", int(tilePos.x), int(tilePos.y), int(tilePos.z))
								}
							*/
						}
					}
				}
			}
		} else if nodeType == OTBMNodeTowns {
			fmt.Printf("read towns")
			for t := range node.children {
				nodeTown := node.children[t]
				if nodeType, err = nodeTown.getByte(); err != nil {
					return
				}

				if nodeType != OTBMNodeTown {
					return fmt.Errorf("Corrupt OTBM file; expected OTBMTown node after OTBMTowns (got: 0x%X)", nodeType)
				}

				var town Town
				if town.id, err = nodeTown.getLong(); err != nil {
					return
				}

				if town.name, err = nodeTown.getString(); err != nil {
					return
				}

				var pos Position
				if pos, err = nodeTown.getPosition(); err != nil {
					return
				}

				fmt.Printf("read town: %d %s %d %d %d\n", town.id, town.name, pos.x, pos.y, pos.z)
				town.templePos = pos
				otMap.towns = append(otMap.towns, town)
			}
		} else if nodeType == OTBMNodeWaypoints && headerVersion > 1 {
			fmt.Printf("start read waypoints\n")
			for w := range node.children {
				nodeWaypoint := node.children[w]
				if nodeType, err = nodeWaypoint.getByte(); err != nil {
					return
				}

				if nodeType != OTBMNodeWaypoint {
					return fmt.Errorf("Corrupt OTBM file; expected OTBMWaypoint after OTBMWaypoints (got: 0x%X)", nodeType)
				}

				var name string
				if name, err = nodeWaypoint.getString(); err != nil {
					return
				}

				var pos Position
				if pos, err = nodeWaypoint.getPosition(); err != nil {
					return
				}

				fmt.Printf("read waypoint: %s %d %d %d\n", name, pos.x, pos.y, pos.z)
				otMap.waypoints[pos] = name
			}
		} else {
			return fmt.Errorf("Unknown OTBM attribute 0x%X!", nodeType)
		}
	}

	return nil
}
