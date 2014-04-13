package main

type ItemType struct {
	category uint8
	name     string
	clientId uint16
	serverId uint16
}

func (itemType ItemType) unserialize(binaryNode *BinaryNode, otbLoader *OtbLoader, lastId int) (err error) {
	if itemType.category, err = binaryNode.getByte(); err != nil {
		return
	}
	binaryNode.pos += 4 // flags

	for binaryNode.pos < len(binaryNode.data) {
		var attr uint8

		attr, err = binaryNode.getByte()
		if err != nil || attr == 0 || attr == 0xFF {
			return
		}

		var length uint16
		if length, err = binaryNode.getShort(); err != nil {
			return
		}

		switch attr {
		case ItemTypeAttrServerId:
			if itemType.serverId, err = binaryNode.getShort(); err != nil {
				return
			}

			// Fugly code, continue with caution.  You have been warned.
			if otbLoader.minorVersion < ClientVersion860 {
				if itemType.serverId > 20000 && itemType.serverId < 20100 {
					itemType.serverId -= 20000
				} else if uint16(lastId) > 99 && uint16(lastId) != itemType.serverId-1 {
					for uint16(lastId) != itemType.serverId-1 {
						var reservedType ItemType
						reservedType.serverId = uint16(lastId)
						lastId += 1
						otbLoader.items = append(otbLoader.items, reservedType)
					}
				}
			} else {
				if itemType.serverId > 30000 && itemType.serverId < 30100 {
					itemType.serverId -= 30000
				} else if uint16(lastId) > 99 && uint16(lastId) != itemType.serverId-1 {
					for uint16(lastId) != itemType.serverId-1 {
						var reservedType ItemType
						reservedType.serverId = uint16(lastId)
						lastId += 1
						otbLoader.items = append(otbLoader.items, reservedType)
					}
				}
			}
		case ItemTypeAttrClientId:
			if itemType.clientId, err = binaryNode.getShort(); err != nil {
				return
			}
		case ItemTypeAttrName:
			if itemType.name, err = binaryNode.getString(); err != nil {
				return
			}
		default:
			binaryNode.pos += int(length)
		}
	}

	return nil
}
