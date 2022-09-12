package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	restaurantbiz "go-dev/module/restaurant/biz"
	storagerestaurant "go-dev/module/restaurant/storage"
	"net/http"
)

func GetRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param(common.RestaurantIdParam))

		uid, err := common.FromBase58(c.Param(common.RestaurantIdParam))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id := int(uid.GetLocalID())

		store := storagerestaurant.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
