package ot

const (
	TileFlag_None           = 0
	TileFlag_ProtectionZone = 1 << 0
	TileFlag_Trashed        = 1 << 1
	TileFlag_OptionalZone   = 1 << 2
	TileFlag_NoLogout       = 1 << 3
	TileFlag_HardcoreZone   = 1 << 4
	TileFlag_Refresh        = 1 << 5

	// internal usage
	TileFlag_House       = 1 << 6
	TileFlag_Teleport    = 1 << 17
	TileFlag_MagicField  = 1 << 18
	TileFlag_MailBox     = 1 << 19
	TileFlag_TrashHolder = 1 << 20
	TileFlag_Bed         = 1 << 21
	TileFlag_Depot       = 1 << 22

	TileFlag_Last = 1 << 24
)

const (
	OTBM_NodeRootV2    = 1
	OTBM_NodeMapData   = 2
	OTBM_NodeTileArea  = 4
	OTBM_NodeTile      = 5
	OTBM_NodeItem      = 6
	OTBM_NodeTowns     = 12
	OTBM_NodeTown      = 13
	OTBM_NodeHouseTile = 14
	OTBM_NodeWaypoints = 15
	OTBM_NodeWaypoint  = 16
)

const (
	OTBM_AttrDescription = 1
	OTBM_AttrTileFlags   = 3
	OTBM_AttrItem        = 9
	OTBM_AttrSpawnFile   = 11
	OTBM_AttrHouseFile   = 13
)

const (
	OTBM_ItemAttrActionId       = 4
	OTBM_ItemAttrUniqueId       = 5
	OTBM_ItemAttrText           = 6
	OTBM_ItemAttrDesc           = 7
	OTBM_ItemAttrTeleDest       = 8
	OTBM_ItemAttrItem           = 9
	OTBM_ItemAttrDepotId        = 10
	OTBM_ItemAttrRuneCharges    = 12
	OTBM_ItemAttrHouseDoorId    = 14
	OTBM_ItemAttrCount          = 15
	OTBM_ItemAttrDuration       = 16
	OTBM_ItemAttrDecayState     = 17
	OTBM_ItemAttrWrittenDate    = 18
	OTBM_ItemAttrWrittenBy      = 19
	OTBM_ItemAttrSleepingGUID   = 20
	OTBM_ItemAttrSleepStart     = 21
	OTBM_ItemAttrCharges        = 22
	OTBM_ItemAttrContainerItems = 23
)

const (
	ThingCategoryItem = iota
	ThingCategoryCreature
	ThingCategoryEffect
	ThingCategoryMissile
	ThingInvalidCategory
	ThingLastCategory
)

const (
	ItemCategoryInvalid = iota
	ItemCategoryGround
	ItemCategoryContainer
	ItemCategoryWeapon
	ItemCategoryAmmunition
	ItemCategoryArmor
	ItemCategoryCharges
	ItemCategoryTeleport
	ItemCategoryMagicField
	ItemCategoryWritable
	ItemCategoryKey
	ItemCategorySplash
	ItemCategoryFluid
	ItemCategoryDoor
	ItemCategoryDeprecated
	ItemCategoryLast
)

const (
	ItemTypeAttrFirst        = 15
	ItemTypeAttrServerId     = 16
	ItemTypeAttrClientId     = 17
	ItemTypeAttrName         = 18 // deprecated
	ItemTypeAttrDesc              //= 19   // deprecated
	ItemTypeAttrSpeed             //= 20
	ItemTypeAttrSlot              //= 21   // deprecated
	ItemTypeAttrMaxItems          //= 22   // deprecated
	ItemTypeAttrWeight            //= 23   // deprecated
	ItemTypeAttrWeapon            //= 24   // deprecated
	ItemTypeAttrAmmunition        //= 25   // deprecated
	ItemTypeAttrArmor             //= 26   // deprecated
	ItemTypeAttrMagicLevel        //= 27   // deprecated
	ItemTypeAttrMagicField        //= 28   // deprecated
	ItemTypeAttrWritable          //= 29   // deprecated
	ItemTypeAttrRotateTo          //= 30   // deprecated
	ItemTypeAttrDecay             //= 31   // deprecated
	ItemTypeAttrSpriteHash        //= 32
	ItemTypeAttrMinimapColor      //= 33
	ItemTypeAttr07                //= 34
	ItemTypeAttr08                //= 35
	ItemTypeAttrLight             //= 36
	ItemTypeAttrDecay2            //= 37   // deprecated
	ItemTypeAttrWeapon2           //= 38   // deprecated
	ItemTypeAttrAmmunition2       //= 39   // deprecated
	ItemTypeAttrArmor2            //= 40   // deprecated
	ItemTypeAttrWritable2         //= 41   // deprecated
	ItemTypeAttrLight2            //= 42
	ItemTypeAttrTopOrder          //= 43
	ItemTypeAttrWrtiable3         //= 44   // deprecated
	ItemTypeAttrWareId            //= 45
	ItemTypeAttrLast              //= 46
)

const (
	ClientVersion750     = 1
	ClientVersion755     //= 2
	ClientVersion760     //= 3
	ClientVersion770     //= 3
	ClientVersion780     //= 4
	ClientVersion790     //= 5
	ClientVersion792     //= 6
	ClientVersion800     //= 7
	ClientVersion810     //= 8
	ClientVersion811     //= 9
	ClientVersion820     //= 10
	ClientVersion830     //= 11
	ClientVersion840     //= 12
	ClientVersion841     //= 13
	ClientVersion842     //= 14
	ClientVersion850     //= 15
	ClientVersion854_OLD //= 16
	ClientVersion854     //= 17
	ClientVersion855     //= 18
	ClientVersion860_OLD //= 19
	ClientVersion860     //= 20
	ClientVersion861     //= 21
	ClientVersion862     //= 22,
	ClientVersion870     //= 23
	ClientVersion871     //= 24
	ClientVersion872     //= 25
	ClientVersion873     //= 26
	ClientVersion900     //= 27
	ClientVersion910     //= 28
	ClientVersion920     //= 29
	ClientVersion940     //= 30
	ClientVersion944_V1  //= 31
	ClientVersion944_V2  //= 32
	ClientVersion944_V3  //= 33
	ClientVersion944_V4  //= 34
	ClientVersion946     //= 35
	ClientVersion950     //= 36
	ClientVersion952     //= 37
	ClientVersion953     //= 38
	ClientVersion954     //= 39
	ClientVersion960     //= 40
	ClientVersion961     //= 41
)
