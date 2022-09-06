package storagerestaurant

import (
	"context"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) UpdateRestaurant(ctx context.Context,
	cond map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
