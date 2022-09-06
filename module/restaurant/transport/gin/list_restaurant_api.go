package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	restaurantbiz "go-dev/module/restaurant/biz"
	restaurantmodel "go-dev/module/restaurant/model"
	storagerestaurant "go-dev/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

func ListRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_ = paging.Validate()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storagerestaurant.NewSqlStore(db)
		biz := restaurantbiz.ListRestaurantStore(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
