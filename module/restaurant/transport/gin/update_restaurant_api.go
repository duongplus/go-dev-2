package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "go-dev/module/restaurant/biz"
	restaurantmodel "go-dev/module/restaurant/model"
	storagerestaurant "go-dev/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storagerestaurant.NewSqlStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurantById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
