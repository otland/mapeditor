package ot

import "fmt"

type Item struct {
	clientId uint16
	serverId uint16
	count    uint16 // Or subtype

	attributes map[uint8]interface{}
	teleportDestination Position

	children []Item
}

func (item *Item) create(serverId uint16) {
	item.serverId = serverId
	item.attributes = make(map[uint8]interface{})
}

func (item *Item) unserialize(binaryNode *BinaryNode) (err error) {
	if item.serverId, err = binaryNode.getShort(); err != nil {
		return
	}

	item.attributes = make(map[uint8]interface{})
	for len(binaryNode.data) != 0 {
		var attribute uint8

		if attribute, err = binaryNode.getByte(); err != nil {
			return
		}

		if attribute == 0 {
			return nil
		}

		switch attribute {
		case OTBMItemAttrCount, OTBMItemAttrRuneCharges:
			var count uint8
			if count, err = binaryNode.getByte(); err != nil {
				return
			}

			item.count = uint16(count)

		case OTBMItemAttrCharges:
			if item.count, err = binaryNode.getShort(); err != nil {
				return
			}

		case OTBMItemAttrHouseDoorID, OTBMItemAttrDecayState:
			var b uint8
			if b, err = binaryNode.getByte(); err != nil {
				return
			}

			item.attributes[attribute] = int(b)

		case OTBMItemAttrActionID, OTBMItemAttrUniqueID, OTBMItemAttrDepotID:
			var s uint16
			if s, err = binaryNode.getShort(); err != nil {
				return
			}

			item.attributes[attribute] = int(s)

		case OTBMItemAttrContainerItems, OTBMItemAttrDuration, OTBMItemAttrWrittenDate,
				OTBMItemAttrSleepingGUID, OTBMItemAttrSleepStart:
			var u uint32
			if u, err = binaryNode.getLong(); err != nil {
				return
			}

			item.attributes[attribute] = int(u)

		case OTBMItemAttrTeleDest:
			if item.teleportDestination, err = binaryNode.getPosition(); err != nil {
				return
			}

		case OTBMItemAttrText, OTBMItemAttrDesc, OTBMItemAttrWrittenBy:
			var s string
			if s, err = binaryNode.getString(); err != nil {
				return
			}

			item.attributes[attribute] = s

		default:
			return fmt.Errorf("Unknown item attribute: %d for id: %d", attribute, item.serverId)
		}
	}

	return nil
}

func (item *Item) isContainer() bool {
	return item.attributes[OTBMItemAttrContainerItems] != 0
}
