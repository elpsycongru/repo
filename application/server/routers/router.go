package routers

import (
	"application/api"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api")
	{
		apiV1.GET("/hello", api.Hello)
		apiV1.GET("/queryPictureList", api.QueryPicture)
		apiV1.POST("/createPicture", api.CreatePicture)
		apiV1.GET("/queryByName", api.QueryByName)
		apiV1.GET("/queryBetweenTime", api.QueryBetweenTime)
		apiV1.GET("/queryByPeer", api.QueryByPeer)
		//apiV1.POST()
		/*
			apiV1.POST("/createRealEstate", v1.CreateRealEstate)
			apiV1.POST("/queryRealEstateList", v1.QueryRealEstateList)
			apiV1.POST("/createSelling", v1.CreateSelling)
			apiV1.POST("/createSellingByBuy", v1.CreateSellingByBuy)
			apiV1.POST("/querySellingList", v1.QuerySellingList)
			apiV1.POST("/querySellingListByBuyer", v1.QuerySellingListByBuyer)
			apiV1.POST("/updateSelling", v1.UpdateSelling)
			apiV1.POST("/createDonating", v1.CreateDonating)
			apiV1.POST("/queryDonatingList", v1.QueryDonatingList)
			apiV1.POST("/queryDonatingListByGrantee", v1.QueryDonatingListByGrantee)
			apiV1.POST("/updateDonating", v1.UpdateDonating)
		*/
	}
	// 静态文件路由
	return r
}
