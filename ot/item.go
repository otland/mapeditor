package ot

import "fmt"

type Item struct {
	clientId uint16
	serverId uint16
	count    uint16 // Or subtype

	intAttributes map[uint8]int
	strAttributes map[uint8]string

	teleportDestination Position

	children []Item
}

func (item *Item) unserialize(binaryNode *BinaryNode) (err error) {
	if item.serverId, err = binaryNode.getShort(); err != nil {
		return
	}

	for len(binaryNode.data) != 0 {
		var attribute uint8

		if attribute, err = binaryNode.getByte(); err != nil {
			return
		}

		if attribute == 0 {
			return nil
		}

		switch attribute {
		case OTBM_ItemAttrCount:
		case OTBM_ItemAttrRuneCharges:
			var count uint8
			if count, err = binaryNode.getByte(); err != nil {
				return
			}

			item.count = uint16(count)

		case OTBM_ItemAttrCharges:
			if item.count, err = binaryNode.getShort(); err != nil {
				return
			}

		case OTBM_ItemAttrHouseDoorId:
		case OTBM_ItemAttrDecayState:
			var b uint8
			if b, err = binaryNode.getByte(); err != nil {
				return
			}

			item.intAttributes[attribute] = int(b)

		case OTBM_ItemAttrActionId:
		case OTBM_ItemAttrUniqueId:
		case OTBM_ItemAttrDepotId:
			var s uint16
			if s, err = binaryNode.getShort(); err != nil {
				return
			}

			fmt.Printf("is depot: %d got: %d\n", attribute == OTBM_ItemAttrDepotId, s)
			item.intAttributes[attribute] = int(s)

		case OTBM_ItemAttrContainerItems:
		case OTBM_ItemAttrDuration:
		case OTBM_ItemAttrWrittenDate:
		case OTBM_ItemAttrSleepingGUID:
		case OTBM_ItemAttrSleepStart:
			var u uint32
			if u, err = binaryNode.getLong(); err != nil {
				return
			}

			item.intAttributes[attribute] = int(u)

		case OTBM_ItemAttrTeleDest:
			if item.teleportDestination, err = binaryNode.getPosition(); err != nil {
				return
			}

		case OTBM_ItemAttrText:
		case OTBM_ItemAttrDesc:
		case OTBM_ItemAttrWrittenBy:
			var s string
			if s, err = binaryNode.getString(); err != nil {
				return
			}

			item.strAttributes[attribute] = s

		default:
			return fmt.Errorf("Unknown item attribute: 0x%X for id: %d", attribute, item.serverId)
		}
	}

	return nil
}

func (item *Item) isContainer() bool {
	return item.intAttributes[OTBM_ItemAttrContainerItems] != 0
}
