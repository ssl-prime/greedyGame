package controller

import (
	"greedyGame/bid_auction/util"

	"github.com/gin-gonic/gin"
)

//GetBid ...
func GetBid(c *gin.Context) {
	var (
		err    error
		adID   string
		status int
		resp   int
	)
	if adID, err = util.ValidateGetBid(c); err == nil {
		resp, err = util.GetBid(adID)
	}
	if err == nil && resp > 0 {
		status = 200
	} else {
		status = 204
	}
	c.JSON(status, gin.H{"highestBid": resp})
}
