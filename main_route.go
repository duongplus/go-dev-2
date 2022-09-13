package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	ginrestaurant "go-dev/module/restaurant/transport/gin"
	"go-dev/module/upload/transport/ginupload"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group(common.Restaurants)
		restaurants.GET("", ginrestaurant.ListRestaurantHandler(appCtx))
		restaurants.POST("", ginrestaurant.CreateRestaurantHandler(appCtx))
		restaurants.GET(fmt.Sprintf("/:%s", common.RestaurantIdParam), ginrestaurant.GetRestaurantHandler(appCtx))
		restaurants.PUT(fmt.Sprintf("/:%s", common.RestaurantIdParam), ginrestaurant.UpdateRestaurantHandler(appCtx))
		restaurants.PATCH(fmt.Sprintf("/:%s", common.RestaurantIdParam), ginrestaurant.SoftDeleteRestaurantHandler(appCtx))
		v1.POST("/upload", ginupload.Upload(appCtx))
		v1.GET(fmt.Sprintf("/uploads/:%s", common.ImageIds), ginupload.ListImageHandler(appCtx))
	}
}
