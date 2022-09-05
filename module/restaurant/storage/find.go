package storagerestaurant

import (
	"context"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) FindRestaurant(ctx context.Context, cond map[string]interface{}, moreKey ...string) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	if err := store.db.Where(cond).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
