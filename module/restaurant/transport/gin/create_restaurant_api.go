package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	restaurantbiz "go-dev/module/restaurant/biz"
	restaurantmodel "go-dev/module/restaurant/model"
	storagerestaurant "go-dev/module/restaurant/storage"
	"net/http"
)

func CreateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storagerestaurant.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
