package storagerestaurant

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) UpdateRestaurant(ctx context.Context,
	cond map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
