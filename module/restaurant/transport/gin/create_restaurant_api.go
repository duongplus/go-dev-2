package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "go-dev/module/restaurant/biz"
	restaurantmodel "go-dev/module/restaurant/model"
	storagerestaurant "go-dev/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storagerestaurant.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
