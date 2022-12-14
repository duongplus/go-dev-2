package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	restaurantbiz "go-dev/module/restaurant/biz"
	restaurantmodel "go-dev/module/restaurant/model"
	storagerestaurant "go-dev/module/restaurant/storage"
	"net/http"
	"strconv"
)

func UpdateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param(common.RestaurantIdParam))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storagerestaurant.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurantById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
