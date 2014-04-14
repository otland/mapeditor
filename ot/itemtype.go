package ot

type ItemType struct {
	category uint8
	name     string
	clientID uint16
	serverID uint16
}

func (itemType ItemType) unserialize(binaryNode *BinaryNode) (err error) {
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
		case ItemTypeAttrServerID:
			if itemType.serverID, err = binaryNode.getShort(); err != nil {
				return
			}

			/*
				if otbLoader.minorVersion < ClientVersion860 {
					if itemType.serverID > 20000 && itemType.serverID < 20100 {
						itemType.serverID -= 20000
					} else if *lastId > 99 && *lastId != itemType.serverID-1 {
						for *lastId != itemType.serverID-1 {
							var reservedType ItemType
							reservedType.serverID = *lastId
							*lastId += 1
							otbLoader.items = append(otbLoader.items, reservedType)
						}
					}
				} else {
					if itemType.serverID > 30000 && itemType.serverID < 30100 {
						itemType.serverID -= 30000
					} else if *lastID > 99 && *lastID != itemType.serverID-1 {
						for *lastID != itemType.serverID-1 {
							var reservedType ItemType
							reservedType.serverID = *lastID
							*lastID += 1
							otbLoader.items = append(otbLoader.items, reservedType)
						}
					}
				}
			*/
		case ItemTypeAttrClientID:
			if itemType.clientID, err = binaryNode.getShort(); err != nil {
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
