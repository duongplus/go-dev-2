package storagerestaurant

import (
	"context"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := store.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
