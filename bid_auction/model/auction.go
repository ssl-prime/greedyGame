package model

//BidResp ...
type BidResp struct {
	Price int `json:"price"`
}

//GetBidStruct ...
type GetBidStruct struct {
	AdPlacementID string `json:"ad_placement_id" valid:"required"`
}
