package storagerestaurant

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
	"gorm.io/gorm"
)

func (store *sqlStore) FindRestaurant(ctx context.Context, cond map[string]interface{}, moreKey ...string) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	if err := store.db.Where(cond).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
