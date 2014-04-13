package main

type ItemType struct {
	category uint8
	name     string
	clientId uint16
	serverId uint16
}

func (itemType ItemType) unserialize(binaryNode *BinaryNode, otbLoader *OtbLoader, lastId uint16) (err error) {
	if itemType.category, err = binaryNode.getByte(); err != nil {
		return
	}
	binaryNode.skip(4) // flags

	for len(binaryNode.data) != 0 {
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
				} else if lastId > 99 && lastId != itemType.serverId-1 {
					for lastId != itemType.serverId-1 {
						var reservedType ItemType
						reservedType.serverId = lastId
						lastId += 1
						otbLoader.items = append(otbLoader.items, reservedType)
					}
				}
			} else {
				if itemType.serverId > 30000 && itemType.serverId < 30100 {
					itemType.serverId -= 30000
				} else if lastId > 99 && lastId != itemType.serverId-1 {
					for lastId != itemType.serverId-1 {
						var reservedType ItemType
						reservedType.serverId = lastId
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
			binaryNode.skip(int(length))
		}
	}

	return nil
}
