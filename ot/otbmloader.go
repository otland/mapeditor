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

	if nodeType != OTBM_NodeMapData {
		return fmt.Errorf("Corrupt OTBM file (Expected OTBM_NodeMapData got: 0x%X)", nodeType)
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
		case OTBM_AttrDescription:
			otMap.description += tmp

		case OTBM_AttrSpawnFile:
			otMap.spawnFile = tmp

		case OTBM_AttrHouseFile:
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

		if nodeType == OTBM_NodeTileArea {
			var basePos Position
			if basePos, err = node.getPosition(); err != nil {
				return err
			}

			for t := range node.children {
				nodeTile := node.children[t]
				if nodeType, err = nodeTile.getByte(); err != nil {
					return
				}

				if nodeType != OTBM_NodeTile && nodeType != OTBM_NodeHouseTile {
					return fmt.Errorf("Corrupt OTBM File, Tile node should be either OTBM_Tile or OTBM_HouseTile (got: 0x%X)", nodeType)
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

				house := &House{}
				if nodeType == OTBM_NodeHouseTile {
					var id uint32
					if id, err = nodeTile.getLong(); err != nil {
						return err
					}

					if tmpHouse := &otMap.houses[id]; tmpHouse == nil {
						otMap.houses = append(otMap.houses, *tmpHouse)
					} else {
						house = tmpHouse
					}

					house.id = id
					tile.flags |= TileFlag_House
					house.tiles = append(house.tiles, tile)
				}

				for len(nodeTile.data) != 0 {
					var tileAttribute uint8
					if tileAttribute, err = nodeTile.getByte(); err != nil {
						return
					}

					switch tileAttribute {
					case OTBM_AttrTileFlags:
						if tile.flags, err = nodeTile.getLong(); err != nil {
							return
						}

					case OTBM_AttrItem:
						// This is the ground item, it's always the bottom-level item in the
						// tile.items array.  So to access it just use the 0 index.

						var tileItem Item
						if tileItem.serverId, err = nodeTile.getShort(); err != nil {
							return
						}

						tile.items = append(tile.items, tileItem)

					default:
						return fmt.Errorf("Unknown tile attribute: 0x%X", tileAttribute)
					}
				}

				for i := range nodeTile.children {
					nodeItem := nodeTile.children[i]
					if nodeType, err = nodeItem.getByte(); err != nil {
						return
					}

					if nodeType != OTBM_NodeItem {
						return fmt.Errorf("Corrupt OTBM file, expected OTBM_Item node in OTBM_Tile node! (got: 0x%X)", nodeType)
					}

					var item Item
					if err = item.unserialize(&nodeItem); err != nil {
						return
					}

					if item.isContainer() {
						fmt.Println("CONTAINER")
						for c := range nodeItem.children {
							nodeContainerItem := nodeItem.children[c]
							if nodeType, err = nodeContainerItem.getByte(); err != nil {
								return
							}

							if nodeType != OTBM_NodeItem {
								return fmt.Errorf("Corrupt OTBM file, expected OTBM_Item node as child of an item container (got: 0x%X)", nodeType)
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
		} else if nodeType == OTBM_NodeTowns {
			for t := range node.children {
				nodeTown := node.children[t]
				if nodeType, err = nodeTown.getByte(); err != nil {
					return
				}

				if nodeType != OTBM_NodeTown {
					return fmt.Errorf("Corrupt OTBM file; expected OTBM_Town node after OTBM_Towns (got: 0x%X)", nodeType)
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

				town.templePos = pos
				otMap.towns = append(otMap.towns, town)
			}
		} else if nodeType == OTBM_NodeWaypoints && headerVersion > 1 {
			for w := range node.children {
				nodeWaypoint := node.children[w]
				if nodeType, err = nodeWaypoint.getByte(); err != nil {
					return
				}

				if nodeType != OTBM_NodeWaypoint {
					return fmt.Errorf("Corrupt OTBM file; expected OTBM_Waypoint after OTBM_Waypoints (got: 0x%X)", nodeType)
				}

				var name string
				if name, err = nodeWaypoint.getString(); err != nil {
					return
				}

				var pos Position
				if pos, err = nodeWaypoint.getPosition(); err != nil {
					return
				}

				otMap.waypoints[pos] = name
			}
		} else {
			return fmt.Errorf("Unknown OTBM attribute 0x%X!", nodeType)
		}
	}

	return nil
}
