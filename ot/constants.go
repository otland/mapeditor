package ot

const (
	TileFlagNone           = 0
	TileFlagProtectionZone = 1 << 0
	TileFlagTrashed        = 1 << 1
	TileFlagOptionalZone   = 1 << 2
	TileFlagNoLogout       = 1 << 3
	TileFlagHardcoreZone   = 1 << 4
	TileFlagRefresh        = 1 << 5

	// internal usage
	TileFlagHouse       = 1 << 6
	TileFlagTeleport    = 1 << 17
	TileFlagMagicField  = 1 << 18
	TileFlagMailBox     = 1 << 19
	TileFlagTrashHolder = 1 << 20
	TileFlagBed         = 1 << 21
	TileFlagDepot       = 1 << 22

	TileFlagLast = 1 << 24
)

const (
	OTBMNodeRootV2    = 1
	OTBMNodeMapData   = 2
	OTBMNodeTileArea  = 4
	OTBMNodeTile      = 5
	OTBMNodeItem      = 6
	OTBMNodeTowns     = 12
	OTBMNodeTown      = 13
	OTBMNodeHouseTile = 14
	OTBMNodeWaypoints = 15
	OTBMNodeWaypoint  = 16
)

const (
	OTBMAttrDescription = 1
	OTBMAttrTileFlags   = 3
	OTBMAttrItem        = 9
	OTBMAttrSpawnFile   = 11
	OTBMAttrHouseFile   = 13
)

const (
	OTBMItemAttrActionID       = 4
	OTBMItemAttrUniqueID       = 5
	OTBMItemAttrText           = 6
	OTBMItemAttrDesc           = 7
	OTBMItemAttrTeleDest       = 8
	OTBMItemAttrItem           = 9
	OTBMItemAttrDepotID        = 10
	OTBMItemAttrRuneCharges    = 12
	OTBMItemAttrHouseDoorID    = 14
	OTBMItemAttrCount          = 15
	OTBMItemAttrDuration       = 16
	OTBMItemAttrDecayState     = 17
	OTBMItemAttrWrittenDate    = 18
	OTBMItemAttrWrittenBy      = 19
	OTBMItemAttrSleepingGUID   = 20
	OTBMItemAttrSleepStart     = 21
	OTBMItemAttrCharges        = 22
	OTBMItemAttrContainerItems = 23
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
	ItemTypeAttrServerID     = 16
	ItemTypeAttrClientID     = 17
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
	ItemTypeAttrWareID            //= 45
	ItemTypeAttrLast              //= 46
)

const (
	ClientVersion750    = 1
	ClientVersion755    //= 2
	ClientVersion760    //= 3
	ClientVersion770    //= 3
	ClientVersion780    //= 4
	ClientVersion790    //= 5
	ClientVersion792    //= 6
	ClientVersion800    //= 7
	ClientVersion810    //= 8
	ClientVersion811    //= 9
	ClientVersion820    //= 10
	ClientVersion830    //= 11
	ClientVersion840    //= 12
	ClientVersion841    //= 13
	ClientVersion842    //= 14
	ClientVersion850    //= 15
	ClientVersion854Old //= 16
	ClientVersion854    //= 17
	ClientVersion855    //= 18
	ClientVersion860Old //= 19
	ClientVersion860    //= 20
	ClientVersion861    //= 21
	ClientVersion862    //= 22,
	ClientVersion870    //= 23
	ClientVersion871    //= 24
	ClientVersion872    //= 25
	ClientVersion873    //= 26
	ClientVersion900    //= 27
	ClientVersion910    //= 28
	ClientVersion920    //= 29
	ClientVersion940    //= 30
	ClientVersion944v1  //= 31
	ClientVersion944v2  //= 32
	ClientVersion944v3  //= 33
	ClientVersion944v4  //= 34
	ClientVersion946    //= 35
	ClientVersion950    //= 36
	ClientVersion952    //= 37
	ClientVersion953    //= 38
	ClientVersion954    //= 39
	ClientVersion960    //= 40
	ClientVersion961    //= 41
)
