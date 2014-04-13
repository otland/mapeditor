package main

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
