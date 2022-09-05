package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "go-dev/module/restaurant/biz"
	storagerestaurant "go-dev/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storagerestaurant.NewSqlStore(db)
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
