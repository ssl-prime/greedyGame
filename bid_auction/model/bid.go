package model

//AdRequest ...
type AdRequest struct {
	AdPlacementID string `json:"ad_placement_id" valid:"required"`
	AdObject      string `json:"ad_object" valid:"required"`
	TypeOfAd      string `json:"type_of_ad" `
}

//ObjectResp ...
type ObjectResp struct {
	Price      int    `json:"price"`
	AdObjectID string `json:"ad_object_id"`
}

//SlotAcceptance ...
type SlotAcceptance struct {
	AdObjectID string `json:"ad_object_id"`
	Price      int    `json:"price"`
}

//AddAdSlot ...
//many more field can be aaded in it
type AddAdSlot struct {
	Name string `json:"name" valid:"required"`
}
