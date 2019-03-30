package main

import (
	"greedyGame/bid_auction/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/bid/v1")
	{
		v1.POST("/adRequest", controller.AdRequest)
		//v1.POST("/acceptance", controller.Acceptance)
		v1.POST("/deleteSlot", controller.DeleteSlot)
		v1.POST("/addSlot", controller.AddSlot)
		v1.GET("/getSlot", controller.GetSlot)
		v1.GET("/getbidbyadplacementid", controller.GetBidByAdPlacementID)
	}

	v2 := router.Group("/auction/v2")
	{
		v2.POST("/getbid", controller.GetBid)
		// 	v1.PUT("/updateInfo", controller.UpdateInfo)
		// 	v1.POST("/deleteInfo", controller.DeleteInfo)
		// 	v1.GET("/getkey", controller.GetKey)
		//
	}

	router.Run(":8090")
}
