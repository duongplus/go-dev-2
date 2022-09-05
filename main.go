package main

import (
	"github.com/gin-gonic/gin"
	"go-dev/module/restaurant/model"
	ginrestaurant "go-dev/module/restaurant/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dsn := os.Getenv("DBConnStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("restaurants")
		restaurants.GET("", getListRestaurant(db))
		restaurants.POST("", ginrestaurant.CreateRestaurantHandler(db))
		restaurants.GET("/:restaurant-id", ginrestaurant.GetRestaurantHandler(db))
		restaurants.PUT("/:restaurant-id", updateRestaurant(db))
	}

	if err := router.Run(); err != nil {

	}
}

func createRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := data.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func getRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.Restaurant

		id, err := strconv.Atoi(c.Param("restaurant-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func getListRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []restaurantmodel.Restaurant

		if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
			Count(&paging.Total).
			Limit(paging.Limit).
			Offset(offset).
			Order("id DESC").
			Find(&result).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})

	}
}

func updateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		id, err := strconv.Atoi(c.Param("restaurant-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func deleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.Restaurant

		id, err := strconv.Atoi(c.Param("restaurant-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}
