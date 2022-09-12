package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	ginrestaurant "go-dev/module/restaurant/transport/gin"
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
	}
}
