package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	restaurantbiz "go-dev/module/restaurant/biz"
	storagerestaurant "go-dev/module/restaurant/storage"
	"net/http"
	"strconv"
)

func SoftDeleteRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param(common.RestaurantIdParam))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagerestaurant.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewSoftDeleteRestaurantBiz(store)

		if err := biz.SoftDeleteRestaurantById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
