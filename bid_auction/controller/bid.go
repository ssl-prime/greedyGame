package controller

import (
	"errors"
	"greedyGame/bid_auction/util"

	"github.com/gin-gonic/gin"
)

//AdRequest ...
func AdRequest(c *gin.Context) {
	var (
		err  error
		resp interface{}
		adID string
	)
	status := 200
	if adID, err = util.ValidateAdRequest(c); err == nil {
		if resp, err = util.AdRequest(adID); err == nil {

		}
	} else {
		err = errors.New("invalid AdRequest: " + err.Error())
	}
	if err != nil {
		status = 204
	}
	//fmt.Println(err)
	c.JSON(status, gin.H{"msg": resp})
}

//Acceptance ...
// func Acceptance(c *gin.Context) {
// 	var (
// 		err  error
// 		resp interface{}
// 	)
// 	if err = util.ValidateAcceptance(c); err == nil {
// 		// if resp, err = util.Acceptance; err == nil {

// 		// }
// 	} else {
// 		err = errors.New("invalid Acceptance param: " + err.Error())
// 	}
// 	c.JSON(200, gin.H{"err": err, "msg": resp})
// }

//DeleteSlot ...
func DeleteSlot(c *gin.Context) {
	var (
		err  error
		resp interface{}
	)
	if err = util.ValidateDeleteSlot(c); err == nil {

		if resp, err = util.DeleteSlot(); err == nil {

		}
	} else {
		err = errors.New("invalid Acceptance param: " + err.Error())
	}
	c.JSON(200, gin.H{"msg": resp})
}

//AddSlot ...
func AddSlot(c *gin.Context) {
	var (
		err        error
		resp       interface{}
		adSlotName string
	)
	if adSlotName, err = util.ValidateAddSlot(c); err == nil {

		if resp, err = util.AddSlot(adSlotName); err == nil {
		}
	} else {
		//err = errors.New("invalid ad slot info ")
		resp = err.Error()
	}
	c.JSON(200, gin.H{"msg": resp})
}

//GetSlot ...
func GetSlot(c *gin.Context) {
	var (
		err  error
		resp interface{}
	)
	if resp, err = util.GetSlot(); err == nil {

	}
	c.JSON(200, gin.H{"msg": resp})
}

//GetBidByAdPlacementID ...
func GetBidByAdPlacementID(c *gin.Context) {
	var (
		status int
		price  int
		err    error
	)
	slotID := c.Query("adplacementid")
	if slotID != `` {
		price, err = util.GetBidByAdPlacementID(slotID)
	}
	if err == nil && slotID != `` {
		status = 200
	} else {
		status = 204
	}
	c.JSON(status, gin.H{"price": price})
}
